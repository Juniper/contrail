package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/types"
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
	securityGroupResourceName          = "security_group"
	projectFieldChildrenSecurityGroups = "security_groups"
	securityGroupAny                   = "any"
	securityGroupDefault               = "default"
	securityGroupLocal                 = "local"
	defaultProject                     = "default-project"
	defaultDomain                      = "default-domain"
	noRuleSecurityGroup                = "__no_rule__"
	egressTrafficVnc                   = ">"
	protocolAny                        = "any"
	ethertypeIPv4                      = "IPv4"
	ethertypeIPv6                      = "IPv6"
	ipv4ZeroValue                      = "0.0.0.0"
	ipv6ZeroValue                      = "::"
	egressTrafficNeutron               = "egress"
	ingressTrafficNeutron              = "ingress"
)

// Create security group logic.
func (sg *SecurityGroup) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	vncSg, err := sg.vncFromNeutron(ctx, rp)
	if err != nil {
		return nil, newSecurityGroupError(err, "error while casting security_group from neutron into vnc resource")
	}

	if sg.Name == securityGroupDefault {
		err = sg.ensureDefaultSecurityGroupExists(ctx, rp)
		if err != nil {
			return nil, newSecurityGroupError(err, "can't ensure default security_group")
		}
		return nil, newNeutronError(securityGroupAlreadyExists, errorFields{
			"id": sg.ID,
		})
	}

	createSgResp, err := rp.WriteService.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{
		SecurityGroup: vncSg,
	})
	if err != nil {
		return nil, newSecurityGroupError(err, "error while saving security_group into database")
	}
	vncSg = createSgResp.GetSecurityGroup()

	if err = assignDefaultSecurityGroupRules(ctx, rp, vncSg); err != nil {
		return nil, newSecurityGroupError(err, "can't create default security group rules")
	}

	resp, err := sg.neutronFromVnc(vncSg)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't cast vnc security_group resource into neutron one")
	}

	return resp, nil
}

// Update security group logic.
func (sg *SecurityGroup) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	sgVnc, err := sg.vncFromNeutron(ctx, rp)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't convert security group from neutron to vnc")
	}
	sgVnc.UUID = id
	sgVnc, err = sg.update(ctx, rp, sgVnc)
	if err != nil {
		return nil, newSecurityGroupError(err, fmt.Sprintf("can't update security group: '%s'", sg.ID))
	}
	resp, err := sg.neutronFromVnc(sgVnc)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't cast contrail security_group resource into neutron one")
	}

	return resp, nil
}

func (sg *SecurityGroup) update(
	ctx context.Context,
	rp RequestParameters,
	sgVnc *models.SecurityGroup,
) (*models.SecurityGroup, error) {
	sgVnc.Name = sg.Name
	sgVnc.DisplayName = sg.Name

	sgVncRes, err := rp.WriteService.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
		SecurityGroup: sgVnc,
		FieldMask: types.FieldMask{
			Paths: []string{models.SecurityGroupFieldName, models.SecurityGroupFieldDisplayName},
		},
	})
	if err != nil {
		return nil, err
	}
	return sgVncRes.GetSecurityGroup(), nil
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
			return nil, newNeutronError(securityGroupNotFound, errorFields{
				"id": id,
			})
		}
		return nil, newSecurityGroupError(err, "error while reading security_group from database")
	}
	return sg.neutronFromVnc(resp.SecurityGroup)
}

// ReadAll security group logic.
func (sg *SecurityGroup) ReadAll(ctx context.Context, rp RequestParameters, f Filters, fields Fields) (
	Response, error,
) {
	err := sg.ensureDefaultSecurityGroupExists(ctx, rp)
	if err != nil {
		return nil, newSecurityGroupError(err, "error while processing default security_group. ")
	}

	sgs, err := sg.getListSecurityGroupsFromDB(ctx, rp, f)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't read security groups from database")
	}

	response, err := sg.neutronFromVncMany(sgs, f)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't cast vnc security_groups into neutron one")
	}

	return response, nil
}

func newSecurityGroupError(err error, message string) error {
	if isNeutronError(err) {
		// If that error is already neutron error than do not override it.
		return err
	}

	if err != nil {
		message = fmt.Sprintf(" %+v: %+v", message, err)
	}

	return newNeutronError(badRequest, errorFields{
		"resource": securityGroupResourceName,
		"msg":      message,
	})
}

func (sg *SecurityGroup) ensureDefaultSecurityGroupExists(ctx context.Context, rp RequestParameters) error {
	project, err := sg.getProject(ctx, rp)
	if err != nil {
		return errors.Wrapf(err, "can't get project")
	}

	for _, sg := range project.GetSecurityGroups() {
		if l := len(sg.GetFQName()); l > 0 && sg.GetFQName()[l-1] == securityGroupDefault {
			return nil
		}
	}

	return sg.createDefaultSecurityGroup(ctx, rp, project)
}

func (sg *SecurityGroup) createDefaultSecurityGroup(
	ctx context.Context, rp RequestParameters, project *models.Project) error {
	projectFQNameString := basemodels.FQNameToString(project.GetFQName())
	vncSg := models.SecurityGroup{
		Name:       securityGroupDefault,
		ParentUUID: project.GetUUID(),
		ParentType: models.KindProject,
		IDPerms: &models.IdPermsType{
			Enable:      true,
			Description: "Default security group",
		},
		SecurityGroupEntries: &models.PolicyEntriesType{
			PolicyRule: []*models.PolicyRuleType{
				sg.createRule(true, securityGroupDefault, "", ethertypeIPv4, projectFQNameString),
				sg.createRule(true, securityGroupDefault, "", ethertypeIPv6, projectFQNameString),
				sg.createRule(false, "", ipv4ZeroValue, ethertypeIPv4, projectFQNameString),
				sg.createRule(false, "", ipv6ZeroValue, ethertypeIPv6, projectFQNameString),
			},
		},
	}

	_, err := rp.WriteService.CreateSecurityGroup(
		ctx,
		&services.CreateSecurityGroupRequest{
			SecurityGroup: &vncSg,
		},
	)

	// TODO chown(sqResponse.GetSecurityGroup().GetUUID(), project.GetUUID())
	return err
}

func (sg *SecurityGroup) createRule(
	ingress bool,
	securityGroup string,
	prefix string,
	ethertype string,
	projectFQNameString string,
) *models.PolicyRuleType {

	uuid := uuid.NewV4().String()
	localAddr := models.AddressType{
		SecurityGroup: securityGroupLocal,
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
		Direction: egressTrafficVnc,
		Protocol:  protocolAny,
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

func (sg *SecurityGroup) getListSecurityGroupsFromDB(
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

	return sg.listSecurityGroups(ctx, rp, securityGroupUUIDS, securityGroupTenantUUID)
}

func (sg *SecurityGroup) listSecurityGroups(
	ctx context.Context, rp RequestParameters, uuids []string, tenantUUIDs []string,
) ([]*models.SecurityGroup, error) {

	var parentUUIDs []string
	for _, uuid := range tenantUUIDs {
		vncUUID, err := neutronIDToVncUUID(uuid)
		parentUUIDs = append(parentUUIDs, vncUUID)
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

func (sg *SecurityGroup) neutronFromVncMany(vncSgs []*models.SecurityGroup, f Filters) (
	[]*SecurityGroupResponse, error) {
	neutronSGs := make([]*SecurityGroupResponse, 0)
	for _, vncSg := range vncSgs {
		if basemodels.FQNameEquals(vncSg.GetFQName(), sgNoRuleFQName) {
			continue
		}

		if !f.Match("name", sg.getSecurityGroupVncName(vncSg)) {
			continue
		}

		neutronSG, err := sg.neutronFromVnc(vncSg)
		if err != nil {
			return nil, err
		}

		if neutronSG != nil {
			neutronSGs = append(neutronSGs, neutronSG)
		}
	}

	return neutronSGs, nil
}

func (sg *SecurityGroup) getSecurityGroupVncName(vncSg *models.SecurityGroup) string {
	if vncSg.GetDisplayName() != "" {
		return vncSg.GetDisplayName()
	}

	return vncSg.GetName()
}

func (sg *SecurityGroup) neutronFromVnc(vncSg *models.SecurityGroup) (*SecurityGroupResponse, error) {

	sgNeutron := SecurityGroupResponse{
		ID:          vncSg.GetUUID(),
		TenantID:    vncUUIDToNeutronID(vncSg.GetParentUUID()),
		CreatedAt:   vncSg.GetIDPerms().GetCreated(),
		UpdatedAt:   vncSg.GetIDPerms().GetLastModified(),
		Description: vncSg.GetIDPerms().GetDescription(),
		Name:        vncSg.GetDisplayName(),
		FQName:      vncSg.GetFQName(),
	}

	if sgNeutron.Name == "" {
		fqName := vncSg.GetFQName()
		if l := len(fqName); l >= 1 {
			sgNeutron.Name = fqName[l-1]
		}
	}

	sgr, err := sg.readSecurityGroupRules(vncSg)
	if err != nil {
		return nil, err
	}
	sgNeutron.SecurityGroupRules = sgr

	// TODO: Implement 'if statement' of FQName parameter assignation when contrail extension will be available.
	return &sgNeutron, nil
}

func (sg *SecurityGroup) readSecurityGroupRules(vncSG *models.SecurityGroup) (
	rules []*SecurityGroupRuleResponse, err error) {
	for _, rule := range vncSG.GetSecurityGroupEntries().GetPolicyRule() {
		sgr, err := (&SecurityGroupRule{}).neutronFromVnc(vncSG, rule)
		if err != nil {
			return nil, err
		}
		rules = append(rules, sgr)
	}

	return rules, nil
}

func (sg *SecurityGroup) getProject(ctx context.Context, rp RequestParameters) (*models.Project, error) {
	projectID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
	if err != nil {
		return nil, newSecurityGroupError(err, "invalid tenant_id parameter")
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
		return nil, newSecurityGroupError(err,
			fmt.Sprintf("can't read project of id '%s' from the database.", projectID))
	}

	return projectResponse.GetProject(), nil
}

func (sg *SecurityGroup) vncFromNeutron(ctx context.Context, rp RequestParameters) (
	*models.SecurityGroup, error) {
	project, err := sg.getProject(ctx, rp)
	if err != nil {
		return nil, err
	}

	idPerms := models.IdPermsType{
		Enable:      true,
		Description: sg.Description,
	}

	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("can't create security_group, bad id: '%s'", sg.ID))
	}

	return &models.SecurityGroup{
		Name:       sg.Name,
		IDPerms:    &idPerms,
		ParentUUID: project.GetUUID(),
		ParentType: models.KindProject,
	}, nil
}

func assignDefaultSecurityGroupRules(
	ctx context.Context,
	rp RequestParameters,
	vncSg *models.SecurityGroup) error {
	vncRuleIPv4, err := getDefaultSecurityGroupRuleIPv4().vncFromNeutron(ctx, rp)
	if err != nil {
		return errors.Wrapf(err, "can't cast security_group_rule IPv4 into vnc resource")
	}
	err = securityGroupRuleCreate(ctx, rp, vncSg, vncRuleIPv4)
	if err != nil {
		return err
	}

	vncRuleIPv6, err := getDefaultSecurityGroupRuleIPv6().vncFromNeutron(ctx, rp)
	if err != nil {
		return errors.Wrapf(err, "can't cast security_group_rule IPv6 into vnc resource")
	}
	if err := securityGroupRuleCreate(ctx, rp, vncSg, vncRuleIPv6); err != nil {
		return err
	}
	return nil
}
