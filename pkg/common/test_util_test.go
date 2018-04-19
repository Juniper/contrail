package common

import (
	"testing"

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
	yaml.Unmarshal([]byte(actualYAML), &actualData)

	var test1Data interface{}
	yaml.Unmarshal([]byte(test1YAML), &test1Data)
	AssertEqual(t, test1Data, actualData, "check same data")

	var test2Data interface{}
	yaml.Unmarshal([]byte(test2YAML), &test2Data)
	AssertEqual(t, test2Data, actualData, "check func works")
}
