package logic

import (
	"context"
	"fmt"
	"github.com/Juniper/contrail/pkg/errutil"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	securityGroupRuleResourceName = "security_group_rule"
	emptyadressIPv4withMask       = "0.0.0.0/0"
	emptyadressIPv6withMask       = "::/0"
	defaultPortMin                = 0
	defaultPortMax                = 65535
	PROTO_NAME_ICMP               = "icmp"
	PROTO_NUM_ICMP                = "1"
	PROTO_NAME_TCP                = "tcp"
	PROTO_NAME_UDP                = "udp"
	protocolMinValue              = 0
	protocolMaxValue              = 255
)

func getGenericDefaultSecurityGroupRule() *SecurityGroupRule {
	return &SecurityGroupRule{
		PortRangeMin: defaultPortMin,
		PortRangeMax: defaultPortMax,
		Direction:    egressTrafficNeutron,
		Protocol:     protocolAny,
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

func securityGroupRuleContrailToNeutronResponse(
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

	if err := addressTypeContrailToNeutron(rule, sg, &responseSgr); err != nil {
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
	nowISOFormat := time.Now().Format(time.RFC3339)
	contrailSgr := &models.PolicyRuleType{
		RuleUUID:     uuid.NewV4().String(),
		Direction:    egressTrafficContrail,
		SRCPorts:     []*models.PortType{{StartPort: defaultPortMin, EndPort: defaultPortMax}},
		Created:      nowISOFormat,
		LastModified: nowISOFormat,
	}

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
	contrailSgr.DSTPorts = []*models.PortType{portType}

	if sgr.RemoteIPPrefix != "" && sgr.RemoteGroupID != "" {
		return nil, newNeutronError(securityGroupRemoteGroupAndRemoteIpPrefix, nil)
	}

	var addrType *models.AddressType
	addrType = &models.AddressType{SecurityGroup: securityGroupAny}

	if sgr.RemoteIPPrefix != "" {
		etherType := sgr.Ethertype
		ipNetworkVersion, err := getIPVersionFromCIDR(sgr.RemoteIPPrefix)

		if err != nil {
			return nil, errors.Wrapf(err, "can't determinate ip version of the IP: "+string(sgr.RemoteIPPrefix))
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
			return nil, errors.Wrapf(err, "can't determinate ip prefix and ip prefix length"+
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
		contrailSgr.SRCAddresses = []*models.AddressType{addrType}
		contrailSgr.DSTAddresses = []*models.AddressType{{SecurityGroup: securityGroupLocal}}
	} else {
		contrailSgr.SRCAddresses = []*models.AddressType{{SecurityGroup: securityGroupLocal}}
		contrailSgr.DSTAddresses = []*models.AddressType{addrType}
	}

	if err := sgr.parseProtocolParameter(); err != nil {
		return nil, errors.Wrapf(err,
			fmt.Sprintf("invalid protocol value (\"%+v\") in security_group_rule ", sgr.Protocol))
	}
	contrailSgr.Protocol = sgr.Protocol
	if contrailSgr.Protocol == "" {
		contrailSgr.Protocol = protocolAny
	}

	if sgr.Ethertype == "" && sgr.RemoteGroupID == "" && sgr.RemoteIPPrefix == "" {
		contrailSgr.Ethertype = ethertypeIPv4
	} else {
		contrailSgr.Ethertype = sgr.Ethertype
	}

	return contrailSgr, nil
}

func addressTypeContrailToNeutron(
	rule *models.PolicyRuleType,
	sg *models.SecurityGroup,
	responseSgr *SecurityGroupRuleResponse,
) error {
	var addr *models.AddressType
	srcAddr := rule.GetSRCAddresses()[0]
	dstAddr := rule.GetDSTAddresses()[0]

	if srcAddr.GetSecurityGroup() == securityGroupLocal {
		responseSgr.Direction = egressTrafficNeutron
		addr = dstAddr
	} else if dstAddr.GetSecurityGroup() == securityGroupLocal {
		responseSgr.Direction = ingressTrafficNeutron
		addr = srcAddr
	} else {
		return newNeutronError(securityGroupRuleNotFound, errorFields{
			"id": rule.GetRuleUUID(),
		})
	}

	if subnet := addr.GetSubnet(); subnet != nil {
		responseSgr.RemoteIPPrefix = getFullNetworkAddress(subnet.GetIPPrefix(), subnet.GetIPPrefixLen())
	} else if remoteSG := addr.GetSecurityGroup(); remoteSG != "" && remoteSG != "any" && remoteSG != securityGroupLocal {
		if remoteSG != basemodels.FQNameToString(sg.GetFQName()) {
			// TODO implement it when service FQNameToID will be available in Neutron package.
			// Origin python code: /src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py:1273
		} else {
			responseSgr.RemoteGroupID = sg.GetUUID()
		}
	}

	return nil
}

func getFullNetworkAddress(ip string, ipLen int64) string {
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

	rules.PolicyRule = append(rules.PolicyRule, sgr)
	sg.SecurityGroupEntries = rules

	if err := saveContrailSecurityGroup(ctx, rp, sg); err != nil {
		if errutil.IsBadRequest(err) {
			return newNeutronError(badRequest, errorFields{
				"resource": securityGroupRuleResourceName,
				"msg":      fmt.Sprintf("%v", err),
			})
		}
		if errutil.IsQuotaExceeded(err) {
			return newNeutronError(overQuota, errorFields{
				"over": []string{securityGroupRuleResourceName},
				"msg":  fmt.Sprintf("%v", err),
			})
		}
		if errutil.IsConflict(err) {
			return newNeutronError(securityGroupRuleExists, errorFields{
				"resource": securityGroupRuleResourceName,
				"id":       10, // TODO: assign proper rule_id. This value is mocked for now.
				"rule_id":  10, //TODO: assign proper rule_id. This value is mocked for now.
			})
		}
		return errors.Wrapf(err, "can't save %s", securityGroupRuleResourceName)
	}
	return nil
}

func (sgr SecurityGroupRule) parseProtocolParameter() error {
	protos := []string{protocolAny, PROTO_NAME_TCP, PROTO_NAME_UDP, PROTO_NAME_ICMP}

	if protocol, err := strconv.ParseInt(sgr.Protocol, 10, 64); err == nil {
		if protocolMinValue < protocol && protocol < protocolMaxValue {
			return nil
		}
	}

	if isStringInSlice(sgr.Protocol, protos) {
		return nil
	}

	return newNeutronError(securityGroupRuleInvalidProtocol, errorFields{
		"protocol": sgr.Protocol,
		"values":   protos,
	})
}
func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
