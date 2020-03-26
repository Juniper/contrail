package httputil

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyHeader(t *testing.T) {
	tests := []struct {
		name               string
		src, dst, expected http.Header
	}{{
		name: "nil",
	}, {
		name:     "one value",
		src:      http.Header{"Foo": []string{"bar"}},
		dst:      http.Header{},
		expected: http.Header{"Foo": []string{"bar"}},
	}, {
		name:     "already existing values",
		src:      http.Header{"Foo": []string{"bar"}},
		dst:      http.Header{contentTypeHeader: []string{applicationJSONValue}},
		expected: http.Header{contentTypeHeader: []string{applicationJSONValue}, "Foo": []string{"bar"}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyHeader(tt.src, tt.dst)
			assert.Equal(t, tt.dst, tt.expected)
		})
	}
}

func TestCloneHeader(t *testing.T) {
	tests := []struct {
		name string
		h    http.Header
		want http.Header
	}{{
		name: "nils",
	}, {
		name: "empty", h: http.Header{}, want: http.Header{},
	}, {
		name: "one value",
		h:    http.Header{contentTypeHeader: []string{"A"}},
		want: http.Header{contentTypeHeader: []string{"A"}},
	}, {
		name: "more headers",
		h: http.Header{
			"Foo-Header":      []string{"X"},
			contentTypeHeader: []string{"A", "B"},
		},
		want: http.Header{
			"Foo-Header":      []string{"X"},
			contentTypeHeader: []string{"A", "B"},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneHeader(tt.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("Clone() mutated original Headers, got %v, want %v", tt.h, tt.want)
			}
		})
	}
}
