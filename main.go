// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/galaxyobe/protoc-gen-redis/plugin"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"path/filepath"
)

func main() {
	gen := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		gen.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, gen.Request); err != nil {
		gen.Error(err, "parsing input proto")
	}

	if len(gen.Request.FileToGenerate) == 0 {
		gen.Fail("no files to generate")
	}

	useGogoImport := false
	// Match parsing algorithm from Generator.CommandLineParameters
	for _, parameter := range strings.Split(gen.Request.GetParameter(), ",") {
		kvp := strings.SplitN(parameter, "=", 2)
		// We only care about key-value pairs where the key is "gogoimport"
		if len(kvp) != 2 || kvp[0] != "gogoimport" {
			continue
		}
		useGogoImport, err = strconv.ParseBool(kvp[1])
		if err != nil {
			gen.Error(err, "parsing gogoimport option")
		}
	}

	gen.CommandLineParameters(gen.Request.GetParameter())

	myPlugin := plugin.NewPlugin(useGogoImport)

	gen.WrapTypes()
	gen.SetPackageNames()
	gen.BuildTypeNameMap()
	gen.GeneratePlugin(myPlugin)

	var Response = new(plugin_go.CodeGeneratorResponse)

	for i := 0; i < len(gen.Response.File); i++ {
		gen.Response.File[i].Name = proto.String(strings.Replace(*gen.Response.File[i].Name, ".pb.go", ".redis.go", -1))
		name := strings.Split(filepath.Base(*gen.Response.File[i].Name), ".")[0]
		if g, ok := myPlugin.GenerateMap[name]; ok {
			if g {
				Response.File = append(Response.File, gen.Response.File[i])
			}
		}
	}

	// Send back the results.
	data, err = proto.Marshal(Response)
	if err != nil {
		gen.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		gen.Error(err, "failed to write output proto")
	}
}
