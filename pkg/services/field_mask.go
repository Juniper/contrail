package services

import (
	"github.com/gogo/protobuf/types"
)

//MapToFieldMask returns updated fields masks.
func MapToFieldMask(request map[string]interface{}) types.FieldMask {
	mask := types.FieldMask{}
	mask.Paths = keys(request, "")
	return mask
}

type toMapper interface {
	ToMap() map[string]interface{}
}

func keys(m map[string]interface{}, prefix string) []string {
	result := []string{}
	for key, value := range m {
		switch v := value.(type) {
		case map[string]interface{}:
			if prefix != "" {
				result = append(result, keys(v, prefix+key+".")...)
			} else {
				result = append(result, keys(v, key+".")...)
			}
		case toMapper:
			m := v.ToMap()
			if prefix != "" {
				result = append(result, keys(m, prefix+key+".")...)
			} else {
				result = append(result, keys(m, key+".")...)
			}
		default:
			result = append(result, prefix+key)
		}
	}
	return result
}
