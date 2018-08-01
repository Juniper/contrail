package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionForList(t *testing.T) {
	tests := []struct {
		name             string
		update           *PropCollectionUpdate
		expectedPosition int
		fails            bool
	}{
		{
			name: "add operation fails without field value",
			update: &PropCollectionUpdate{
				Operation: PropCollectionUpdateOperationAdd,
				Value:     []byte{},
			},
			fails: true,
		},
		{
			name: "add operation returns position 0",
			update: &PropCollectionUpdate{
				Operation: PropCollectionUpdateOperationAdd,
				Value:     []byte("value"),
			},
			fails:            false,
			expectedPosition: 0,
		},
		{
			name: "modify operation fails without field value",
			update: &PropCollectionUpdate{
				Operation: PropCollectionUpdateOperationModify,
				Value:     []byte{},
			},
			fails: true,
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
