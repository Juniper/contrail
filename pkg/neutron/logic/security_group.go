package logic

import (
	"context"
	"strings"

	uuid "github.com/satori/go.uuid"

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

	project := projectResponse.GetProject()
	for _, sg := range project.GetSecurityGroups() {
		if sg.GetFQName()[len(sg.GetFQName())-1] == defaultSGName {
			return nil
		}
	}

	return createDefaultSecurityGroup(ctx, rp, project)
}

func createDefaultSecurityGroup(
	ctx context.Context, rp RequestParameters, project *models.Project) error {
	createRule := func(ingress bool, sg string, prefix string, ethertype string) *models.PolicyRuleType {
		uuid := uuid.NewV4().String()
		localAddr := models.AddressType{
			SecurityGroup: "local",
		}

		var addr models.AddressType
		if sg != "" {
			fqName := basemodels.FQNameToString(project.GetFQName())
			addr = models.AddressType{
				SecurityGroup: strings.Join([]string{fqName, sg}, ":"),
			}
		}

		if prefix != "" {
			addr = models.AddressType{
				Subnet: &models.SubnetType{
					IPPrefix:    prefix,
					IPPrefixLen: 0,
				},
			}
		}

		rule := models.PolicyRuleType{
			RuleUUID:  uuid,
			Direction: ">",
			Protocol:  "any",
			SRCPorts: []*models.PortType{
				{
					StartPort: 0,
					EndPort:   65535,
				},
			},
			DSTPorts: []*models.PortType{
				{
					StartPort: 0,
					EndPort:   65535,
				},
			},
			Ethertype: ethertype,
		}
		if ingress {
			rule.SRCAddresses = []*models.AddressType{&addr}
			rule.DSTAddresses = []*models.AddressType{&localAddr}
			return &rule
		}

		rule.SRCAddresses = []*models.AddressType{&localAddr}
		rule.DSTAddresses = []*models.AddressType{&addr}
		return &rule
	}

	sg := models.SecurityGroup{
		Name:       "default",
		ParentUUID: project.GetUUID(),
		ParentType: models.KindProject,
		IDPerms: &models.IdPermsType{
			Enable:      true,
			Description: "Default security group",
		},
		SecurityGroupEntries: &models.PolicyEntriesType{
			PolicyRule: []*models.PolicyRuleType{
				createRule(true, "default", "", "IPv4"),
				createRule(true, "default", "", "IPv6"),
				createRule(false, "", "0.0.0.0", "IPv4"),
				createRule(false, "", "::", "IPv6"),
			},
		},
	}

	_, err := rp.WriteService.CreateSecurityGroup(
		ctx,
		&services.CreateSecurityGroupRequest{
			SecurityGroup: &sg,
		},
	)

	// TODO chown(sqResponse.GetSecurityGroup().GetUUID(), project.GetUUID())
	return err
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
