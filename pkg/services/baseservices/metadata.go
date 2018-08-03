package baseservices

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

//MetadataGetter provides getter for metadata.
type MetadataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*basemodels.MetaData, error)
}
