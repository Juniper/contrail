package basedb

import "strings"

// Structure describes fields in schema.
type Structure map[string]interface{}

func (s *Structure) getPaths(prefix string) []string {
	var paths []string
	for k, v := range *s {
		p := prefix + "." + k
		switch o := v.(type) {
		case struct{}:
			paths = append(paths, p)
		case *Structure:
			paths = append(paths, o.getPaths(p)...)
		}
	}
	return paths
}

// GetInnerPaths gets all child for given fieldMask.
func (s *Structure) GetInnerPaths(fieldMask string) (paths []string) {
	innerStructure := s
	for _, segment := range strings.Split(fieldMask, ".") {
		switch o := (*innerStructure)[segment].(type) {
		case *Structure:
			innerStructure = o
		default:
			return nil
		}
	}
	return innerStructure.getPaths(fieldMask)
}
