package strutil

import (
	"bytes"
	"github.com/gogo/protobuf/types"
	"github.com/volatiletech/sqlboiler/strmangle"
	"strings"
	"unicode"
)

func isUpperOrDigit(c rune) bool {
	return unicode.IsUpper(c) || unicode.IsDigit(c)
}

//CamelToSnake translate camel case to snake case
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

//SnakeToCamel translates snake case to camel case
func SnakeToCamel(s string) string {
	return strmangle.TitleCase(s)
}

//ContainsString check if a string is in a string list
func ContainsString(list []string, a string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

//CheckPath check if fieldMask includes provided path
func CheckPath(fieldMask *types.FieldMask, path []string) bool {
	genPath := strings.Join(path, ".")
	return ContainsString(fieldMask.GetPaths(), genPath)
}
