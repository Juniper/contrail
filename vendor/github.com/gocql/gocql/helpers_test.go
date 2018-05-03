package gocql

import (
	"reflect"
	"testing"
)

func TestGetCassandraType_Set(t *testing.T) {
	typ := getCassandraType("set<text>")
	set, ok := typ.(CollectionType)
	if !ok {
		t.Fatalf("expected CollectionType got %T", typ)
	} else if set.typ != TypeSet {
		t.Fatalf("expected type %v got %v", TypeSet, set.typ)
	}

	inner, ok := set.Elem.(NativeType)
	if !ok {
		t.Fatalf("expected to get NativeType got %T", set.Elem)
	} else if inner.typ != TypeText {
		t.Fatalf("expected to get %v got %v for set value", TypeText, set.typ)
	}
}

func TestGetCassandraType(t *testing.T) {
	tests := []struct {
		input string
		exp   TypeInfo
	}{
		{
			"set<text>", CollectionType{
				NativeType: NativeType{typ: TypeSet},

				Elem: NativeType{typ: TypeText},
			},
		},
		{
			"map<text, varchar>", CollectionType{
				NativeType: NativeType{typ: TypeMap},

				Key:  NativeType{typ: TypeText},
				Elem: NativeType{typ: TypeVarchar},
			},
		},
		{
			"list<int>", CollectionType{
				NativeType: NativeType{typ: TypeList},
				Elem:       NativeType{typ: TypeInt},
			},
		},
		{
			"tuple<int, int, text>", TupleTypeInfo{
				NativeType: NativeType{typ: TypeTuple},

				Elems: []TypeInfo{
					NativeType{typ: TypeInt},
					NativeType{typ: TypeInt},
					NativeType{typ: TypeText},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := getCassandraType(test.input)

			// TODO(zariel): define an equal method on the types?
			if !reflect.DeepEqual(got, test.exp) {
				t.Fatalf("expected %v got %v", test.exp, got)
			}
		})
	}
}
