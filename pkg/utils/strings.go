package utils

import (
	"bytes"
	"unicode"

	"github.com/volatiletech/sqlboiler/strmangle"
)

func isUpperOrDigit(c rune) bool {
	return unicode.IsUpper(c) || unicode.IsDigit(c)
}

//CamelToSnake translate camel case to snake case
func CamelToSnake(s string) string {
	var buf bytes.Buffer
	runes := []rune(s)
	for i, c := range runes {
		if i != 0 && i != len(runes)-1 && isUpperOrDigit(c) && !isUpperOrDigit(runes[i+1]) {
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
