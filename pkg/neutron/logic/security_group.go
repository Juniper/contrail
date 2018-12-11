package logic

import (
	"context"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"strings"
)

func (s *SecurityGroup) Read(ctx Context, id string) (Response, error) {
	resp, err := ctx.ReadService.GetSecurityGroup(context.Background(), &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		return nil, NewNeutronError(SecurityGroupNotFound, ErrorFields{
			"id": id,
		})
	}
	neutronSG := securityGroupVNCToNeutron(resp.SecurityGroup, READ)
	return neutronSG, nil
}

func securityGroupVNCToNeutron(sg *models.SecurityGroup, operation string) Response {
	sgNeutron := SecurityGroup{}
	sgNeutron.Name = sg.GetName()
	sgNeutron.TenantID = strings.Replace(sg.GetParentUUID(), "-", "", -1)

	if sg.GetDisplayName() != "" {
		sgNeutron.Name = sg.GetDisplayName()
	} else {
		fgName := sg.GetFQName()
		sgNeutron.Name =fgName[len(fgName) - 1]
	}

	if sg.IDPerms.GetDescription() != "" {
		sgNeutron.Description = sg.IDPerms.GetDescription()
	}

	// TODO: security_group_rules

	sgNeutron.CreatedAt = sg.IDPerms.GetCreated()
	sgNeutron.UpdatedAt = sg.IDPerms.GetLastModified()

	// TODO: finish it
	//if contrailExtensionsEnabled {
	//	sgNeutron.FGName = sg.GetFQName()
	//}
	//
	return sgNeutron
}
