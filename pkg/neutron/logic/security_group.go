package logic

import (
	"context"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func (s *SecurityGroup) Read(ctx Context, id string) (Response, error) {
	resp, err := ctx.ReadService.GetSecurityGroup(context.Background(), &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		return nil, NewNeutronError(SecurityGroupNotFound, map[string]interface{}{
			"id": id,
		})
	}
	neutronSG := securityGroupVNCToNeutron(resp.SecurityGroup, READ)
	return neutronSG, nil
}

func securityGroupVNCToNeutron(sg *models.SecurityGroup, operation string) Response {
	// TODO
	return nil
}
