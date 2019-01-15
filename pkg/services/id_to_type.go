package services

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// IDToType translates UUID to corresponding type stored in database
func (svc *ContrailService) IDToType(
	ctx context.Context,
	request *IDToTypeRequest,
) (*IDToTypeResponse, error) {
	metadata, err := svc.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{UUID: request.UUID})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to retrieve metadata for UUID %v", request.UUID)
	}

	return &IDToTypeResponse{
		Type: metadata.Type,
	}, nil
}
