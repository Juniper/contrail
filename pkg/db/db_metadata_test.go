package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestService_ListMetadata(t *testing.T) {
	tests := []struct {
		name        string
		dbMetadatas []*models.MetaData
		args        []*models.FQNameUUIDPair
		want        []*models.MetaData
		wantErr     bool
	}{
		{
			name: "Simple test",
			dbMetadatas: []*models.MetaData{
				{
					UUID:   "uuid-a",
					FQName: []string{"default", "uuid-a"},
				},
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
				},
				{
					UUID:   "uuid-c",
					FQName: []string{"default", "uuid-c"},
				},
			},
			args: []*models.FQNameUUIDPair{
				{
					FQName: []string{"default", "uuid-b"},
				},
			},
			want: []*models.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
			defer cancel()

			for _, metadata := range tt.dbMetadatas {
				err := db.CreateMetaData(ctx, metadata)
				assert.NoError(t, err)
				defer db.DeleteMetaData(ctx, metadata.UUID)
			}

			got, err := db.ListMetadata(ctx, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
