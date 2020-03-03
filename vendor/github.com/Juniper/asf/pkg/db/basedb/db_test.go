package basedb

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestScanRow(t *testing.T) {
	tests := []struct {
		name     string
		scanner  sqlScanner
		columns  Columns
		wantErr  bool
		expected RowData
	}{{
		name:    "nil",
		wantErr: true,
	}, {
		name:     "nil columns",
		scanner:  &mockScanner{},
		expected: RowData{},
	}, {
		name:     "empty data read",
		scanner:  &mockScanner{},
		columns:  Columns{"a": 0, "b": 1},
		expected: RowData{"a": nil, "b": nil},
	}, {
		name:    "bad column index",
		scanner: &mockScanner{},
		columns: Columns{"a": 0, "b": 3},
		wantErr: true,
	}, {
		name:    "scanner returns error",
		scanner: &mockScanner{err: errors.New("scanner error")},
		wantErr: true,
	}, {
		name:     "read scanner values",
		scanner:  &mockScanner{data: []interface{}{3, "value"}},
		columns:  Columns{"a": 1, "b": 0},
		expected: RowData{"a": "value", "b": 3},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := scanRow(tt.scanner, tt.columns)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, data)
		})
	}
}

type mockScanner struct {
	err  error
	data []interface{}
}

func (s mockScanner) Scan(values ...interface{}) error {
	if s.err != nil {
		return s.err
	}

	for i, v := range values {
		if i >= len(s.data) {
			return nil
		}
		if dest := v.(*interface{}); dest != nil {
			*dest = s.data[i]
		}
	}
	return nil
}
