package types

import (
	"context"

	"github.com/Juniper/contrail/extension/pkg/services"
)

// CreateTagType performs Tag Type specific logic.
func (sv *ContrailTypeLogicService) CreateTagType(
	ctx context.Context, request *services.CreateTagTypeRequest,
) (*services.CreateTagTypeResponse, error) {
	var response services.CreateTagTypeResponse

	return response, nil
}
