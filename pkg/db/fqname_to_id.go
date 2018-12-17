package db

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

func (db *Service) FQNameToID(
	ctx context.Context,
	request *services.FQNameToIDRequest,
) (*services.FQNameToIDResponse, error) {
	metadata, err := db.GetMetadata(ctx, basemodels.Metadata{Type: request.Type, FQName: request.FQName})
	if err != nil {
		//TODO adding Project
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return nil, errors.Wrapf(err, errMsg)
	}

	//TODO permissions check

	return &services.FQNameToIDResponse{
		UUID: metadata.UUID,
	}, nil
}
