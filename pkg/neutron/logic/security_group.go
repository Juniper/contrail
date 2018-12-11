package logic

import (
	"context"
	"fmt"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/pkg/errors"
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	localSecurityGroup = "local"
)

func (s *SecurityGroup) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	resp, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
		ID: id,
	})
	if err != nil {
		return nil, NewNeutronError(BadRequest, map[string]interface{}{
			"id": id,
		})
	}

	return securityGroupContrailToNeutron(resp.SecurityGroup)
}

func securityGroupContrailToNeutron(sg *models.SecurityGroup) (*SecurityGroupResponse, error) {

	sgNeutron := SecurityGroupResponse{
		ID:          sg.GetUUID(),
		TenantID:    contrailUUIDToNeutronID(sg.GetParentUUID()),
		CreatedAt:   sg.GetIDPerms().GetCreated(),
		UpdatedAt:   sg.GetIDPerms().GetLastModified(),
		Description: sg.GetIDPerms().GetDescription(),
	}

	if sg.GetDisplayName() != "" {
		sgNeutron.Name = sg.GetDisplayName()
	} else {
		fqName := sg.GetFQName()
		sgNeutron.Name = fqName[len(fqName)-1]
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
	var rules []*SecurityGroupRuleResponse
	for _, rule := range sg.GetSecurityGroupEntries().GetPolicyRule() {
		sgr, err := securityGroupRuleContrailToNeutron(sg, rule)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("SecurityGroupRuleNotFound: %v", err))
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
	}

	if err := castAddressTypeContrailToNeutron(rule, sg, &sgr); err != nil {
		return nil, err
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

func castAddressTypeContrailToNeutron(
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
		return errors.New(fmt.Sprintf("SecurityGroupRuleNotFound with id %s", rule.GetRuleUUID()))
	}

	remoteSG := addr.GetSecurityGroup()
	if subnet := addr.GetSubnet(); subnet != nil {
		sgr.RemoteIPPrefix = strings.Join(
			[]string{subnet.GetIPPrefix(), strconv.FormatInt(subnet.GetIPPrefixLen(), 10)}, "/",
		)
	} else if remoteSG != "" && remoteSG != "any" && remoteSG != localSecurityGroup {
		if remoteSG != basemodels.FQNameToString(sg.GetFQName()) {
			// TODO(drapek) implement it when service FQNameToID will be available in Neutron package.
			//sgr.RemoteGroupID, err = FQNameToID(append([]string{"security-group"}, basemodels.ParseFQName(remoteSG)...))
			//if err != nil {
			//	return &sgr, nil
			//}
		} else {
			sgr.RemoteGroupID = sg.GetUUID()
		}
	}

	return nil
}
