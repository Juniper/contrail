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
		return nil, NewNeutronError(SecurityGroupNotFound, ErrorFields{
			"id": id,
		})
	}
	neutronSG := securityGroupVNCToNeutron(resp.SecurityGroup, READ)
	return neutronSG, nil
}

func securityGroupVNCToNeutron(sg *models.SecurityGroup, operation string) Response {
	sgNeutron := SecurityGroupResponse{}
	sgNeutron.ID = sg.GetUUID()
	sgNeutron.TenantID = contrailUUIDToNeutronID(sg.GetParentUUID())

	if sg.GetDisplayName() != "" {
		sgNeutron.Name = sg.GetDisplayName()
	} else {
		fgName := sg.GetFQName()
		sgNeutron.Name = fgName[len(fgName)-1]
	}

	if sg.IDPerms.GetDescription() != "" {
		sgNeutron.Description = sg.IDPerms.GetDescription()
	}

	sgNeutron.CreatedAt = sg.IDPerms.GetCreated()
	sgNeutron.UpdatedAt = sg.IDPerms.GetLastModified()

	// TODO: Uncomment and adjust when contrail extension support will be done.
	//if contrailExtensionsEnabled {
	//	sgNeutron.FQName = sg.GetFQName()
	//}

	return sgNeutron
}

func readSecurityGroupRules() []*security_group_rule {
	// TODO: implement it
	return nil
}
