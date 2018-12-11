package logic

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

func (s *SecurityGroup) Read(ctx context.Context, id string) (Response, error) {
	resp, err := ctx.ReadService.GetSecurityGroup(context.Background(), &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		return nil, NewNeutronError(SecurityGroupNotFound, map[string]interface{}{
			"id": id,
		})
	}
	neutronSG := securityGroupVNCToNeutron(ctx, resp.SecurityGroup, READ)
	return neutronSG, nil
}

func securityGroupVNCToNeutron(ctx context.Context, sg *models.SecurityGroup, operation string) Response {
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

	sgNeutron.SecurityGroupRules, _ = readSecurityGroupRules(ctx)

	// TODO: Uncomment and adjust when contrail extension support will be done.
	//if contrailExtensionsEnabled {
	//	sgNeutron.FQName = sg.GetFQName()
	//}
	//ListCloudSecurityGroupRuleRequest
	return sgNeutron
}

func readSecurityGroupRules(ctx context.Context) ([]*security_group_rule, error) {
	// TODO: implement it
	sgRules, err := ctx.ReadService.ListCloudSecurityGroupRule(
		ctx,
		&services.ListCloudSecurityGroupRuleRequest{
			Spec: &baseservices.ListSpec{
				Detail: true, // TODO
			},
		},
	)

	fmt.Println(sgRules) // TODO: delete

	if err != nil {
		return nil, err
	}

	return nil, nil
}
