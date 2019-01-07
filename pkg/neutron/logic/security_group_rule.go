package logic

import (
	"context"
	"fmt"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"strconv"
)

// TODO change function on methods assigned to securityGroupRule

const (
	// TODO: think about moving out constants used by many resources into one file eg. neutron/logic/constants.go.
	neutronSecurityGroupRuleResourceName = "security_group_rule"
	emptyadressIPv4withMask              = "0.0.0.0/0"
	emptyadressIPv6withMask              = "::/0"
	defaultPortMin                       = 0
	defaultPortMax                       = 65535
	PROTO_NAME_ICMP                      = "icmp"
	PROTO_NUM_ICMP                       = "1"
)

func newSecurityGroupRuleError(err error, message string) error {
	if isNeutronError(err) {
		// If that error is already neutron error than do not override it.
		return err
	}

	if err != nil {
		message = fmt.Sprintf(" %+v: %+v", message, err)
	}

	return newNeutronError(badRequest, errorFields{
		"resource": neutronSecurityGroupRuleResourceName,
		"msg":      message,
	})
}

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
	sgr.RemoteIPPrefix = emptyadressIPv4withMask
	return sgr
}

func getDefaultSecurityGroupRuleIPv6() *SecurityGroupRule {
	sgr := getGenericDefaultSecurityGroupRule()
	sgr.Ethertype = ethertypeIPv6
	sgr.RemoteIPPrefix = emptyadressIPv6withMask
	return sgr
}

func (sgr *SecurityGroupRule) securityGroupRuleContrailToNeutronResponse(
	sg *models.SecurityGroup,
	rule *models.PolicyRuleType,
) (*SecurityGroupRuleResponse, error) {

	responseSgr := SecurityGroupRuleResponse{
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

	if err := sgr.addressTypeContrailToNeutron(rule, sg, &responseSgr); err != nil {
		return nil, err
	}

	if len(rule.GetDSTPorts()) != 0 {
		responseSgr.PortRangeMin = rule.GetDSTPorts()[0].GetStartPort()
		responseSgr.PortRangeMax = rule.GetDSTPorts()[0].GetEndPort()
	}

	return &responseSgr, nil
}

func (sgr *SecurityGroupRule) securityGroupRuleNeutronToContrail(ctx context.Context, rp RequestParameters,
) (*models.PolicyRuleType, error) {
	// TODO: think about splitting this method into few smaller ones
	contrailPolicyRule := &models.PolicyRuleType{}

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

	var addrType *models.AddressType
	addrType = &models.AddressType{SecurityGroup: anySecurityGroup}

	if sgr.RemoteIPPrefix != "" {
		etherType := sgr.Ethertype
		ipNetworkVersion, err := getIPVersion(sgr.RemoteIPPrefix)

		if err != nil {
			return nil, newSecurityGroupRuleError(err, "can't determinate ip version of the ip: "+string(ipNetworkVersion))
		}

		if (ipNetworkVersion == 4 && etherType != ethertypeIPv4) ||
			(ipNetworkVersion == 6 && etherType != ethertypeIPv6) {
			return nil, newNeutronError(securityGroupRuleParameterConflict, errorFields{
				"ethertype": etherType,
				"cidr":      sgr.RemoteIPPrefix,
			})
		}

		ipPrefix, ipPrefixLen, err := getIPPrefixAndPrefixLen(sgr.RemoteIPPrefix)

		if err != nil {
			return nil, newSecurityGroupRuleError(err, "can't determinate ip prefix and ip prefix length"+
				" of the cidr: "+sgr.RemoteIPPrefix)
		}

		addrType = &models.AddressType{Subnet: &models.SubnetType{IPPrefix: ipPrefix, IPPrefixLen: ipPrefixLen}}
	}

	if sgr.RemoteGroupID != "" {
		sgResponse, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
			ID: sgr.RemoteGroupID,
		})

		if err != nil {
			return nil, newNeutronError(securityGroupNotFound, errorFields{
				"id": sgr.RemoteGroupID,
			})
		}
		addrType = &models.AddressType{
			SecurityGroup: basemodels.FQNameToString(sgResponse.GetSecurityGroup().GetFQName()),
		}
	}

	if sgr.Direction == ingressTrafficNeutron {
		// TODO: [!] Finished here. Python code: neutron_plugin_db.py:1349
	}

	_ = addrType // TODO: debug only. Delete later.

	return contrailPolicyRule, nil
}

func (sgr *SecurityGroupRule) addressTypeContrailToNeutron(
	rule *models.PolicyRuleType,
	sg *models.SecurityGroup,
	responseSgr *SecurityGroupRuleResponse,
) error {
	var addr *models.AddressType
	srcAddr := rule.GetSRCAddresses()[0]
	dstAddr := rule.GetDSTAddresses()[0]

	if srcAddr.GetSecurityGroup() == localSecurityGroup {
		responseSgr.Direction = egressTrafficNeutron
		addr = dstAddr
	} else if dstAddr.GetSecurityGroup() == localSecurityGroup {
		responseSgr.Direction = ingressTrafficNeutron
		addr = srcAddr
	} else {
		return newNeutronError(securityGroupRuleNotFound, errorFields{
			"id": rule.GetRuleUUID(),
		})
	}

	if subnet := addr.GetSubnet(); subnet != nil {
		responseSgr.RemoteIPPrefix = sgr.getFullNetworkAddress(subnet.GetIPPrefix(), subnet.GetIPPrefixLen())
	} else if remoteSG := addr.GetSecurityGroup(); remoteSG != "" && remoteSG != "any" && remoteSG != localSecurityGroup {
		if remoteSG != basemodels.FQNameToString(sg.GetFQName()) {
			// TODO implement it when service FQNameToID will be available in Neutron package.
			// Origin python code: /src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py:1273
		} else {
			responseSgr.RemoteGroupID = sg.GetUUID()
		}
	}

	return nil
}

func (sgr *SecurityGroupRule) getFullNetworkAddress(ip string, ipLen int64) string {
	return ip + "/" + strconv.FormatInt(ipLen, 10)
}

func securityGroupRuleCreate(
	ctx context.Context,
	rp RequestParameters,
	sg *models.SecurityGroup,
	sgr *models.PolicyRuleType) error {
	rules := sg.GetSecurityGroupEntries()
	if rules == nil {
		rules = &models.PolicyEntriesType{}
	}

	rules.AddPolicyRule(sgr) // TODO: AddPolicyRule not implemented...
	sg.SecurityGroupEntries = rules

	if err := writeContrailSecurityGroup(ctx, rp, sg); err != nil {
		return err // TODO: neutron_plugin_db.py:454 throws some exceptions. Try to imitate them here.
	}
	//TODO: finished here (29.01.2019 17:53)
	return nil
}
