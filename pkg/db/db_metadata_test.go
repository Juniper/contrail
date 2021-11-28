package db

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/asf/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestService_ListMetadata(t *testing.T) {

	dbMetadatas := []*models.Metadata{
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
		name    string
		args    []*models.Metadata
		want    []*models.Metadata
		wantErr bool
	}{
		{
			name: "Get multiple metadata using UUID and FQName",
			args: []*models.Metadata{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
			want: []*models.Metadata{
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
			name: "Get multiple metadata using UUIDs",
			args: []*models.Metadata{
				{
					UUID: "uuid-b",
				},
				{
					UUID: "uuid-c",
				},
			},
			want: []*models.Metadata{
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
			name: "Get multiple metadata using FQNames",
			args: []*models.Metadata{
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
				{
					FQName: []string{"default", "uuid-c"},
					Type:   "hoge",
				},
			},
			want: []*models.Metadata{
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
			name: "Provide only FQNames - fail",
			args: []*models.Metadata{
				{
					FQName: []string{"default", "uuid-b"},
				},
				{
					FQName: []string{"default", "uuid-c"},
				},
			},
			wantErr: true,
		},
		{
			name: "Get metadata using FQName",
			args: []*models.Metadata{
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
			want: []*models.Metadata{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
		},
		{
			name: "Get metadata using UUID",
			args: []*models.Metadata{
				{
					UUID: "uuid-b",
				},
			},
			want: []*models.Metadata{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
		},

		{
			name: "Get single metadata using UUID and FQName",
			args: []*models.Metadata{
				{
					UUID: "uuid-b",
				},
				{
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
			want: []*models.Metadata{
				{
					UUID:   "uuid-b",
					FQName: []string{"default", "uuid-b"},
					Type:   "hoge",
				},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	for _, metadata := range dbMetadatas {
		err := db.CreateMetadata(ctx, metadata)
		assert.NoError(t, err)
		defer db.DeleteMetadata(ctx, metadata.UUID) // nolint: errcheck
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

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
