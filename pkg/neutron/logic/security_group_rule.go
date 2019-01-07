package logic

import (
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"strconv"
)

const (
	// TODO: think about moving out all constants into one file.
	//       Because there is problem with searching for already defined constants.
	emptyFulladressIPv4 = "0.0.0.0/0"
	emptyFulladressIPv6 = "::/0"
	defaultPortMin      = 0
	defaultPortMax      = 65535

	PROTO_NAME_ICMP = "icmp"
	PROTO_NUM_ICMP  = "1"
)

func getGenericDefaultSecurityGroupRule() *SecurityGroupRule {
	return &SecurityGroupRule{
		PortRangeMin: defaultPortMin,
		PortRangeMax: defaultPortMax,
		Direction:    egressTrafficNeutron,
		Protocol:     anyProtocol,
	}
}

func getDefaultSecurityGroupRuleIPv4() *SecurityGroupRule {
	sgr := getGenericDefaultSecurityGroupRule()
	sgr.Ethertype = ethertypeIPv4
	sgr.RemoteIPPrefix = emptyFulladressIPv4
	return sgr
}

func getDefaultSecurityGroupRuleIPv6() *SecurityGroupRule {
	sgr := getGenericDefaultSecurityGroupRule()
	sgr.Ethertype = ethertypeIPv6
	sgr.RemoteIPPrefix = emptyFulladressIPv6
	return sgr
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
		PortRangeMin:    defaultPortMin,
		PortRangeMax:    defaultPortMax,
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

func (sgr *SecurityGroupRule) securityGroupRuleNeutronToContrail() (*models.PolicyRuleType, err) {
	//contrailRule := &models.PolicyRuleType{
	//
	//} TODO: delete this if won't be unnecessary

	// TODO: maybe move this block into separate method
	portType := &models.PortType{}
	if !(sgr.Protocol == PROTO_NAME_ICMP || sgr.Protocol == PROTO_NUM_ICMP) {
		portType.StartPort = defaultPortMin
		portType.EndPort = defaultPortMax
	}
	if sgr.PortRangeMin > 0 {
		portType.StartPort = sgr.PortRangeMin
	}
	if sgr.PortRangeMax > 0 {
		portType.EndPort = sgr.PortRangeMax
	}

	if sgr.RemoteIPPrefix != "" && sgr.RemoteGroupID != "" {
		return nil, newNeutronError(securityGroupRemoteGroupAndRemoteIpPrefix, nil)
	}

	// TODO: Continue work on this method. Python code: neutron_plugin_db.py:1324

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
