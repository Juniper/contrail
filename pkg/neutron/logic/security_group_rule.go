package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	uuid "github.com/satori/go.uuid"
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
	ctx context.Context, rp RequestParameters, f Filters, _ Fields,
) (Response, error) {
	response, err := listSecurityGroupRules(ctx, rp, f)
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": securityGroupRuleResourceName,
			"msg":      err,
		})
	}
	return response, nil
}

// Create security group rule.
func (sgr *SecurityGroupRule) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	if sgr.ID == "" {
		sgr.ID = uuid.NewV4().String()
	}

	var sgRes *services.GetSecurityGroupResponse
	var err error

	if sgRes, err = rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
		ID: sgr.SecurityGroupID,
	}); err != nil {
		return nil, err
	}

	var rule *models.PolicyRuleType
	if rule, err = sgr.vncFromNeutron(ctx, rp); err != nil {
		return nil, err
	}

	if err = securityGroupRuleCreate(ctx, rp, sgRes.GetSecurityGroup(), rule); err != nil {
		return nil, err
	}

	return sgr.neutronFromVnc(sgRes.GetSecurityGroup(), rule)
}

// Read security group rule logic.
func (sgr *SecurityGroupRule) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	sgrResponses, err := sgr.getFilterSecurityGroupRules(ctx, rp, id)
	if err != nil {
		return nil, err
	}
	for _, sgrResponse := range sgrResponses {
		if sgrResponse.ID == id {
			return sgrResponse, nil
		}
	}

	return nil, newNeutronError(securityGroupRuleNotFound, nil)
}

// Delete security group rule.
func (sgr *SecurityGroupRule) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	sg, err := sgr.findSecurityGroupByRuleID(ctx, rp, id)
	if err != nil {
		return nil, err
	}

	if sg == nil {
		return nil, newNeutronError(securityGroupNotFound, errorFields{
			"security_group_rule_id": id,
		})
	}

	var rules []*models.PolicyRuleType
	for _, r := range sg.GetSecurityGroupEntries().GetPolicyRule() {
		if r.RuleUUID != id {
			rules = append(rules, r)
		}
	}
	sg.SecurityGroupEntries.PolicyRule = rules
	return rp.WriteService.UpdateSecurityGroup(ctx, &services.UpdateSecurityGroupRequest{
		SecurityGroup: sg,
		FieldMask: types.FieldMask{
			Paths: []string{models.SecurityGroupFieldSecurityGroupEntries},
		},
	})
}

func (sgr *SecurityGroupRule) getFilterSecurityGroupRules(
	ctx context.Context,
	rp RequestParameters,
	id string,
) ([]*SecurityGroupRuleResponse, error) {
	var f Filters
	if !rp.RequestContext.IsAdmin {
		projectUUID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
		if err != nil {
			return nil, newNeutronError(badRequest, errorFields{
				"resource": securityGroupRuleResourceName,
				"msg":      err,
			})
		}
		// Trigger a project read to ensure project sync
		if _, err := rp.ReadService.GetProject(ctx, &services.GetProjectRequest{ID: projectUUID}); err != nil {
			return nil, newNeutronError(badRequest, errorFields{
				"resource": securityGroupRuleResourceName,
				"msg":      err,
			})
		}
		f = Filters{tenantIDKey: []string{projectUUID}}
	}

	sgrResponses, err := listSecurityGroupRules(ctx, rp, f)
	if errutil.IsNotFound(err) {
		return nil, newNeutronError(securityGroupNotFound, errorFields{
			"resource":               securityGroupRuleResourceName,
			"security_group_rule_id": id,
			"msg":                    err,
		})
	}
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": securityGroupRuleResourceName,
			"msg":      err,
		})
	}
	return sgrResponses, nil
}

func (sgr *SecurityGroupRule) findSecurityGroupByRuleID(
	ctx context.Context,
	rp RequestParameters,
	id string,
) (*models.SecurityGroup, error) {
	sgrResponses, err := sgr.getFilterSecurityGroupRules(ctx, rp, id)
	if err != nil {
		return nil, err
	}
	for _, sgrResponse := range sgrResponses {
		if sgrResponse.ID != id {
			continue
		}
		var sgResponse *services.GetSecurityGroupResponse
		sgResponse, err = rp.ReadService.GetSecurityGroup(ctx, &services.GetSecurityGroupRequest{
			ID: sgrResponse.SecurityGroupID,
		})
		if errutil.IsNotFound(err) {
			return nil, newNeutronError(securityGroupNotFound, errorFields{
				"resource":               securityGroupRuleResourceName,
				"security_group_rule_id": id,
				"msg":                    err,
			})
		}
		if err != nil {
			return nil, newNeutronError(badRequest, errorFields{
				"resource":               securityGroupRuleResourceName,
				"security_group_rule_id": id,
				"msg":                    err,
			})
		}
		return sgResponse.GetSecurityGroup(), nil
	}

	return nil, newNeutronError(securityGroupNotFound, errorFields{
		"resource":               securityGroupRuleResourceName,
		"security_group_rule_id": id,
		"msg":                    err,
	})
}

func listSecurityGroupRules(
	ctx context.Context, rp RequestParameters, f Filters,
) ([]*SecurityGroupRuleResponse, error) {
	securityGroups, err := listSecurityGroups(ctx, rp, Filters{
		tenantIDKey: f[tenantIDKey],
		idKey:       f[securityGroupIDKey],
	})

	if err != nil {
		return nil, errors.Wrapf(err, "can't read security groups")
	}

	securityGroupRules := make([]*SecurityGroupRuleResponse, 0)
	for _, sg := range securityGroups {
		sgr, err := (&SecurityGroup{}).readSecurityGroupRules(sg)
		if err != nil {
			return nil, errors.Wrapf(err, "can't read security group rules")
		}
		securityGroupRules = append(securityGroupRules, sgr...)
	}

	return securityGroupRules, nil
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
		RuleUUID:     sgr.ID,
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
