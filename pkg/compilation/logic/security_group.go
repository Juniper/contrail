package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/services"
)

// CreateSecurityGroup creates default AccessControlList's for the already created SecurityGroup.
func (s *Service) CreateSecurityGroup(
	ctx context.Context, request *services.CreateSecurityGroupRequest,
) (*services.CreateSecurityGroupResponse, error) {
	ingressACL, egressACL := defaultSecurityGroupACLs(request.SecurityGroup)

	_, err := s.WriteService.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		AccessControlList: ingressACL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ingress access control list")
	}

	_, err = s.WriteService.CreateAccessControlList(ctx, &services.CreateAccessControlListRequest{
		AccessControlList: egressACL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create egress access control list")
	}

	return &services.CreateSecurityGroupResponse{SecurityGroup: request.SecurityGroup}, nil
}
