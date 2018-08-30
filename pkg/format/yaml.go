package format

import (
	"gopkg.in/yaml.v2"
)

//MustYAML Marshal yaml
func MustYAML(data interface{}) string {
	b, err := yaml.Marshal(data)
	if err != nil {
		return ""
	}
	return string(b)
}
