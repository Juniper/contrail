package format

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Map map[string]string

func (f *Map) ApplyMap(m map[string]interface{}) error {
	return nil
}

type SomeStruct struct {
	Integer                  int  `json:"int_tag"`
	Boolean                  bool `json:"boolean_tag"`
	String                   string
	ArrayOfStrings           []string         `json:"str_array"`
	ArrayOfStructs           []AnotherStruct  `json:"struct_array"`
	ArrayOfPointersToStructs []*AnotherStruct `json:"pointer_to_struct_array"`
	Struct                   AnotherStruct
	PointerToStruct          *AnotherStruct
	Interface                interface{}
	Map                      Map
}

type StructWithPointerToNonStruct struct {
	PointerToString *string
}

type AnotherStruct struct {
	A string
}

func TestApplyMapWithJSONUnmarshalMaps(t *testing.T) {
	var someString = "hoge"
	tests := []struct {
		name     string
		expected interface{}
		fails    bool
	}{
		{
			name: "nil_object",
		},
		{
			name: "pointer_to_non_struct_field",
			expected: &StructWithPointerToNonStruct{
				PointerToString: &someString,
			},
			fails: true,
		},
		{
			name: "integer_field",
			expected: &SomeStruct{
				Integer: 234,
			},
		},
		{
			name: "array_field",
			expected: &SomeStruct{
				ArrayOfStrings: []string{"aa", "bb"},
			},
		},
		{
			name: "struct_field",
			expected: &SomeStruct{
				Struct: AnotherStruct{
					A: "hoge",
				},
			},
		},
		{
			name: "struct_array_field",
			expected: &SomeStruct{
				ArrayOfStructs: []AnotherStruct{
					{
						A: "aa",
					},
					{
						A: "bb",
					},
				},
			},
		},
		{
			name: "pointer_to_struct_array_field",
			expected: &SomeStruct{
				ArrayOfPointersToStructs: []*AnotherStruct{
					{
						A: "aa",
					},
					{
						A: "bb",
					},
				},
			},
		},
		{
			name: "pointer_to_struct_field",
			expected: &SomeStruct{
				PointerToStruct: &AnotherStruct{
					A: "hoge",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw, err := json.Marshal(tt.expected)
			require.NoError(t, err)
			var m map[string]interface{}
			require.NoError(t, json.Unmarshal(raw, &m))
			var output interface{}
			if t := reflect.TypeOf(tt.expected); t != nil {
				output = reflect.New(indirect(t)).Interface()
			}

			err = ApplyMap(m, output)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, output)
			}
		})
	}
}

func TestApplyMap(t *testing.T) {
	var someInt = 6
	var nilStruct *SomeStruct

	type args struct {
		m map[string]interface{}
		o interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected interface{}
		fails    bool
	}{
		{
			name: "apply_nil_map_to_nil",
		},
		{
			name: "apply_non_nil_map_to_nil",
			args: args{
				m: map[string]interface{}{
					"hoge": "hoge",
				},
			},
			fails: true,
		},
		{
			name: "pointer_to_non_struct",
			args: args{
				o: &someInt,
				m: map[string]interface{}{
					"hoge": "hoge",
				},
			},
			expected: &someInt,
			fails:    true,
		},
		{
			name: "non_struct",
			args: args{
				o: someInt,
				m: map[string]interface{}{
					"hoge": "hoge",
				},
			},
			expected: someInt,
			fails:    true,
		},
		{
			name: "struct",
			args: args{
				o: &SomeStruct{},
				m: map[string]interface{}{
					"hoge": "hoge",
				},
			},
			expected: &SomeStruct{},
		},
		{
			name: "pointer_to_nil_struct",
			args: args{
				o: &nilStruct,
				m: map[string]interface{}{
					"hoge": "hoge",
				},
			},
			expected: &nilStruct,
		},
		{
			name: "apply_fields",
			args: args{
				o: &SomeStruct{},
				m: map[string]interface{}{
					"int_tag":     1,
					"boolean_tag": true,
					"String":      "hoge",
				},
			},
			expected: &SomeStruct{
				Integer: 1,
				Boolean: true,
				String:  "hoge",
			},
		},
		{
			name: "pointer_to_struct_under_interface_field",
			args: args{
				o: &SomeStruct{
					Interface: &AnotherStruct{},
				},
				m: map[string]interface{}{
					"Interface": map[string]interface{}{
						"A": "aaa",
					},
				},
			},
			expected: &SomeStruct{
				Interface: &AnotherStruct{A: "aaa"},
			},
		},
		{
			name: "struct_under_interface_field",
			args: args{
				o: &SomeStruct{
					Interface: AnotherStruct{},
				},
				m: map[string]interface{}{
					"Interface": map[string]interface{}{
						"A": "aaa",
					},
				},
			},
			expected: &SomeStruct{
				Interface: AnotherStruct{},
			},
			fails: true,
		},
		{
			name: "apply_invalid_map_to_integer_field",
			args: args{
				o: &SomeStruct{},
				m: map[string]interface{}{
					"int_tag": "hoge",
				},
			},
			expected: &SomeStruct{},
		},
		{
			name: "apply_invalid_map_to_string_field",
			args: args{
				o: &SomeStruct{},
				m: map[string]interface{}{
					"String": 123,
				},
			},
			expected: &SomeStruct{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ApplyMap(tt.args.m, tt.args.o)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, tt.args.o)
		})
	}
}
