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
			Type:   "hoge",
		},
		{
			UUID:   "uuid-b",
			FQName: []string{"default", "uuid-b"},
			Type:   "hoge",
		},
		{
			UUID:   "uuid-c",
			FQName: []string{"default", "uuid-c"},
			Type:   "hoge",
		},
	}

	tests := []struct {
		name        string
		dbMetadatas []*basemodels.MetaData
		args        []*basemodels.MetaData
		want        []*basemodels.MetaData
		wantErr     bool
	}{
		{
			name:        "Get multiple metadatas using UUID and FQName",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.MetaData{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
				{
					UUID:   "uuid-c",
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
		},
		{
			name:        "Get multiple metadatas using UUIDs",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.MetaData{
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
					Type:   "hoge",
				},
				{
					UUID:   "uuid-c",
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
		},
		{
			name:        "Get multiple metadatas using FQNames",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.MetaData{
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
				{
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
				{
					UUID:   "uuid-c",
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
		},
		{
			name:        "Get metadata using FQName",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.MetaData{
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
		},
		{
			name:        "Get metadata using UUID",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.MetaData{
				{
					UUID: "uuid-b",
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
		},

		{
			name:        "Get single metadata using UUID and FQName",
			dbMetadatas: dbMetadatas,
			args: []*basemodels.MetaData{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
			want: []*basemodels.MetaData{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
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
