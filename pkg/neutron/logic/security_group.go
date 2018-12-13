package logic

import (
	"context"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

var sgNoRuleFQName = []string{"default-domain", "default-project", "__no_rule__"}

const (
	projectFieldChildrenSecurityGroups = "security_groups"
	defaultSGName                      = "default"
	localSecurityGroup                 = "local"
)

// Read security group logic.
func (sg *SecurityGroup) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	resp, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		if errutil.IsNotFound(err) {
			err = NewNeutronError(SecurityGroupNotFound, map[string]interface{}{
				"id": id,
			})
		} else {
			err = NewNeutronError(BadRequest, nil)
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
		return nil, err
	}

	sgs, err := getSecurityGroupsFromDB(ctx, rp, f)
	if err != nil {
		return nil, err
	}

	response, err := getFixedSecurityGroups(sgs, f)

	if err != nil {
		return nil, NewNeutronError(BadRequest, nil) // TODO: make sure that it will be sufficient.
	}

	return response, nil
}

func ensureDefaultSecurityGroupExists(ctx context.Context, rp RequestParameters) error {
	projectID, err := neutronIDToContrailUUID(rp.RequestContext.TenantID)
	if err != nil {
		return NewNeutronError(ProjectNotFound, map[string]interface{}{
			"resource": "project",
			"msg":      "Invalid tenant_id parameter.",
			"err":      err.Error(),
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
		return NewNeutronError(ProjectNotFound, map[string]interface{}{
			"id": projectID,
		})
	}

	project := projectResponse.GetProject()
	for _, sg := range project.GetSecurityGroups() {
		if len(sg.GetFQName()) > 0 && sg.GetFQName()[len(sg.GetFQName())-1] == defaultSGName {
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
		} else if prefix != "" {
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
		neutronUUID, err := neutronIDToContrailUUID(uuid)
		parentUUIDs = append(parentUUIDs, neutronUUID)
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

func getFixedSecurityGroups(sgs []*models.SecurityGroup, f Filters) ([]*SecurityGroupResponse, error) {

	var neutronSGs []*SecurityGroupResponse
	for _, sg := range sgs {
		if basemodels.FQNameEquals(sg.GetFQName(), sgNoRuleFQName) {
			continue
		}

		if !isPresentInFilters(f, "name", getSecurityGroupName(sg)) {
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
	sgNeutron.SecurityGroupRules = sgr
	if err != nil {
		return nil, err
	}

	// TODO: Implement 'if loop' when contrail extension will be available.
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
		sgr.Direction = "egress"
		addr = dstAddr
	} else if dstAddr.GetSecurityGroup() == localSecurityGroup {
		sgr.Direction = "ingress"
		addr = srcAddr
	} else {
		return NewNeutronError("SecurityGroupRuleNotFound", map[string]interface{}{
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

func getFullNetworkAddress(ip string, cidr int64) string {
	return ip + "/" + strconv.FormatInt(cidr, 10)
}
