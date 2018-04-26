package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

var actualYAML = `
string: value
list:
- element1
- element2
- element3
map:
  key1: value1
  key2: value2
number: 1
bool: true
nilValue: null
`

var test1YAML = `
string: value
list:
- element1
- element2
- element3
map:
  key1: value1
number: 1
bool: true
`

var test2YAML = `
string: value
list:
- element1
- $any
- $any
map:
  key1: $any
number: $int
bool: true
nilValue: $null
`

func TestAssertEquals(t *testing.T) {

	var actualData interface{}
	err := yaml.Unmarshal([]byte(actualYAML), &actualData)
	assert.NoError(t, err, "no error expected")

	var test1Data interface{}
	err = yaml.Unmarshal([]byte(test1YAML), &test1Data)
	assert.NoError(t, err, "no error expected")
	AssertEqual(t, test1Data, actualData, "check same data")

	var test2Data interface{}
	err = yaml.Unmarshal([]byte(test2YAML), &test2Data)
	assert.NoError(t, err, "no error expected")
	AssertEqual(t, test2Data, actualData, "check func works")
}
