package schema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAPI(t *testing.T) {
	api, err := MakeAPI("test_data/schema")
	assert.Nil(t, err, "API reading failed")
	fmt.Println(api)
	openAPI, err := api.ToOpenAPI()
	assert.Nil(t, err, "OpenAPI generation failed")
	SaveFile("test_output/openapi.json", openAPI)
}
