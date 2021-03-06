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
	_ "golang.org/x/lint/golint"
	// Required to vendor them, used by schema tool
	_ "github.com/Juniper/asf/pkg/cache"
	_ "github.com/Juniper/asf/pkg/db"
	_ "github.com/Juniper/asf/tools"
)
