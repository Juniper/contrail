package services

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// FQNameToID translates FQName to corresponding UUID stored in database
func (svc *ContrailService) FQNameToID(
	ctx context.Context,
	request *FQNameToIDRequest,
) (*FQNameToIDResponse, error) {
	metadata, err := svc.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: request.Type, FQName: request.FQName})
	if err != nil {
		//TODO adding Project
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return nil, errors.Wrapf(err, errMsg)
	}

	//TODO permissions check

	return &FQNameToIDResponse{
		UUID: metadata.UUID,
	}, nil
}
