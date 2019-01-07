package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

var sgNoRuleFQName = []string{defaultDomain, defaultProject, noRuleSecurityGroup}

const (
	projectFieldChildrenSecurityGroups = "security_groups"
	anySecurityGroup                   = "any"
	defaultSecurityGroup               = "default"
	localSecurityGroup                 = "local"
	defaultProject                     = "default-project"
	defaultDomain                      = "default-domain"
	noRuleSecurityGroup                = "__no_rule__"
	egressTrafficContrail              = ">"
	anyProtocol                        = "any"
	ethertypeIPv4                      = "IPv4"
	ethertypeIPv6                      = "IPv6"
	zeroMaskIPv4                       = "0.0.0.0"
	zeroMaskIPv6                       = "::"
	egressTrafficNeutron               = "egress"
	ingressTrafficNeutron              = "ingress"
)

// Create security group logic.
func (sg *SecurityGroup) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	contrailSecurityGroup, err := sg.securityGroupNeutronToContrailOnCreate(ctx, rp)
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "security_group",
			"msg": fmt.Sprintf("Error while casting security group from neutorn to contrail resource."+
				" Error details: %+v", err),
		})
	}

	if sg.Name == defaultSecurityGroup {
		ensureDefaultSecurityGroupExists(ctx, rp)
		return nil, newNeutronError(securityGroupAlreadyExists, errorFields{
			"id": sg.ID,
		})
		// TODO: write unit test to it.
	}

	if err := writeContrailSecurityGroup(ctx, rp, contrailSecurityGroup); err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "security_group",
			"msg":      fmt.Sprintf("Error while saving. Error details: %+v", err),
		})
	}

	// TODO: implement rest of the logic. I ended on neutron_plugin_db.py:4595.

	return securityGroupContrailToNeutron(contrailSecurityGroup)
}

// Update security group logic.
func (sg *SecurityGroup) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
	// TODO implement it.
}

// Delete security group logic.
func (sg *SecurityGroup) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
	// TODO implement it.
}

// Read security group logic.
func (sg *SecurityGroup) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	resp, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		if errutil.IsNotFound(err) {
			err = newNeutronError(securityGroupNotFound, errorFields{
				"id": id,
			})
		} else {
			err = newNeutronError(badRequest, errorFields{
				"resource": "security_group",
				"msg":      fmt.Sprintf("Error while reading security group from database. Error details: %+v", err),
			})
		}
		return nil, err
	}
	return securityGroupContrailToNeutron(resp.SecurityGroup)
}

// ReadAll security group logic.
func (sg *SecurityGroup) ReadAll(ctx context.Context, rp RequestParameters, f Filters, fields Fields) (
	Response, error,
) {
	err := ensureDefaultSecurityGroupExists(ctx, rp)
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "security_group",
			"msg": fmt.Sprintf("Error while checking / creating Default security group. "+
				"Error message: \n%+v", err),
		})
	}

	sgs, err := getListSecurityGroupsFromDB(ctx, rp, f)
	if err != nil {
		return nil, newNeutronError(securityGroupNotFound, errorFields{
			"resource": "security_group",
			"msg": fmt.Sprintf("Error while reading security group from database. "+
				"Error message: \n%+v", err),
		})
	}

	response, err := getNeutronSecurityGroups(sgs, f)

	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "security_group",
			"msg": fmt.Sprintf("Error while converting security group from contrail to neutron resource. "+
				"Error message: \n%+v", err),
		})
	}

	return response, nil
}

func ensureDefaultSecurityGroupExists(ctx context.Context, rp RequestParameters) error {

	project, err := getProject(ctx, rp)
	if err != nil {
		// TODO: think about wrapping this error.
		return err
	}

	for _, sg := range project.GetSecurityGroups() {
		if l := len(sg.GetFQName()); l > 0 && sg.GetFQName()[len(sg.GetFQName())-1] == defaultSecurityGroup {
			return nil
		}
	}

	return createDefaultSecurityGroup(ctx, rp, project)
}

func createDefaultSecurityGroup(
	ctx context.Context, rp RequestParameters, project *models.Project) error {
	projectFQNameString := basemodels.FQNameToString(project.GetFQName())
	sg := models.SecurityGroup{
		Name:       defaultSecurityGroup,
		ParentUUID: project.GetUUID(),
		ParentType: models.KindProject,
		IDPerms: &models.IdPermsType{
			Enable:      true,
			Description: "Default security group",
		},
		SecurityGroupEntries: &models.PolicyEntriesType{
			PolicyRule: []*models.PolicyRuleType{
				createRule(true, defaultSecurityGroup, "", ethertypeIPv4, projectFQNameString),
				createRule(true, defaultSecurityGroup, "", ethertypeIPv6, projectFQNameString),
				createRule(false, "", zeroMaskIPv4, ethertypeIPv4, projectFQNameString),
				createRule(false, "", zeroMaskIPv6, ethertypeIPv6, projectFQNameString),
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

func createRule(ingress bool, securityGroup string, prefix string, ethertype string, projectFQNameString string,
) *models.PolicyRuleType {

	uuid := uuid.NewV4().String()
	localAddr := models.AddressType{
		SecurityGroup: localSecurityGroup,
	}

	var addr models.AddressType
	if securityGroup != "" {
		addr = models.AddressType{
			SecurityGroup: strings.Join([]string{projectFQNameString, securityGroup}, ":"),
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
		Direction: egressTrafficContrail,
		Protocol:  anyProtocol,
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

func getListSecurityGroupsFromDB(
	ctx context.Context, rp RequestParameters, f Filters,
) ([]*models.SecurityGroup, error) {
	var securityGroupUUIDS []string
	var securityGroupTenantUUID []string

	if ids, ok := f["id"]; ok {
		securityGroupUUIDS = ids
	}

	if !rp.RequestContext.IsAdmin {
		securityGroupTenantUUID = []string{rp.RequestContext.Tenant}
	}

	if projectIDs, ok := f["tenant_id"]; ok {
		securityGroupTenantUUID = projectIDs
	}

	return listSecurityGroups(ctx, rp, securityGroupUUIDS, securityGroupTenantUUID)
}

func listSecurityGroups(
	ctx context.Context, rp RequestParameters, uuids []string, tenantUUIDs []string,
) ([]*models.SecurityGroup, error) {

	var parentUUIDs []string
	for _, uuid := range tenantUUIDs {
		contrailUUID, err := neutronIDToContrailUUID(uuid)
		parentUUIDs = append(parentUUIDs, contrailUUID)
		if err != nil {
			return nil, err
		}
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

func getNeutronSecurityGroups(sgs []*models.SecurityGroup, f Filters) ([]*SecurityGroupResponse, error) {

	neutronSGs := make([]*SecurityGroupResponse, 0)
	for _, sg := range sgs {
		if basemodels.FQNameEquals(sg.GetFQName(), sgNoRuleFQName) {
			continue
		}

		if !f.checkValue("name", getSecurityGroupName(sg)) {
			continue
		}

		neutronSG, err := securityGroupContrailToNeutron(sg)
		if err != nil {
			return nil, err
		}

		if neutronSG != nil {
			neutronSGs = append(neutronSGs, neutronSG)
		}
	}

	return neutronSGs, nil
}

func getSecurityGroupName(sg *models.SecurityGroup) string {
	if sg.GetDisplayName() != "" {
		return sg.GetDisplayName()
	}

	return sg.GetName()
}

func securityGroupContrailToNeutron(sg *models.SecurityGroup) (*SecurityGroupResponse, error) {

	sgNeutron := SecurityGroupResponse{
		ID:          sg.GetUUID(),
		TenantID:    contrailUUIDToNeutronID(sg.GetParentUUID()),
		CreatedAt:   sg.GetIDPerms().GetCreated(),
		UpdatedAt:   sg.GetIDPerms().GetLastModified(),
		Description: sg.GetIDPerms().GetDescription(),
		Name:        sg.GetDisplayName(),
		FQName:      sg.GetFQName(),
	}

	if sgNeutron.Name == "" {
		fqName := sg.GetFQName()
		if len(fqName) >= 1 {
			sgNeutron.Name = fqName[len(fqName)-1]
		}
	}

	sgr, err := readSecurityGroupRules(sg)
	if err != nil {
		return nil, err
	}
	sgNeutron.SecurityGroupRules = sgr

	// TODO: Implement 'if loop' of FQName parameter assignation when contrail extension will be available.
	return &sgNeutron, nil
}

func readSecurityGroupRules(sg *models.SecurityGroup) (rules []*SecurityGroupRuleResponse, err error) {
	for _, rule := range sg.GetSecurityGroupEntries().GetPolicyRule() {
		sgr, err := securityGroupRuleContrailToNeutron(sg, rule)
		if err != nil {
			return nil, err
		}
		rules = append(rules, sgr)
	}

	return rules, nil
}

func getProject(ctx context.Context, rp RequestParameters) (*models.Project, error) {
	projectID, err := neutronIDToContrailUUID(rp.RequestContext.TenantID)
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "project",
			"msg":      fmt.Sprintf("Invalid tenant_id parameter. Error message: \n%+v", err),
		})
	}
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
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "project",
			"msg": fmt.Sprintf("Error while reading project of id %s from the DB. "+
				"Error message: \n%+v", projectID, err),
		})
	}

	return projectResponse.GetProject(), nil
}

func (sg *SecurityGroup) securityGroupNeutronToContrailOnCreate(ctx context.Context, rp RequestParameters) (
	*models.SecurityGroup, error) {
	project, err := getProject(ctx, rp)
	if err != nil {
		// TODO: think about wrapping this error.
		return nil, err
	}

	idPerms := models.IdPermsType{
		Enable:      true,
		Description: sg.Description, // TODO: Ask for get method. (GetDescription()). Maybe it is already developed.
		// If exists then refactor all neutron SecuirtyGroup properties getters.
	}
	// TODO: maybe we should save it using Write Service. But remember doing it in transaction!

	sgUUID, err := neutronIDToContrailUUID(sg.ID)

	if err != nil {
		return nil, err // TODO: wrap this error.
	}

	if sgUUID == "" {
		sgUUID = uuid.NewV4().String()
	}

	sgContrail := models.SecurityGroup{
		Name:       sg.Name,
		IDPerms:    &idPerms,
		ParentUUID: project.GetUUID(),
		ParentType: models.KindProject,
		UUID:       sgUUID,
	}

	return &sgContrail, nil
}

func writeContrailSecurityGroup(ctx context.Context, rp RequestParameters, contrailSG *models.SecurityGroup) error {
	_, err := rp.WriteService.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
		SecurityGroup: contrailSG,
	})
	return err
}
