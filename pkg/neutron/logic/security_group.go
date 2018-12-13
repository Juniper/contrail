package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// ReadLLL logic
func (sg *SecurityGroup) ReadLLL(ctx context.Context, rp RequestParameters, f Filters, fields Fields) (Response, error) {

	//TODO ensure default security group exists

	var securityGroups []*models.SecurityGroup
	var err error
	if ids, ok := f["id"]; ok {
		securityGroups, err = listSecurityGroups(ctx, rp, ids, nil)
		if err != nil {
			return nil, err
		}
	} else if !rp.RequestContext.IsAdmin {
		securityGroups, err = listSecurityGroups(
			ctx, rp, nil, []string{rp.RequestContext.Tenant},
		)
		if err != nil {
			return nil, err
		}
	} else if rp.RequestContext.TenantID != "" {
		securityGroups, err = listSecurityGroups(
			ctx, rp, nil, f["tenant_id"],
		)
		if err != nil {
			return nil, err
		}
	} else {
		securityGroups, err = listSecurityGroups(ctx, rp, nil, nil)
		if err != nil {
			return nil, err
		}
	}

	_ = securityGroups

	return nil, errors.New("not implemented")
}

func listSecurityGroups(
	ctx context.Context, rp RequestParameters, uuids []string, parentUUIDs []string,
) ([]*models.SecurityGroup, error) {
	sgResponse, err := rp.ReadService.ListSecurityGroup(
		ctx,
		&services.ListSecurityGroupRequest{
			Spec: &baseservices.ListSpec{
				ObjectUUIDs: uuids,
				ParentUUIDs: parentUUIDs,
				Detail:      true,
			},
		},
	)

	return sgResponse.GetSecurityGroups(), err
}
