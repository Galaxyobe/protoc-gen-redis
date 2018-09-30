package plugin

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/vanity"
	"github.com/galaxyobe/protoc-gen-redis/proto"
	"text/template"
	"log"
	"strings"
)

const (
	contextPkg      = "context"
	redisPkg        = "github.com/gomodule/redigo/redis"
	mapStructurePkg = "github.com/mitchellh/mapstructure"
	jsonPkg         = "github.com/json-iterator/go"
)

type generateField struct {
	Name     string
	Value    string
	Type     string
	TypeName string
}

type generateData struct {
	Package         string
	MessageName     string
	ContextPkg      string
	RedisPkg        string
	MapStructurePkg string
	CodecPkg        string
	StorageType     string
	Expired         bool
	Fields          []*generateField
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
		expired := proto.GetBoolExtension(message.Options, redis.E_Ttl, true)

		// generateData
		data := &generateData{
			ContextPkg:  p.NewImport(contextPkg).Use(),
			RedisPkg:    p.NewImport(redisPkg).Use(),
			Expired:     expired,
			MessageName: generator.CamelCaseSlice(message.TypeName()),
		}

		storageCodec, _ := proto.GetExtension(message.Options, redis.E_StorageCodec)
		if storageCodec != nil && *storageCodec.(*string) == "json" {
			data.CodecPkg = p.NewImport(jsonPkg).Use()
		} else {
			data.CodecPkg = "proto"
		}

		storageType, _ := proto.GetExtension(message.Options, redis.E_StorageType)
		p.generateRedisControllerCommon(data, file, message)

		if storageType != nil && *storageType.(*string) == "hash" {
			data.MapStructurePkg = p.NewImport(mapStructurePkg).Use()
			// hash handler
			p.generateRedisHashFunc(data, file, message)
		} else {
			// string handler
			p.generateRedisStringFunc(data, file, message)
		}
	}
}

// redis controller common template
const redisControllerCommonTemplate = `
// new {{.MessageName}} redis controller with redis pool
func (m *{{.MessageName}}) RedisController(pool *{{.RedisPkg}}.Pool) *{{.MessageName}}RedisController {
	return &{{.MessageName}}RedisController{
		pool: pool,
		m:    m,
	}
}

// {{.MessageName}} redis controller
type {{.MessageName}}RedisController struct {
	pool *{{.RedisPkg}}.Pool
	m    *{{.MessageName}}
}

// new {{.MessageName}} redis controller with redis pool
func New{{.MessageName}}RedisController(pool *{{.RedisPkg}}.Pool) *{{.MessageName}}RedisController {
	return &{{.MessageName}}RedisController{pool: pool}
}

// get {{.MessageName}}
func (r *{{.MessageName}}RedisController) {{.MessageName}}() *{{.MessageName}} {
	return r.m
}
`

// generate redis controller common
func (p *plugin) generateRedisControllerCommon(data *generateData, file *generator.FileDescriptor, message *generator.Descriptor) {
	tmpl, _ := template.New("RedisController").Parse(redisControllerCommonTemplate)
	tmpl.Execute(p.Buffer, data)
}

// load from redis by string type
const loadFromRedisStringFuncTemplate = `
// load {{.MessageName}} from redis string with context and key
func (r *{{.MessageName}}RedisController) Load(ctx {{.ContextPkg}}.Context, key string) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// load data from redis string
	data, err := {{.RedisPkg}}.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}

	// unmarshal data to StringStorageType
	return {{.CodecPkg}}.Unmarshal(data, r.m)
}
`

// store to redis by string type
const storeToRedisStringFuncTemplate = `
// store {{.MessageName}} to redis string with context and key
func (r *{{.MessageName}}RedisController) Store(ctx {{.ContextPkg}}.Context, key string{{ if .Expired }}, ttl uint64{{ end }}) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal {{.MessageName}} to []byte
	data, err := {{.CodecPkg}}.Marshal(r.m)
	if err != nil {
		return err
	}

	{{ if .Expired }}
	// use redis string store {{.MessageName}} data with expire second
	_, err = conn.Do("SETEX", key, ttl, data)
	{{- else }}
	// use redis string store {{.MessageName}} data
	_, err = conn.Do("SET", key, data)
	{{- end }}

	return err
}
`

// generate Redis handler by string type
func (p *plugin) generateRedisStringFunc(data *generateData, file *generator.FileDescriptor, message *generator.Descriptor) {
	tmpl, _ := template.New("StoreToRedis").Parse(storeToRedisStringFuncTemplate)
	if err := tmpl.Execute(p.Buffer, data); err != nil {
		log.Println("storeToRedisStringFuncTemplate", data)
	}
	tmpl, _ = template.New("StoreToRedis").Parse(loadFromRedisStringFuncTemplate)
	if err := tmpl.Execute(p.Buffer, data); err != nil {
		log.Println("loadFromRedisStringFuncTemplate", data)
	}
}

// load from redis by hash type
const loadFromRedisHashFuncTemplate = `
// load {{.MessageName}} from redis hash with context and key
func (r *{{.MessageName}}RedisController) Load(ctx {{.ContextPkg}}.Context, key string) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// load data from redis hash
	data, err := {{.RedisPkg}}.ByteSlices(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}

	// parse redis hash field name and value
	structure := make(map[string]interface{})
	for i := 0; i < len(data); i += 2 {
		switch string(data[i]) {
		{{- range .Fields}}
			{{- if eq .Type "TYPE_MESSAGE" }}
			case "{{.Name}}":
				// unmarshal {{.Name}}
				r.m.{{.Name}} = new({{.TypeName}})
				if err := {{$.CodecPkg}}.Unmarshal(data[i+1], r.m.{{.Name}}); err != nil {
					return err	
				}
			{{- end }}
		{{- end }}
		default:
			structure[string(data[i])] = string(data[i+1])
		}
	}

	// use mapstructure weak decode structure to {{.MessageName}}
	return {{.MapStructurePkg}}.WeakDecode(structure, r.m)
}
`

// store to redis by hash type
const storeToRedisHashFuncTemplate = `
// store {{.MessageName}} to redis hash with context and key
func (r *{{.MessageName}}RedisController) Store(ctx {{.ContextPkg}}.Context, key string{{ if .Expired }}, ttl uint64{{ end }}) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// make args
	args := make([]interface{}, 0)

	// add redis key
	args = append(args, key)

	// add redis field and value
	{{- range .Fields}}
		{{- if eq .Type "TYPE_MESSAGE" }}
			// marshal {{.Name}}
			if r.m.{{.Name}} != nil {
				{{.Name}}, {{.Name}}Error := {{$.CodecPkg}}.Marshal(r.m.{{.Name}})
				if {{.Name}}Error != nil {
					return {{.Name}}Error
				}
				args = append(args, "{{.Name}}", {{.Name}})
			}
		{{- else }}
			args = append(args, "{{.Name}}", {{.Value}})
		{{- end }}
	{{- end}}

	{{if .Expired }}
	// use redis hash store {{.MessageName}} data with expire second
	err := conn.Send("MULTI")
	if err != nil{
		return err
	}
	err = conn.Send("HMSET", args...)
	if err != nil{
		return err
	}
	err = conn.Send("EXPIRE", key, ttl)
	if err != nil{
		return err
	}
	_, err = conn.Do("EXEC")
	{{- else }}
	// use redis hash store {{.MessageName}} data
	_, err := conn.Do("HMSET", args...)
	{{- end }}

	return err
}
`

// generate Redis handler by hash type
func (p *plugin) generateRedisHashFunc(data *generateData, file *generator.FileDescriptor, message *generator.Descriptor) {

	log.Println(file)
	// range fields
	for _, field := range message.Field {
		name := generator.CamelCase(*field.Name)
		log.Println(field)
		generateField := &generateField{
			Name:  name,
			Value: "r.m." + name,
			Type:  field.Type.String(),
		}
		switch *field.Type {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			typeName := *field.TypeName
			// same package
			selfPackage := "." + *file.Package + "."
			if strings.HasPrefix(typeName, selfPackage) {
				generateField.TypeName = strings.Split(typeName, selfPackage)[1]
			} else {
				// TODO: 其他proto文件的引用
				generateField.TypeName = typeName[1:]
			}
		default:
		}
		log.Println(generateField)
		data.Fields = append(data.Fields, generateField)
	}

	tmpl, _ := template.New("hash").Parse(loadFromRedisHashFuncTemplate)
	if err := tmpl.Execute(p.Buffer, data); err != nil {
		log.Println("loadFromRedisHashFuncTemplate", data)
	}

	tmpl, _ = template.New("hash").Parse(storeToRedisHashFuncTemplate)
	if err := tmpl.Execute(p.Buffer, data); err != nil {
		log.Println("storeToRedisHashFuncTemplate", data)
	}
}
