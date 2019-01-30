package logic

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilters_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		expected Filters
		data     string
		wantErr  bool
	}{
		{
			name:     "empty filter",
			data:     `{"data": []}`,
			expected: Filters{"data": nil},
		},
		{
			name:     "string filter",
			data:     `{"data": ["false", "true"]}`,
			expected: Filters{"data": []string{"false", "true"}},
		},
		{
			name:     "boolean filter",
			data:     `{"data": [false, true]}`,
			expected: Filters{"data": []string{"false", "true"}},
		},
		{
			name:     "integer filter",
			data:     `{"data": [123, 1414]}`,
			expected: Filters{},
			wantErr:  true,
		},
		{
			name:     "string filter",
			data:     `{"data": "hoge"}`,
			expected: Filters{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual Filters
			if err := json.Unmarshal([]byte(tt.data), &actual); (err != nil) != tt.wantErr {
				t.Errorf("Filters.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
