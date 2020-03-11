package baseservices

import (
	"context"

	"github.com/Juniper/asf/pkg/models"
)

//MetadataGetter provides getter for metadata.
type MetadataGetter interface {
	GetMetadata(ctx context.Context, requested models.Metadata) (*models.Metadata, error)
	ListMetadata(ctx context.Context, requested []*models.Metadata) ([]*models.Metadata, error)
}
