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
	Name             string
	Value            string
	Type             string
	GoType           string
	NewGoType        string
	RedisType        string
	RedisTypeReplace bool
	Setter           bool
	Getter           bool
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
				r.m.{{.Name}} = new({{.NewGoType}})
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

	log.Println("-------------------------------------")
	log.Println(file)

	//for _, dep := range file.Dependency {
	//	gp := p.Generator.GoPackageName(generator.GoImportPath(dep))
	//	log.Println(gp)
	//}
	// range fields
	for _, field := range message.Field {
		name := generator.CamelCase(*field.Name)
		log.Println(field)
		log.Println(field.Options)

		generateField := &generateField{
			Name:   name,
			Value:  "r.m." + name,
			Type:   field.Type.String(),
			Setter: true,
			Getter: true,
		}

		generateField.GoType, _ = p.Generator.GoType(message, field)
		generateField.NewGoType = strings.Replace(generateField.GoType, "*", "", -1)
		generateField.RedisType = generator.CamelCase(generateField.GoType)
		// redis go just have 64bit function
		if strings.Contains(generateField.RedisType, "32") {
			generateField.RedisType = strings.Replace(generateField.RedisType, "32", "64", -1)
			generateField.RedisTypeReplace = true
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

	p.generateRedisHashGetBasicFieldFunc(data)
}

// get basic type from redis by hash field
const getBasicTypeFromRedisHashFuncTemplate = `
// get {{.MessageName}} {{.Name}} field value with key 
func (r *{{.MessageName}}RedisController) Get{{.Name}}(key string) ({{.Name}} {{.GoType}}, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get {{.Name}} field
	if value, err := {{.RedisPkg}}.{{.RedisType}}(conn.Do("HGET", key, "{{.Name}}")); err != nil {
		return {{.Name}}, err
	} else {
		{{- if .RedisTypeReplace}}
			r.m.{{.Name}} = {{.GoType}}(value)
		{{else}}
			r.m.{{.Name}} = value
		{{- end -}}
    }

	return r.m.{{.Name}}, nil
}
`

// set basic type from redis by hash field
const setBasicTypeFromRedisHashFuncTemplate = `
// set {{.MessageName}} {{.Name}} field with key and {{.Name}} 
func (r *{{.MessageName}}RedisController) Set{{.Name}}(key string, {{.Name}} {{.GoType}}) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set {{.Name}} field
	r.m.{{.Name}} = {{.Name}}
	_, err = conn.Do("HSET", key, "{{.Name}}", {{.Name}})

	return
}
`

// get message type from redis by hash field
const getMessageTypeFromRedisHashFuncTemplate = `
// get {{.MessageName}} {{.Name}} field value with key 
func (r *{{.MessageName}}RedisController) Get{{.Name}}(key string) ({{.Name}} {{.GoType}}, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get {{.Name}} field
	if value, err := {{.RedisPkg}}.{{.RedisType}}(conn.Do("HGET", key, "{{.Name}}")); err != nil {
		return {{.Name}}, err
	} else {
		// unmarshal {{.Name}}
		r.m.{{.Name}} = new({{.NewGoType}})
		if err = {{.CodecPkg}}.Unmarshal(value, r.m.{{.Name}}); err != nil {
			return {{.Name}}, err
		}
    }

	return r.m.{{.Name}}, nil
}
`

// set message type from redis by hash field
const setMessageTypeFromRedisHashFuncTemplate = `
// set {{.MessageName}} {{.Name}} field with key and {{.Name}} 
func (r *{{.MessageName}}RedisController) Set{{.Name}}(key string, {{.Name}} {{.GoType}}) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal {{.Name}}
	if r.m.{{.Name}} != nil {
		r.m.{{.Name}} = {{.Name}}
		if data, err := {{.CodecPkg}}.Marshal(r.m.{{.Name}}); err != nil {
			return err
		} else {
			// set {{.Name}} field
			_, err = conn.Do("HSET", key, "{{.Name}}", data)
			return err 
		}
	}

	return nil
}
`

// generate Redis basic type get handler by hash type
func (p *plugin) generateRedisHashGetBasicFieldFunc(data *generateData) {

	type FiledType struct {
		*generateField
		MessageName string
		RedisPkg    string
		CodecPkg    string
	}

	for _, field := range data.Fields {

		if !field.Setter || !field.Getter {
			continue
		}

		fieldData := FiledType{
			MessageName: data.MessageName,
			RedisPkg:    data.RedisPkg,
			CodecPkg:    data.CodecPkg,
		}
		fieldData.generateField = field

		getTemplateName := ""
		setTemplateName := ""
		tpy := descriptor.FieldDescriptorProto_Type_value[field.Type]
		switch descriptor.FieldDescriptorProto_Type(tpy) {
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
			descriptor.FieldDescriptorProto_TYPE_FLOAT,
			descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_UINT64,
			descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_UINT32,
			descriptor.FieldDescriptorProto_TYPE_FIXED64,
			descriptor.FieldDescriptorProto_TYPE_SFIXED64,
			descriptor.FieldDescriptorProto_TYPE_FIXED32,
			descriptor.FieldDescriptorProto_TYPE_SFIXED32,
			descriptor.FieldDescriptorProto_TYPE_BOOL,
			descriptor.FieldDescriptorProto_TYPE_STRING:
			getTemplateName = getBasicTypeFromRedisHashFuncTemplate
			setTemplateName = setBasicTypeFromRedisHashFuncTemplate
		case descriptor.FieldDescriptorProto_TYPE_ENUM:
			getTemplateName = getBasicTypeFromRedisHashFuncTemplate
			fieldData.RedisType = "Int64"
			fieldData.RedisTypeReplace = true
			setTemplateName = setBasicTypeFromRedisHashFuncTemplate
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			getTemplateName = getMessageTypeFromRedisHashFuncTemplate
			fieldData.RedisType = "Bytes"
			setTemplateName = setMessageTypeFromRedisHashFuncTemplate
		default:
			return
		}

		if getTemplateName != "" {
			tmpl, _ := template.New("hash-get").Parse(getTemplateName)
			if err := tmpl.Execute(p.Buffer, fieldData); err != nil {
				log.Println(getTemplateName, fieldData)
			}
		}

		if setTemplateName != "" {
			tmpl, _ := template.New("hash-set").Parse(setTemplateName)
			if err := tmpl.Execute(p.Buffer, fieldData); err != nil {
				log.Println(setTemplateName, fieldData)
			}
		}
	}
}
