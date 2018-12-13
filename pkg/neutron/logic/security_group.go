package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

var sgNoRuleFQName = []string{"default-domain", "default-project", "__no_rule__"}

const (
	projectFieldChildrenSecurityGroups = "security_groups"
	defaultSGName                      = "default"
)

// ReadAll logic
func (sg *SecurityGroup) ReadAll(rp RequestParameters, f Filters, fields Fields) (Response, error) {

	//TODO ensure default security group exists
	ctx := context.Background()

	err := ensureDefaultSecurityGroupExists(ctx, rp)
	if err != nil {
		return nil, err
	}

	sgs, err := getSecurityGroupsFromDB(ctx, rp, f)
	if err != nil {
		return nil, err
	}

	return getFixedSecurityGroups(ctx, sgs, f), nil
}

func ensureDefaultSecurityGroupExists(ctx context.Context, rp RequestParameters) error {
	projectID := neutronIDToContrailUUID(rp.RequestContext.TenantID)
	projectResponse, err := rp.ReadService.GetProject(
		ctx,
		&services.GetProjectRequest{
			ID: projectID,
			Fields: []string{
				projectFieldChildrenSecurityGroups,
			},
		},
	)
	if err != nil {
		return err
	}

	for _, sg := range projectResponse.GetProject().GetSecurityGroups() {
		if sg.GetFQName()[len(sg.GetFQName())-1] == defaultSGName {
			return nil
		}
	}

	return createDefaultSecurityGroup()
}

func createDefaultSecurityGroup() error {
	//TODO add logic to this
	return nil
}

func getSecurityGroupsFromDB(
	ctx context.Context, rp RequestParameters, f Filters,
) ([]*models.SecurityGroup, error) {
	if ids, ok := f["id"]; ok {
		return listSecurityGroups(ctx, rp, ids, nil)
	}

	if !rp.RequestContext.IsAdmin {
		return listSecurityGroups(ctx, rp, nil, []string{rp.RequestContext.Tenant})

	}

	if projectIDs, ok := f["tenant_id"]; ok {
		return listSecurityGroups(ctx, rp, nil, projectIDs)
	}

	return listSecurityGroups(ctx, rp, nil, nil)
}

func listSecurityGroups(
	ctx context.Context, rp RequestParameters, uuids []string, tenantUUIDs []string,
) ([]*models.SecurityGroup, error) {

	var parentUUIDs []string
	for _, uuid := range tenantUUIDs {
		parentUUIDs = append(parentUUIDs, neutronIDToContrailUUID(uuid))
	}

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

func getFixedSecurityGroups(
	ctx context.Context, sgs []*models.SecurityGroup, f Filters,
) []*SecurityGroup {

	var neutronSGs []*SecurityGroup
	for _, sg := range sgs {
		if basemodels.FQNameEquals(sg.GetFQName(), sgNoRuleFQName) {
			continue
		}

		if !isPresentInFilters(f, "name", getSecurityGroupName(sg)) {
			continue
		}

		neutronSG := securityGroupVNCToNeutron(ctx, sg)
		if neutronSG != nil {
			neutronSGs = append(neutronSGs, neutronSG)
		}
	}

	return neutronSGs
}

func getSecurityGroupName(sg *models.SecurityGroup) string {
	if sg.GetDisplayName() != "" {
		return sg.GetDisplayName()
	}

	return sg.GetName()
}

//TODO delete that when implemented by Pawel
func securityGroupVNCToNeutron(ctx context.Context, sg *models.SecurityGroup) *SecurityGroup {
	return nil
}
