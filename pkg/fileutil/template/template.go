package template

import (
	"regexp"

	"github.com/flosch/pongo2"
)

//Apply applies src template to a context and returns output
func Apply(templateSrc string, context map[string]interface{}) ([]byte, error) {
	template, err := pongo2.FromFile(templateSrc)
	if err != nil {
		return nil, err
	}
	output, err := template.ExecuteBytes(context)
	if err != nil {
		return nil, err
	}
	// strip empty lines in output content
	emptyLinesregex, _ := regexp.Compile("\n[ \r\n\t]*\n") // nolint: errcheck
	outputString := emptyLinesregex.ReplaceAllString(string(output), "\n")
	// remove trailing spaces and tabs
	trailingRegex, _ := regexp.Compile("([^ \t\r\n])[ \t]") // nolint: errcheck
	outputString = trailingRegex.ReplaceAllString(string(outputString), "")
	return []byte(outputString), nil
}
