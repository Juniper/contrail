package baseservices

import (
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"golang.org/x/net/context"
)

//MetadataGetter provides getter for metadata.
type MetadataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*basemodels.MetaData, error)
}
