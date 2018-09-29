package plugin

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
	"protoc-gen-redis/proto"
	"text/template"
)

const (
	contextPkg = "context"
	redisPkg   = "github.com/gomodule/redigo/redis"
)

type generateData struct {
	TypeName    string
	ContextPkg  string
	RedisPkg    string
	StorageType string
	Expired     bool
}

type plugin struct {
	*generator.Generator
	generator.PluginImports
	fmtPkg        generator.Single
	protoPkg      generator.Single
	useGogoImport bool
}

func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "redis"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	if len(file.Messages()) == 0 {
		return
	}

	if !p.useGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}

	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.fmtPkg = p.NewImport("fmt")

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.generateRedisFunc(file, msg)
	}
}

func (p *plugin) generateRedisFunc(file *generator.FileDescriptor, message *generator.Descriptor) {
	// enable redis
	if proto.GetBoolExtension(message.Options, redis.E_Enabled, false) {
		StorageType, _ := proto.GetExtension(message.Options, redis.E_StorageType)
		expired := proto.GetBoolExtension(message.Options, redis.E_Ttl, true)

		// generateData
		data := &generateData{
			ContextPkg: p.NewImport(contextPkg).Use(),
			RedisPkg:   p.NewImport(redisPkg).Use(),
			Expired:    expired,
			TypeName:   generator.CamelCaseSlice(message.TypeName()),
		}

		// hash handler
		if StorageType != nil && *StorageType.(*string) == "hash" {
			p.generateRedisHash(data, file, message)
		} else {
			p.generateRedisProto(data, file, message)
		}
	}
}

// load from redis by proto
const loadFromRedisFuncTemplate = `
// load {{.TypeName}} from redis
func (this *{{.TypeName}}) LoadFromRedis(ctx {{.ContextPkg}}.Context, key string) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// load data from redis string
	data, err := {{.RedisPkg}}.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}

	// unmarshal data to Demo
	err = proto.Unmarshal(data, this.Demo)
	if err != nil {
		return err
	}

	return nil
}
`

// store to redis by proto
const storeToRedisFuncTemplate = `
// store {{.TypeName}} to redis
// {{.TypeName}} will not expire when ttl is 0
func (this *{{.TypeName}}) StoreToRedis(ctx {{.ContextPkg}}.Context, key string{{ if .Expired }}, ttl uint64{{ end }}) error {
		// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// marshal {{.TypeName}} to []byte
	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	{{ if .Expired }}
	// use redis string store the {{.TypeName}} data with expire second
	_, err = conn.Do("SETEX", key, ttl, data)
	{{ else }}
	// use redis string store the {{.TypeName}} data
	_, err = conn.Do("SET", key, data)
	{{ end }}

	return err
}
`

// generate Redis handler by proto type
func (p *plugin) generateRedisProto(data *generateData, file *generator.FileDescriptor, message *generator.Descriptor) {
	tmpl, _ := template.New("StoreToRedis").Parse(storeToRedisFuncTemplate)
	tmpl.Execute(p.Buffer, data)
	tmpl, _ = template.New("StoreToRedis").Parse(loadFromRedisFuncTemplate)
	tmpl.Execute(p.Buffer, data)
}

// load from redis by hash
const loadFromRedisHashFuncTemplate = `
// load {{.TypeName}} from redis by hash type
func (this *{{.TypeName}}) LoadFromRedis(ctx {{.ContextPkg}}.Context, key string) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// load data from redis string
	data, err := {{.RedisPkg}}.Values(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}

	err = {{.RedisPkg}}.ScanStruct(data, this.Demo)
	
	return err
}
`

// load from redis hash by one key
const loadFromRedisHashByKeyFuncTemplate string = `
// store {{.TypeName}} to redis by one hash key
// field param is full match with original struct name
func (this *{{.TypeName}}) LoadFromRedisByKey(ctx {{.ContextPkg}}.Context, key string, field string) (interface{}, error) {
		// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// get by key
	return conn.Do("HGET", key, field)
}
`

// store to redis by hash
const storeToRedisHashFuncTemplate string = `
// store {{.TypeName}} to redis by hash type 
// {{.TypeName}} will not expire when ttl is 0
func (this *{{.TypeName}}) StoreToRedis(ctx {{.ContextPkg}}.Context, key string{{ if .Expired }}, ttl uint64{{ end }}) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	fields := make([]interface{}, 0)

	v := reflect.ValueOf(this.Demo)
	if v.IsValid() == false {
		return errors.New("reflect is InValid")
	}

	//找到最后指向的值，或者空指针，空指针是需要进行初始化的
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	st := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// 判断是否为可导出字段
		if v.Field(i).CanInterface() {
			name := st.FieldByIndex([]int{i}).Name
			if strings.Contains(name, "XXX_") {
				continue
			}
			fields = append(fields, name, v.Field(i).Interface())
		}
	}

	if len(fields) == 0 {
		return errors.New("struct has no fields can export")
	}

	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, fields...)

	{{ if .Expired }}
	// set expire
    conn.Send("MULTI")
	conn.Send("HMSET", args...)
	conn.Send("EXPIRE", key, ttl)
	_, err := conn.Do("EXEC")
	{{ else }}
	_, err := conn.Do("HMSET", args...)
	{{ end }}

	return err	
}
`

// store to redis by hash key
const storeToRedisHashByKeyFuncTemplate string = `
// store {{.TypeName}} to redis by hash key 
// field param is full match with original struct name
// DemoRedis will not expire when ttl is 0
func (this *{{.TypeName}}) StoreToRedisByKey(ctx {{.ContextPkg}}.Context, key string, field string, val interface{}{{ if .Expired }}, ttl uint64{{ end }}) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	{{ if .Expired }}
	// set expire
	conn.Send("MULTI")
	conn.Send("HSET", key, field, val)
	conn.Send("EXPIRE", key, ttl)
	_, err := conn.Do("EXEC")
	{{ else }}
	_, err := conn.Do("HSET", key, field, val)
	{{ end }}
	return err
}
`

// generate Redis handler by hash type
func (p *plugin) generateRedisHash(data *generateData, file *generator.FileDescriptor, message *generator.Descriptor) {
	tmpl, _ := template.New("hash").Parse(loadFromRedisHashFuncTemplate)
	tmpl.Execute(p.Buffer, data)
	tmpl, _ = template.New("hash").Parse(loadFromRedisHashByKeyFuncTemplate)
	tmpl.Execute(p.Buffer, data)
	tmpl, _ = template.New("hash").Parse(storeToRedisHashFuncTemplate)
	tmpl.Execute(p.Buffer, data)
	tmpl, _ = template.New("hash").Parse(storeToRedisHashByKeyFuncTemplate)
	tmpl.Execute(p.Buffer, data)
}
