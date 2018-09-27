package plugin

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
	"text/template"
	"github.com/gogo/protobuf/proto"
	"protoc-gen-redis/proto"
)

const (
	contextPkg = "context"
	redisPkg   = "github.com/gomodule/redigo/redis"
)

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
	if proto.GetBoolExtension(message.Options, redis.E_Enabled, false) {
		p.generateRedisLoadFunc(file, message)
		p.generateRedisStoreFunc(file, message, proto.GetBoolExtension(message.Options, redis.E_Ttl, true))
	}
}

type generateData struct {
	TypeName   string
	ContextPkg string
	RedisPkg   string
	Expired    bool
}

const loadFromRedisFuncTemplate = `
// load {{.TypeName}} from redis
func (m *{{.TypeName}}) LoadFromRedis(ctx {{.ContextPkg}}.Context, conn {{.RedisPkg}}.Conn, key string) error {
	// load data from redis string
	data, err := {{.RedisPkg}}.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}
	// unmarshal data to {{.TypeName}}
	err = proto.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}
`

func (p *plugin) generateRedisLoadFunc(file *generator.FileDescriptor, message *generator.Descriptor) {
	data := &generateData{
		ContextPkg: p.NewImport(contextPkg).Use(),
		RedisPkg:   p.NewImport(redisPkg).Use(),
	}

	data.TypeName = generator.CamelCaseSlice(message.TypeName())

	tmpl, _ := template.New("LoadFromRedis").Parse(loadFromRedisFuncTemplate)
	tmpl.Execute(p.Buffer, data)
}

const storeToRedisFuncTemplate = `
// store {{.TypeName}} to redis
// {{.TypeName}} will not expire when ttl is 0
func (m *{{.TypeName}}) StoreToRedis(ctx {{.ContextPkg}}.Context, conn {{.RedisPkg}}.Conn, key string{{ if .Expired }}, ttl uint64{{ end }}) error {
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
	if err != nil {
		return err
	}

	return nil
}
`

func (p *plugin) generateRedisStoreFunc(file *generator.FileDescriptor, message *generator.Descriptor, expired bool) {
	data := &generateData{
		ContextPkg: p.NewImport(contextPkg).Use(),
		RedisPkg:   p.NewImport(redisPkg).Use(),
		Expired:    expired,
	}

	data.TypeName = generator.CamelCaseSlice(message.TypeName())

	tmpl, _ := template.New("StoreToRedis").Parse(storeToRedisFuncTemplate)
	tmpl.Execute(p.Buffer, data)
}
