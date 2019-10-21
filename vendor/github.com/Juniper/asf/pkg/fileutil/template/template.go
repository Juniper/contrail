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
	// TODO: Do not compile regex inside a function.
	// strip empty lines in output content
	emptyLinesregex, err := regexp.Compile("\n[ \r\n\t]*\n")
	if err != nil {
		return nil, err
	}
	outputString := emptyLinesregex.ReplaceAllString(string(output), "\n")
	// remove trailing spaces and tabs
	trailingRegex, err := regexp.Compile("([^ \t\r\n]) +(\n)+")
	if err != nil {
		return nil, err
	}
	outputString = trailingRegex.ReplaceAllString(outputString, "$1\n")
	return []byte(outputString), nil
}
