package gendoc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"io/ioutil"
	"path"
	"strings"
)

// PluginOptions encapsulates options for the plugin. The type of renderer, template file, and the name of the output
// file are included.
type PluginOptions struct {
	Type         RenderType
	TemplateFile string
	OutputFile   string
}

// ParseOptions parses plugin options from a CodeGeneratorRequest. It does this by splitting the `Parameter` field from
// the request object and parsing out the type of renderer to use and the name of the file to be generated.
//
// The parameter (`--doc_opt`) must be of the format <TYPE|TEMPLATE_FILE>,<OUTPUT_FILE>. The file will be written to the
// directory specified with the `--doc_out` argument to protoc.
func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{
		Type:         RenderTypeHTML,
		TemplateFile: "",
		OutputFile:   "index.html",
	}

	params := req.GetParameter()
	if params == "" {
		return options, nil
	}

	if !strings.Contains(params, ",") {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	parts := strings.Split(params, ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	options.TemplateFile = parts[0]
	options.OutputFile = path.Base(parts[1])

	renderType, err := NewRenderType(options.TemplateFile)
	if err == nil {
		options.Type = renderType
		options.TemplateFile = ""
	}

	return options, nil
}

// RunPlugin compiles the documentation and generates the CodeGeneratorResponse to send back to protoc. It does this
// by rendering a template based on the options parsed from the CodeGeneratorRequest.
func RunPlugin(request *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	result := parser.ParseCodeRequest(request)
	template := NewTemplate(result)

	options, err := ParseOptions(request)
	if err != nil {
		return nil, err
	}

	customTemplate := ""

	if options.TemplateFile != "" {
		data, err := ioutil.ReadFile(options.TemplateFile)
		if err != nil {
			return nil, err
		}

		customTemplate = string(data)
	}

	output, err := RenderTemplate(options.Type, template, customTemplate)
	if err != nil {
		return nil, err
	}

	resp := new(plugin_go.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(options.OutputFile),
		Content: proto.String(string(output)),
	})

	return resp, nil
}
