package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

var sgNoRuleFQName = []string{defaultDomain, defaultProject, noRuleSecurityGroup}

const (
	projectFieldChildrenSecurityGroups = "security_groups"
	defaultSGName                      = "default"
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

	sgs, err := getSecurityGroupsFromDB(ctx, rp, f)
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
	projectID, err := neutronIDToContrailUUID(rp.RequestContext.TenantID)
	if err != nil {
		return newNeutronError(badRequest, errorFields{
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
		return newNeutronError(badRequest, errorFields{
			"resource": "project",
			"msg": fmt.Sprintf("Error while reading project of id %s from the DB. "+
				"Error message: \n%+v", projectID, err),
		})
	}

	project := projectResponse.GetProject()
	for _, sg := range project.GetSecurityGroups() {

		if l := len(sg.GetFQName()); l > 0 && sg.GetFQName()[l-1] == defaultSGName {
			return nil
		}
	}

	return createDefaultSecurityGroup(ctx, rp, project)
}

func createDefaultSecurityGroup(
	ctx context.Context, rp RequestParameters, project *models.Project) error {
	projectFQNameString := basemodels.FQNameToString(project.GetFQName())
	sg := models.SecurityGroup{
		Name:       defaultSGName,
		ParentUUID: project.GetUUID(),
		ParentType: models.KindProject,
		IDPerms: &models.IdPermsType{
			Enable:      true,
			Description: "Default security group",
		},
		SecurityGroupEntries: &models.PolicyEntriesType{
			PolicyRule: []*models.PolicyRuleType{
				createRule(true, defaultSGName, "", ethertypeIPv4, projectFQNameString),
				createRule(true, defaultSGName, "", ethertypeIPv6, projectFQNameString),
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

func getSecurityGroupsFromDB(
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
		if l := len(fqName); l >= 1 {
			sgNeutron.Name = fqName[l-1]
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

func securityGroupRuleContrailToNeutron(
	sg *models.SecurityGroup, rule *models.PolicyRuleType,
) (*SecurityGroupRuleResponse, error) {

	sgr := SecurityGroupRuleResponse{
		ID:              rule.GetRuleUUID(),
		TenantID:        contrailUUIDToNeutronID(sg.GetParentUUID()),
		CreatedAt:       rule.GetCreated(),
		UpdatedAt:       rule.GetLastModified(),
		SecurityGroupID: sg.GetUUID(),
		Ethertype:       rule.GetEthertype(),
		Protocol:        rule.GetProtocol(),
		PortRangeMin:    0,
		PortRangeMax:    65535,
	}

	if err := addressTypeContrailToNeutron(rule, sg, &sgr); err != nil {
		return nil, err
	}

	if len(rule.GetDSTPorts()) != 0 {
		sgr.PortRangeMin = rule.GetDSTPorts()[0].GetStartPort()
		sgr.PortRangeMax = rule.GetDSTPorts()[0].GetEndPort()
	}

	return &sgr, nil
}

func addressTypeContrailToNeutron(
	rule *models.PolicyRuleType,
	sg *models.SecurityGroup,
	sgr *SecurityGroupRuleResponse,
) error {
	var addr *models.AddressType
	srcAddr := rule.GetSRCAddresses()[0]
	dstAddr := rule.GetDSTAddresses()[0]

	if srcAddr.GetSecurityGroup() == localSecurityGroup {
		sgr.Direction = egressTrafficNeutron
		addr = dstAddr
	} else if dstAddr.GetSecurityGroup() == localSecurityGroup {
		sgr.Direction = ingressTrafficNeutron
		addr = srcAddr
	} else {
		return newNeutronError(securityGroupRuleNotFound, errorFields{
			"id": rule.GetRuleUUID(),
		})
	}

	if subnet := addr.GetSubnet(); subnet != nil {
		sgr.RemoteIPPrefix = getFullNetworkAddress(subnet.GetIPPrefix(), subnet.GetIPPrefixLen())
	} else if remoteSG := addr.GetSecurityGroup(); remoteSG != "" && remoteSG != "any" && remoteSG != localSecurityGroup {
		if remoteSG != basemodels.FQNameToString(sg.GetFQName()) {
			// TODO implement it when service FQNameToID will be available in Neutron package.
			// Origin python code: /src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py:1273
		} else {
			sgr.RemoteGroupID = sg.GetUUID()
		}
	}

	return nil
}

func getFullNetworkAddress(ip string, mask int64) string {
	return ip + "/" + strconv.FormatInt(mask, 10)
}
