package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropCollectionUpdate_PositionForList(t *testing.T) {
	correctValue := []byte("value")
	var emptyValue []byte

	tests := []struct {
		name             string
		update           *PropCollectionUpdate
		expectedPosition int
		fails            bool
	}{
		{
			name: "fails for add operation without value",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationAdd,
				Value:     emptyValue,
			},
			fails: true,
		},
		{
			name: "returns position 0 for add operation",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationAdd,
				Value:     correctValue,
			},
			expectedPosition: 0,
		},
		{
			name: "fails for modify operation without value",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationModify,
				Value:     emptyValue,
			},
			fails: true,
		},
		{
			name: "fails for modify operation with invalid position",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationModify,
				Value:     correctValue,
				Position:  "five",
			},
			fails: true,
		},
		{
			name: "returns position for modify operation",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationModify,
				Value:     correctValue,
				Position:  "5",
			},
			expectedPosition: 5,
		},
		{
			name: "fails for delete operation with invalid position",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationDelete,
				Value:     correctValue,
				Position:  "five",
			},
			fails: true,
		},
		{
			name: "returns position for delete operation",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationDelete,
				Value:     correctValue,
				Position:  "5",
			},
			expectedPosition: 5,
		},
		{
			name: "fails for set operation",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationSet,
			},
			fails: true,
		},
		{
			name: "fails for invalid operation",
			update: &PropCollectionUpdate{
				Operation: "invalid",
			},
			fails: true,
		},
		{
			name: "returns position for mixed case operation string",
			update: &PropCollectionUpdate{
				Operation: "aDd",
				Value:     correctValue,
			},
			expectedPosition: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position, err := tt.update.positionForList()

			if tt.fails {
				assert.Error(t, err)
				assert.Equal(t, 0, position)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPosition, position)
			}
		})
	}
}

func TestPropCollectionUpdate_ValidateForMap(t *testing.T) {
	correctValue := []byte("value")
	var emptyValue []byte

	tests := []struct {
		name   string
		update *PropCollectionUpdate
		fails  bool
	}{
		{
			name: "fails for set operation without value",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationSet,
				Value:     emptyValue,
			},
			fails: true,
		},
		{
			name: "fails for delete operation without position",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationDelete,
				Position:  "",
			},
			fails: true,
		},
		{
			name: "succeeds for mixed case operation string",
			update: &PropCollectionUpdate{
				Operation: "sEt",
				Value:     correctValue,
			},
		},
		{
			name: "fails for add operation",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationAdd,
				Value:     correctValue,
			},
			fails: true,
		},
		{
			name: "fails for modify operation",
			update: &PropCollectionUpdate{
				Operation: propCollectionUpdateOperationModify,
				Value:     correctValue,
			},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.update.validateForMap()

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
