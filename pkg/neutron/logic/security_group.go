package logic

import (
	"context"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

func (s *SecurityGroup) Read(rp RequestParameters, id string) (Response, error) {
	resp, err := rp.ReadService.GetSecurityGroup(context.Background(), &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		return nil, NewNeutronError(SecurityGroupNotFound, map[string]interface{}{
			"id": id,
		})
	}

	return securityGroupVNCToNeutron(rp, resp.SecurityGroup)
}

func securityGroupVNCToNeutron(rp RequestParameters, sg *models.SecurityGroup) (*SecurityGroupResponse, error) {
	var err error
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

	sgNeutron.SecurityGroupRules, err = readSecurityGroupRules(rp, sg.GetUUID())

	if err != nil {
		return nil, err
	}

	// TODO(pawel_drapiewski): Uncomment and adjust when contrail extension support will be done.
	//if contrailExtensionsEnabled {
	//	sgNeutron.FQName = sg.GetFQName()
	//}

	return &sgNeutron, nil
}

func readSecurityGroupRules(rp RequestParameters, parentID string) ([]*SecurityGroupRule, error) {
	sgRules, err := rp.ReadService.ListCloudSecurityGroupRule(
		context.Background(),
		&services.ListCloudSecurityGroupRuleRequest{
			Spec: &baseservices.ListSpec{
				ParentUUIDs: []string{parentID},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return sgRules.CloudSecurityGroupRules, nil
}
