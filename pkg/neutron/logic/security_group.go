package logic

import (
	"context"
	"strconv"
	"strings"

	"github.com/juju/errors"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	localSecurityGroup = "local"
)

func (s *SecurityGroup) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	resp, err := rp.ReadService.GetSecurityGroup(context.Background(), &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		return nil, NewNeutronError(SecurityGroupNotFound, map[string]interface{}{
			"id": id,
		})
	}

	return securityGroupContrailToNeutron(resp.SecurityGroup)
}

func securityGroupContrailToNeutron(sg *models.SecurityGroup) (*SecurityGroupResponse, error) {

	sgNeutron := SecurityGroupResponse{}
	sgNeutron.ID = sg.GetUUID()
	sgNeutron.TenantID = contrailUUIDToNeutronID(sg.GetParentUUID())
	sgNeutron.CreatedAt = sg.GetIDPerms().GetCreated()
	sgNeutron.UpdatedAt = sg.GetIDPerms().GetLastModified()

	if sg.GetDisplayName() != "" {
		sgNeutron.Name = sg.GetDisplayName()
	} else {
		fgName := sg.GetFQName()
		sgNeutron.Name = fgName[len(fgName)-1]
	}

	if sg.IDPerms.GetDescription() == "" {
		sgNeutron.Description = sg.GetIDPerms().GetDescription()
	}

	var err error
	sgNeutron.SecurityGroupRules, err = readSecurityGroupRules(sg)
	if err != nil {
		return nil, err
	}

	// TODO(drapek): Uncomment and adjust when contrail extension support will be done.
	//if globalConfig.contrailExtensionsEnabled {
	//	sgNeutron.FQName = sg.GetFQName()
	//}

	return &sgNeutron, nil
}

func readSecurityGroupRules(sg *models.SecurityGroup) ([]*SecurityGroupRuleResponse, error) {
	sge := sg.GetSecurityGroupEntries()

	var rules []*SecurityGroupRuleResponse
	for _, rule := range sge.GetPolicyRule() {
		sgr, err := securityGroupRuleContrailToNeutron(sg, rule)
		if err != nil {
			return nil, errors.BadRequestf("SecurityGroupRuleNotFound: %v", err)
		}
		rules = append(rules, sgr)
	}

	return rules, nil
}

func securityGroupRuleContrailToNeutron(
	sg *models.SecurityGroup, rule *models.PolicyRuleType,
) (*SecurityGroupRuleResponse, error) {

	var sgr SecurityGroupRuleResponse

	sgr.ID = rule.GetRuleUUID()
	sgr.TenantID = contrailUUIDToNeutronID(sg.GetParentUUID())
	sgr.SecurityGroupID = sg.GetUUID()
	sgr.Ethertype = rule.GetEthertype()
	sgr.Protocol = rule.GetProtocol()
	sgr.CreatedAt = rule.GetCreated()
	sgr.UpdatedAt = rule.GetLastModified()

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
		return nil, errors.BadRequestf("SecurityGroupRuleNotFound with id %s", rule.GetRuleUUID())
	}

	subnet := addr.GetSubnet()
	remoteSG := addr.GetSecurityGroup()
	if subnet != nil {
		sgr.RemoteIPPrefix = strings.Join(
			[]string{subnet.GetIPPrefix(), strconv.FormatInt(subnet.GetIPPrefixLen(), 10)}, "/",
		)
	} else if remoteSG != "" && remoteSG != "any" && remoteSG != localSecurityGroup {
		if remoteSG != basemodels.FQNameToString(sg.GetFQName()) {
			// TODO(drapek) implement it when service FQNameToID will be available in Neutron package
			//sgr.RemoteGroupID, err = FQNameToID(append([]string{"security-group"}, basemodels.ParseFQName(remoteSG)...))
			//if err != nil {
			//	return &sgr, nil
			//}
		} else {
			sgr.RemoteGroupID = sg.GetUUID()
		}
	}

	if rule.GetDSTPorts() != nil {
		sgr.PortRangeMin = rule.GetDSTPorts()[0].GetStartPort()
		sgr.PortRangeMax = rule.GetDSTPorts()[0].GetEndPort()
	} else {
		sgr.PortRangeMin = 0
		sgr.PortRangeMax = 65535
	}

	return &sgr, nil
}
