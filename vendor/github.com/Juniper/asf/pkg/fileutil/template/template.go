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
	regex, _ := regexp.Compile("\n[ \r\n\t]*\n") // nolint: errcheck
	outputString := regex.ReplaceAllString(string(output), "\n")
	return []byte(outputString), nil
}
