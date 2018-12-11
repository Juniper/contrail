package logic

import (
	"context"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/juju/errors"
	"strconv"
	"strings"
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

	return securityGroupContrailToNeutron(rp, resp.SecurityGroup)
}

func securityGroupContrailToNeutron(rp RequestParameters, sg *models.SecurityGroup) (*SecurityGroupResponse, error) {
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
			return nil, errors.BadRequestf("SecurityGroupNotFound: %v", err)
		}
		rules = append(rules, sgr)
	}

	return rules, nil
}


func securityGroupRuleContrailToNeutron(
	sg *models.SecurityGroup, rule *models.PolicyRuleType,
	)(*SecurityGroupRuleResponse, error) {

	var sgr SecurityGroupRuleResponse

	var addr *models.AddressType
	srcAddr := rule.GetSRCAddresses()[0]
	dstAddr := rule.GetDSTAddresses()[0]

	if srcAddr.GetSecurityGroup() == "local" {
		sgr.Direction = "egress"
		addr = dstAddr
	} else if dstAddr.GetSecurityGroup() == "local" {
		sgr.Direction = "ingress"
		addr = srcAddr
	} else {
		return nil, errors.BadRequestf("SecurityGroupRuleNotFound with id %s", rule.GetRuleUUID())
	}

	subnet := addr.GetSubnet()
	remoteSG := addr.GetSecurityGroup()
	if subnet != nil {
		sgr.RemoteIPPrefix = strings.Join(
			[]string{subnet.GetIPPrefix(), strconv.FormatInt(subnet.GetIPPrefixLen(), 10}, "/",
		)
	} else if remoteSG != "" && remoteSG != "any" && remoteSG != "local" {
		if remoteSG != basemodels.FQNameToString(sg.GetFQName()) {
			sgr.RemoteGroupID = FQnameToID(append([]string{"security-group"}, basemodels.ParseFQName(remoteSG))) // TODO(drapek): find FQName to ID function in our system, improve joining two slices
		} else {
			sgr.RemoteGroupID = sg.GetUUID()
		}
	}

}
