package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/plugin"
)

// A ParseResult contains a set of parsed proto files
type ParseResult struct {
	Files []*File
}

// GetFile returns the parsed proto file specified by the name (base name without path)
// e.g. pr.GetFile("Vehicle.proto")
func (pr *ParseResult) GetFile(name string) *File {
	for _, f := range pr.Files {
		if f.Name == name {
			return f
		}
	}

	return nil
}

// ParseCodeRequest iterates through all the proto files in the code gen request, parses them, and finally adds them to
// the returned ParseResult
func ParseCodeRequest(req *plugin_go.CodeGeneratorRequest) *ParseResult {
	result := new(ParseResult)

	for _, file := range req.GetProtoFile() {
		result.Files = append(result.Files, parseProtoFile(file))
	}

	return result
}
