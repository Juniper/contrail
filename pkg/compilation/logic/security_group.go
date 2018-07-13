package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

func (s *Service) CreateSecurityGroup(
	ctx context.Context, request *services.CreateSecurityGroupRequest,
) (*services.CreateSecurityGroupResponse, error) {
	_, err := s.api.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		&models.AccessControlList{
			Name: "ingress-access-control-list",
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create ingress access control list")
	}

	_, err = s.api.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		&models.AccessControlList{
			Name: "egress-access-control-list",
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create egress access control list")
	}

	return &services.CreateSecurityGroupResponse{request.SecurityGroup}, nil
}
