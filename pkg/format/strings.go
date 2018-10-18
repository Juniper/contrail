package format

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/gogo/protobuf/types"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// CamelToSnake translate camel case to snake case.
func CamelToSnake(s string) string {
	var buf bytes.Buffer
	runes := []rune(s)
	for i, c := range runes {
		if i != 0 && i != len(runes)-1 && isUpperOrDigit(c) && (!isUpperOrDigit(runes[i+1]) || !isUpperOrDigit(runes[i-1])) {
			buf.WriteRune('_')
		}
		buf.WriteRune(unicode.ToLower(c))
	}
	return buf.String()
}

func isUpperOrDigit(c rune) bool {
	return unicode.IsUpper(c) || unicode.IsDigit(c)
}

// SnakeToCamel translates snake case to camel case.
func SnakeToCamel(s string) string {
	return strmangle.TitleCase(s)
}

// ContainsString check if a string is in a string list.
func ContainsString(list []string, a string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

// CheckPath check if fieldMask includes provided path.
func CheckPath(fieldMask *types.FieldMask, path []string) bool {
	genPath := strings.Join(path, ".")
	return ContainsString(fieldMask.GetPaths(), genPath)
}

// RemoveFromStringSlice removes given values from slice of strings.
// It preserves order of values.
func RemoveFromStringSlice(slice []string, values map[string]struct{}) []string {
	if len(values) == 0 {
		return slice
	}

	var indexesToRemove []int
	for i, v := range slice {
		if _, ok := values[v]; ok {
			indexesToRemove = append(indexesToRemove, i)
		}
	}

	for i, v := range indexesToRemove {
		slice = append(slice[:v-i], slice[v-i+1:]...)
	}

	return slice
}
