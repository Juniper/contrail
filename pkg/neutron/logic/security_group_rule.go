package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	securityGroupRuleResourceName = "security_group_rule"
	maskZeroValue                 = "/0"
	defaultPortMin                = 0
	defaultPortMax                = 65535
	protoNumICMP                  = "1"
	protocolMinValue              = 0
	protocolMaxValue              = 255
)

// ReadAll security group rule logic.
func (sgr *SecurityGroupRule) ReadAll(
	ctx context.Context, rp RequestParameters, f Filters, fields Fields,
) (Response, error) {
	// TODO: implement it when Read operation will be available
	sgs := sgr.collectSecurityGroups(f)

	// TODO: for every seg group collect sg rules by sg.readSecurityGroupRules() make sure that there are implemented
	//       functions for OPER=LIST

	// TODO: Return it as list

	_ = sgs // TODO: debug only to fool the compiler
	return nil, nil
}

func (sgr *SecurityGroupRule) collectSecurityGroups(filters Filters) []*SecurityGroup {
	if len(filters) > 0 && filters.HaveKeys(SecurityGroupRuleFieldTenantID) {
		// TODO: if 'tenant_id' in filters than collect all sg withing projects ids
	}

	// TODO: rest of python code: neutron_plugin_db.py:4884 with no filters
	return nil
}

func getGenericDefaultSecurityGroupRule() *SecurityGroupRule {
	return &SecurityGroupRule{
		PortRangeMin: defaultPortMin,
		PortRangeMax: defaultPortMax,
		Direction:    egressTrafficNeutron,
		Protocol:     models.AnyProtocol,
	}
}

func getDefaultSecurityGroupRuleIPv4() *SecurityGroupRule {
	sgr := getGenericDefaultSecurityGroupRule()
	sgr.Ethertype = models.IPv4Ethertype
	sgr.RemoteIPPrefix = models.IPv4ZeroValue + maskZeroValue
	return sgr
}

func getDefaultSecurityGroupRuleIPv6() *SecurityGroupRule {
	sgr := getGenericDefaultSecurityGroupRule()
	sgr.Ethertype = models.IPv6Ethertype
	sgr.RemoteIPPrefix = models.IPv6ZeroValue + maskZeroValue
	return sgr
}

func (*SecurityGroupRule) neutronFromVnc(
	sg *models.SecurityGroup, rule *models.PolicyRuleType,
) (*SecurityGroupRuleResponse, error) {
	responseSgr := SecurityGroupRuleResponse{
		ID:              rule.GetRuleUUID(),
		TenantID:        VncUUIDToNeutronID(sg.GetParentUUID()),
		CreatedAt:       rule.GetCreated(),
		UpdatedAt:       rule.GetLastModified(),
		SecurityGroupID: sg.GetUUID(),
		Ethertype:       rule.GetEthertype(),
		Protocol:        rule.GetProtocol(),
		PortRangeMin:    defaultPortMin,
		PortRangeMax:    defaultPortMax,
	}

	if err := addressTypeNeutronFromVnc(rule, sg, &responseSgr); err != nil {
		return nil, err
	}

	if len(rule.GetDSTPorts()) != 0 {
		responseSgr.PortRangeMin = rule.GetDSTPorts()[0].GetStartPort()
		responseSgr.PortRangeMax = rule.GetDSTPorts()[0].GetEndPort()
	}

	return &responseSgr, nil
}

func (sgr *SecurityGroupRule) vncFromNeutron(
	ctx context.Context, rp RequestParameters,
) (*models.PolicyRuleType, error) {
	nowISOFormat := time.Now().Format(constants.ISO8601TimeFormat)
	vncSgr := &models.PolicyRuleType{
		RuleUUID:     uuid.NewV4().String(),
		Direction:    models.SRCToDSTDirection,
		SRCPorts:     []*models.PortType{{StartPort: defaultPortMin, EndPort: defaultPortMax}},
		Created:      nowISOFormat,
		LastModified: nowISOFormat,
	}

	vncSgr.DSTPorts = []*models.PortType{sgr.getPortType()}

	if sgr.RemoteIPPrefix != "" && sgr.RemoteGroupID != "" {
		return nil, newNeutronError(securityGroupRemoteGroupAndRemoteIPPrefix, nil)
	}

	if err := sgr.initAddressType(ctx, rp, vncSgr); err != nil {
		return nil, err
	}

	if err := sgr.parseProtocolParameter(); err != nil {
		return nil, errors.Wrapf(err,
			fmt.Sprintf("invalid protocol value ('%+v') in security_group_rule ", sgr.Protocol))
	}
	vncSgr.Protocol = sgr.Protocol
	if vncSgr.Protocol == "" {
		vncSgr.Protocol = models.AnyProtocol
	}

	if sgr.Ethertype == "" && sgr.RemoteGroupID == "" && sgr.RemoteIPPrefix == "" {
		vncSgr.Ethertype = models.IPv4Ethertype
	} else {
		vncSgr.Ethertype = sgr.Ethertype
	}

	return vncSgr, nil
}

func (sgr *SecurityGroupRule) getPortType() *models.PortType {
	portType := &models.PortType{}
	if sgr.Protocol != models.ICMPProtocol && sgr.Protocol != protoNumICMP {
		portType.StartPort = defaultPortMin
		portType.EndPort = defaultPortMax
	}
	if sgr.PortRangeMin > 0 {
		portType.StartPort = sgr.PortRangeMin
	}
	if sgr.PortRangeMax > 0 {
		portType.EndPort = sgr.PortRangeMax
	}
	return portType
}

func (sgr *SecurityGroupRule) initAddressType(
	ctx context.Context, rp RequestParameters, vncSgr *models.PolicyRuleType,
) error {
	var addrType *models.AddressType
	addrType = &models.AddressType{SecurityGroup: models.AnySecurityGroup}

	if sgr.RemoteIPPrefix != "" {
		ipPrefix, ipPrefixLen, err := sgr.getIPPrefixWithLen()
		if err != nil {
			return err
		}
		addrType = &models.AddressType{Subnet: &models.SubnetType{IPPrefix: ipPrefix, IPPrefixLen: ipPrefixLen}}
	}

	if sgr.RemoteGroupID != "" {
		sgResponse, err := rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
			ID: sgr.RemoteGroupID,
		})

		if err != nil {
			return newNeutronError(securityGroupNotFound, errorFields{
				"id": sgr.RemoteGroupID,
			})
		}
		addrType = &models.AddressType{
			SecurityGroup: basemodels.FQNameToString(sgResponse.GetSecurityGroup().GetFQName()),
		}
	}

	if sgr.Direction == ingressTrafficNeutron {
		vncSgr.SRCAddresses = []*models.AddressType{addrType}
		vncSgr.DSTAddresses = []*models.AddressType{{SecurityGroup: models.LocalSecurityGroup}}
	} else {
		vncSgr.SRCAddresses = []*models.AddressType{{SecurityGroup: models.LocalSecurityGroup}}
		vncSgr.DSTAddresses = []*models.AddressType{addrType}
	}

	return nil
}

func (sgr *SecurityGroupRule) getIPPrefixWithLen() (string, int64, error) {
	etherType := sgr.Ethertype
	ipNetworkVersion, err := getIPVersionFromCIDR(sgr.RemoteIPPrefix)

	if err != nil {
		return "", 0, errors.Wrapf(err, "can't determinate ip version of the IP: '%s'", sgr.RemoteIPPrefix)
	}

	if (ipNetworkVersion == 4 && etherType != models.IPv4Ethertype) ||
		(ipNetworkVersion == 6 && etherType != models.IPv6Ethertype) {
		return "", 0, newNeutronError(securityGroupRuleParameterConflict, errorFields{
			"ethertype": etherType,
			"cidr":      sgr.RemoteIPPrefix,
		})
	}

	ipPrefix, _, ipLen, err := getIPPrefixAndPrefixLen(sgr.RemoteIPPrefix)
	return ipPrefix, ipLen, err
}

func addressTypeNeutronFromVnc(
	rule *models.PolicyRuleType,
	sg *models.SecurityGroup,
	responseSgr *SecurityGroupRuleResponse,
) error {
	pair := models.PolicyAddressPair{
		SourceAddress:      rule.GetSRCAddresses()[0],
		DestinationAddress: rule.GetDSTAddresses()[0],
	}
	isIngress, err := pair.IsIngress()
	if err != nil {
		return newNeutronError(securityGroupRuleNotFound, errorFields{
			"id": rule.GetRuleUUID(),
		})
	}

	var remoteAddr *models.AddressType
	if isIngress {
		responseSgr.Direction = ingressTrafficNeutron
		remoteAddr = pair.SourceAddress
	} else {
		responseSgr.Direction = egressTrafficNeutron
		remoteAddr = pair.DestinationAddress
	}

	if subnet := remoteAddr.GetSubnet(); subnet != nil {
		responseSgr.RemoteIPPrefix = getFullNetworkAddress(subnet.GetIPPrefix(), subnet.GetIPPrefixLen())
	} else if remoteAddr.IsSecurityGroupNameAReference() {
		remoteSG := remoteAddr.GetSecurityGroup()
		if remoteSG == basemodels.FQNameToString(sg.GetFQName()) {
			responseSgr.RemoteGroupID = sg.GetUUID()
		} else {
			// TODO implement it when service FQNameToID will be available in Neutron package.
			// Origin python code: /src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py:1273
		}
	}

	return nil
}

func getFullNetworkAddress(ip string, ipLen int64) string {
	return ip + "/" + strconv.FormatInt(ipLen, 10)
}

func securityGroupRuleCreate(
	ctx context.Context, rp RequestParameters, sg *models.SecurityGroup, sgr *models.PolicyRuleType,
) error {
	rules := sg.GetSecurityGroupEntries()
	if rules == nil {
		rules = &models.PolicyEntriesType{}
	}

	rules.PolicyRule = append(rules.PolicyRule, sgr)
	sg.SecurityGroupEntries = rules

	_, err := rp.WriteService.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
		SecurityGroup: sg,
		FieldMask: types.FieldMask{
			Paths: []string{models.SecurityGroupFieldSecurityGroupEntries},
		},
	})
	if err != nil {
		// TODO: write test to this errors while implementing Security Group Rule - CREATE.
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
			conflictRuleUUID := findConflictedRuleUUID(err, rules.PolicyRule)
			return newNeutronError(securityGroupRuleExists, errorFields{
				"resource": securityGroupRuleResourceName,
				"id":       conflictRuleUUID,
				"rule_id":  conflictRuleUUID,
			})
		}
		return errors.Wrapf(err, "can't save security_group_rule")
	}
	return nil
}

func findConflictedRuleUUID(err error, rules []*models.PolicyRuleType) string {
	errMsg := err.Error()
	for _, rule := range rules {
		if strings.Contains(errMsg, rule.RuleUUID) {
			return rule.RuleUUID
		}
	}
	return ""
}

func (sgr *SecurityGroupRule) parseProtocolParameter() error {
	protos := []string{models.AnyProtocol, models.TCPProtocol, models.UDPProtocol, models.ICMPProtocol}

	if protocol, err := strconv.ParseInt(sgr.Protocol, 10, 64); err == nil {
		if protocolMinValue < protocol && protocol < protocolMaxValue {
			return nil
		}
	}

	if containsString(protos, sgr.Protocol) {
		return nil
	}

	return newNeutronError(securityGroupRuleInvalidProtocol, errorFields{
		"protocol": sgr.Protocol,
		"values":   protos,
	})
}

func containsString(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
