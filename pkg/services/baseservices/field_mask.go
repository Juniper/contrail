package baseservices

import (
	"github.com/gogo/protobuf/types"
)

//MapToFieldMask returns updated fields masks.
func MapToFieldMask(data map[string]interface{}) types.FieldMask {
	var mask types.FieldMask
	mask.Paths = paths(data, "")
	return mask
}

type toMapper interface {
	ToMap() map[string]interface{}
}

func paths(data map[string]interface{}, prefix string) []string {
	var result []string
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			if prefix != "" {
				result = append(result, paths(v, prefix+key+".")...)
			} else {
				result = append(result, paths(v, key+".")...)
			}
		case toMapper:
			m := v.ToMap()
			if prefix != "" {
				result = append(result, paths(m, prefix+key+".")...)
			} else {
				result = append(result, paths(m, key+".")...)
			}
		default:
			result = append(result, prefix+key)
		}
	}
	return result
}
