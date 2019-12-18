package logic

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

var sgNoRuleFQName = []string{defaultDomainName, defaultProjectName, noRuleSecurityGroup}

const (
	securityGroupResourceName          = "security_group"
	projectFieldChildrenSecurityGroups = "security_groups"
	noRuleSecurityGroup                = "__no_rule__"
	egressTrafficNeutron               = "egress"
	ingressTrafficNeutron              = "ingress"
)

// Create security group logic.
func (sg *SecurityGroup) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	vncSg, err := sg.neutronToVnc(ctx, rp)
	if err != nil {
		return nil, newSecurityGroupError(err, "error while casting security_group from neutron into vnc resource")
	}

	if sg.Name == models.DefaultSecurityGroupName {
		if err = sg.ensureDefaultSecurityGroupExists(ctx, rp); err != nil {
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
		return nil, newSecurityGroupError(err, "Cannot create security_group resource")
	}
	vncSg = createSgResp.GetSecurityGroup()

	if err = sg.assignDefaultSecurityGroupRules(ctx, rp, vncSg); err != nil {
		return nil, newSecurityGroupError(err, "can't create default security group rules")
	}

	resp, err := sg.vncToNeutron(vncSg)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't cast vnc security_group resource into neutron one")
	}

	return resp, nil
}

// Update security group logic.
func (sg *SecurityGroup) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	sgVnc, err := sg.neutronToVnc(ctx, rp)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't convert security group from neutron to vnc")
	}
	sgVnc.UUID = id
	if err = sg.update(ctx, rp, sgVnc); err != nil {
		return nil, newSecurityGroupError(err, fmt.Sprintf("can't update security group: '%s'", sg.ID))
	}
	response, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{ID: id})
	if err != nil {
		return nil, err
	}
	resp, err := sg.vncToNeutron(response.SecurityGroup)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't cast contrail security_group resource into neutron one")
	}

	return resp, nil
}

func (sg *SecurityGroup) update(ctx context.Context, rp RequestParameters, sgVnc *models.SecurityGroup) error {
	var fm types.FieldMask
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(SecurityGroupFieldDescription)) {
		basemodels.FieldMaskAppend(&fm, models.SecurityGroupFieldIDPerms, models.IdPermsTypeFieldDescription)
	}
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(SecurityGroupFieldName)) {
		basemodels.FieldMaskAppend(&fm, models.SecurityGroupFieldName)
		basemodels.FieldMaskAppend(&fm, models.SecurityGroupFieldDisplayName)
	}

	_, err := rp.WriteService.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
		SecurityGroup: sgVnc,
		FieldMask:     fm,
	})

	return err
}

// Delete security group logic.
func (sg *SecurityGroup) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	resp, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		return nil, newSecurityGroupError(err, "error while reading security_group from database")
	}
	vncSg := resp.GetSecurityGroup()

	projectUUID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
	if err != nil {
		return nil, newSecurityGroupError(err, "invalid tenant id")
	}

	if vncSg.GetName() == models.DefaultSecurityGroupName && vncSg.GetParentUUID() == projectUUID {
		return nil, newNeutronError(securityGroupCannotRemoveDefault, errorFields{})
	}

	_, err = rp.WriteService.DeleteSecurityGroup(ctx, &services.DeleteSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		if errutil.IsConflict(err) {
			return nil, newNeutronError(securityGroupInUse, errorFields{
				"id": id,
			})
		}
		return nil, newSecurityGroupError(err,
			fmt.Sprintf("can't delete resource security_group of id '%s'", id))
	}

	return nil, nil
}

// Read security group logic.
func (sg *SecurityGroup) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	resp, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
		ID: id,
	})
	if errutil.IsNotFound(err) {
		return nil, newNeutronError(securityGroupNotFound, errorFields{
			"id": id,
		})
	}
	if err != nil {
		return nil, newSecurityGroupError(err, "error while reading security_group from database")
	}

	return sg.vncToNeutron(resp.SecurityGroup)
}

// ReadAll security group logic.
func (sg *SecurityGroup) ReadAll(
	ctx context.Context, rp RequestParameters, f Filters, fields Fields,
) (Response, error) {
	if err := sg.ensureDefaultSecurityGroupExists(ctx, rp); err != nil {
		return nil, newSecurityGroupError(err, "error while processing default security_group. ")
	}

	sgs, err := listSecurityGroups(ctx, rp, f)
	if err != nil {
		return nil, newSecurityGroupError(err, "can't read security groups from database")
	}

	response, err := sg.convertVncSecurityGroupsToNeutron(sgs, f)
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
		if l := len(sg.GetFQName()); l > 0 && sg.GetFQName()[l-1] == models.DefaultSecurityGroupName {
			return nil
		}
	}

	return sg.createDefaultSecurityGroup(ctx, rp, project)
}

func (sg *SecurityGroup) createDefaultSecurityGroup(
	ctx context.Context, rp RequestParameters, project *models.Project,
) error {
	_, err := rp.WriteService.CreateSecurityGroup(
		ctx,
		&services.CreateSecurityGroupRequest{
			SecurityGroup: project.DefaultSecurityGroup(),
		},
	)

	// TODO chown(sqResponse.GetSecurityGroup().GetUUID(), project.GetUUID())
	return err
}

func listSecurityGroups(
	ctx context.Context, rp RequestParameters, filter Filters,
) ([]*models.SecurityGroup, error) {
	var parentUUIDs []string
	for _, uuid := range getFilterProjectIDS(ctx, rp, filter) {
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
				ObjectUUIDs: filter[idKey],
				ParentUUIDs: parentUUIDs,
				Detail:      true,
			},
		},
	)

	return sgResponse.GetSecurityGroups(), err
}

func getFilterProjectIDS(ctx context.Context, rp RequestParameters, f Filters) []string {
	if !rp.RequestContext.IsAdmin {
		return []string{rp.RequestContext.Tenant}
	}
	return f[tenantIDKey]
}

func (sg *SecurityGroup) convertVncSecurityGroupsToNeutron(
	vncSgs []*models.SecurityGroup, f Filters,
) ([]*SecurityGroupResponse, error) {
	neutronSGs := make([]*SecurityGroupResponse, 0)
	for _, vncSg := range vncSgs {
		if basemodels.FQNameEquals(vncSg.GetFQName(), sgNoRuleFQName) {
			continue
		}

		if !f.Match(SecurityGroupFieldName, sg.getSecurityGroupName(vncSg)) {
			continue
		}

		neutronSG, err := sg.vncToNeutron(vncSg)
		if err != nil {
			return nil, err
		}

		if neutronSG != nil {
			neutronSGs = append(neutronSGs, neutronSG)
		}
	}

	return neutronSGs, nil
}

func (sg *SecurityGroup) getSecurityGroupName(vncSg *models.SecurityGroup) string {
	if vncSg.GetDisplayName() != "" {
		return vncSg.GetDisplayName()
	}

	return vncSg.GetName()
}

func (sg *SecurityGroup) vncToNeutron(vncSg *models.SecurityGroup) (*SecurityGroupResponse, error) {

	sgNeutron := SecurityGroupResponse{
		ID:          vncSg.GetUUID(),
		TenantID:    VncUUIDToNeutronID(vncSg.GetParentUUID()),
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

func (sg *SecurityGroup) readSecurityGroupRules(
	vncSG *models.SecurityGroup,
) (rules []*SecurityGroupRuleResponse, err error) {
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

func (sg *SecurityGroup) neutronToVnc(ctx context.Context, rp RequestParameters) (*models.SecurityGroup, error) {
	project, err := sg.getProject(ctx, rp)
	if err != nil {
		return nil, err
	}

	idPerms := models.IdPermsType{
		Enable:      true,
		Description: sg.Description,
	}

	return &models.SecurityGroup{
		UUID:        sg.ID,
		Name:        sg.Name,
		DisplayName: sg.Name,
		IDPerms:     &idPerms,
		ParentUUID:  project.GetUUID(),
		ParentType:  models.KindProject,
	}, nil
}

func (SecurityGroup) assignDefaultSecurityGroupRules(
	ctx context.Context, rp RequestParameters, vncSg *models.SecurityGroup,
) error {
	vncRuleIPv4, err := getDefaultSecurityGroupRuleIPv4().vncFromNeutron(ctx, rp)
	if err != nil {
		return errors.Wrapf(err, "can't cast security_group_rule IPv4 into vnc resource")
	}
	if err = securityGroupRuleCreate(ctx, rp, vncSg, vncRuleIPv4); err != nil {
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
