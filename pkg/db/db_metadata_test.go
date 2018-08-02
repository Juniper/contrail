package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

func TestService_ListMetadata(t *testing.T) {

	dbMetadatas := []*basemodels.MetaData{
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
	}

	tests := []struct {
		name        string
		dbMetadatas []*basemodels.MetaData
		args        []*basemodels.FQNameUUIDPair
		want        []*basemodels.MetaData
		wantErr     bool
	}{
		{
			name:        "Get multiple metadatas using UUID and FQName",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.FQNameUUIDPair{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-c"},
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
				},
				{
					UUID:   "uuid-c",
					FQName: []string{"default", "uuid-c"},
				},
			},
		},
		{
			name:        "Get multiple metadatas using UUIDs",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.FQNameUUIDPair{
				{
					UUID: "uuid-b",
				},
				{
					UUID: "uuid-c",
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
				},
				{
					UUID:   "uuid-c",
					FQName: []string{"default", "uuid-c"},
				},
			},
		},
		{
			name:        "Get multiple metadatas using FQNames",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.FQNameUUIDPair{
				{
					FQName: []string{"default", "uuid-b"},
				},
				{
					FQName: []string{"default", "uuid-c"},
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
				},
				{
					UUID:   "uuid-c",
					FQName: []string{"default", "uuid-c"},
				},
			},
		},
		{
			name:        "Get metadata using FQName",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.FQNameUUIDPair{
				{
					FQName: []string{"default", "uuid-b"},
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
				},
			},
		},
		{
			name:        "Get metadata using UUID",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.FQNameUUIDPair{
				{
					UUID: "uuid-b",
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
				},
			},
		},

		{
			name:        "Get single metadata using UUID and FQName",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.FQNameUUIDPair{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-b"},
				},
			},
			want: []*basemodels.MetaData{
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
				defer db.DeleteMetaData(ctx, metadata.UUID) //nolint
			}

			got, err := db.ListMetadata(ctx, tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, len(tt.want), len(got))
			for _, metadata := range got {
				assert.Contains(t, tt.want, metadata)
			}
		})
	}
}
