package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/stretchr/testify/assert"
)

const (
	// TODO: Use relative path
	templFilePath = "/var/tmp/contrail/template_test.tmpl"
)

func TestApply(t *testing.T) {
	tests := []struct {
		name         string
		context      map[string]interface{}
		templateData string
		expectedData string
	}{
		{
			name: "Add context values",
			context: map[string]interface{}{
				"key1": interface{}("value1"),
				"key2": interface{}("value2"),
			},
			templateData: strings.Join([]string{
				"line_a: {{ key1 }}",
				"line_b: {{ key2 }}",
				"line_c: otherValue\n",
			}, "\n"),
			expectedData: strings.Join([]string{
				"line_a: value1",
				"line_b: value2",
				"line_c: otherValue\n",
			}, "\n"),
		},
		{
			name:    "Remove trailing spaces",
			context: nil,
			templateData: strings.Join([]string{
				"line a: ",
				"line b: value",
				"line c:",
				"line d: \n",
			}, "\n"),
			expectedData: strings.Join([]string{
				"line a:",
				"line b: value",
				"line c:",
				"line d:\n",
			}, "\n"),
		},
		{
			name:         "Remove multiple trailing spaces",
			context:      nil,
			templateData: "abc   \n",
			expectedData: "abc\n",
		},
		{
			name:    "Remove extra lines",
			context: nil,
			templateData: strings.Join([]string{
				"a\n",
				"b",
			}, "\n"),
			expectedData: strings.Join([]string{
				"a",
				"b",
			}, "\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTemporaryTemplateFile(t, tt.templateData, templFilePath)
			defer func() {
				removeErr := os.Remove(templFilePath)
				assert.NoErrorf(t, removeErr, "Couldn't remove %s file. Please remove it manually", templFilePath)
			}()
			result, err := Apply(templFilePath, tt.context)
			assert.NoError(t, err, "Couldn't apply change to template")
			assert.Equal(t, []byte(tt.expectedData), result)
		})
	}
}

func createTemporaryTemplateFile(t *testing.T, data string, destFile string) {
	err := os.MkdirAll(filepath.Dir(destFile), 0664)
	assert.NoError(t, err, "Couldn't create directory for template test file")
	err = fileutil.WriteToFile(destFile, []byte(data), 0777)
	assert.NoError(t, err, "Couldn't create template test file")
}
