// +build tools

package tools

import (
	_ "github.com/Juniper/asf/cmd/contrailschema"
	_ "github.com/gogo/protobuf/protoc-gen-gogofaster"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/mattn/goveralls"
	_ "github.com/mitchellh/gox"
	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
	_ "golang.org/x/tools/cmd/goimports"
)
