package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// FQNameToUUID translates fqName to UUID.
func (sv *ContrailTypeLogicService) FQNameToUUID(
	ctx context.Context,
	fqName []string,
	resourceType string,
) (string, error) {

	metadata, err := sv.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: resourceType, FQName: fqName})
	if err != nil {
		return "", err
	}

	return metadata.UUID, nil
}
