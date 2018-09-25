package models

import (
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// FirewallRule constants.
const (
	DefaultMatchTagType = "application"
)

var protocolIDs = map[string]int64{
	"any":  0,
	"icmp": 1,
	"tcp":  6,
	"udp":  17,
}

//TagTypeIDs contains tag-type names with corresponding ids
var TagTypeIDs = map[string]int64{ //TODO move to TagType file
	"label":       0,
	"application": 1,
	"tier":        2,
	"deployment":  3,
	"site":        4,
}

// CheckAssociatedRefsInSameScope TODO
func (fr *FirewallRule) CheckAssociatedRefsInSameScope() error {
	if len(fr.GetFQName()) != 2 {
		return nil
	}

	for _, ref := range fr.GetAddressGroupRefs() {
		if err := fr.checkRefInSameScope(ref); err != nil {
			return err
		}
	}

	for _, ref := range fr.GetServiceGroupRefs() {
		if err := fr.checkRefInSameScope(ref); err != nil {
			return err
		}
	}

	for _, ref := range fr.GetVirtualNetworkRefs() {
		if err := fr.checkRefInSameScope(ref); err != nil {
			return err
		}
	}

	return nil
}

func (fr *FirewallRule) checkRefInSameScope(ref basemodels.Reference) error {
	if len(ref.GetTo()) == 2 {
		return nil
	}

	refKind := strings.Replace(ref.GetReferredKind(), "-", " ", -1)
	return common.ErrorBadRequestf(
		"global Firewall Rule %s (%s) cannot reference a scoped %s %s (%s)",
		basemodels.FQNameToString(fr.GetFQName()),
		fr.GetUUID(),
		strings.Title(refKind),
		basemodels.FQNameToString(ref.GetTo()),
		ref.GetUUID(),
	)
}

// CheckServiceProperties TODO
func (fr *FirewallRule) CheckServiceProperties() error {
	serviceGroupRefs := fr.GetServiceGroupRefs()
	service := fr.GetService()

	if service == nil && len(serviceGroupRefs) == 0 {
		return common.ErrorBadRequest("firewall Rule requires at least 'service' property or Service Group references(s)")
	}

	if service != nil && len(serviceGroupRefs) > 0 {
		return common.ErrorBadRequest("firewall Rule cannot have both defined 'service' property and Service Group reference(s)")
	}

	return nil
}

// AddDefaultMatchTag TODO
func (fr *FirewallRule) AddDefaultMatchTag() {
	if fr.GetMatchTags() == nil {
		fr.MatchTags = &FirewallRuleMatchTagsType{
			TagList: []string{DefaultMatchTagType},
		}
	}
}

// SetProtocolID TODO
func (fr *FirewallRule) SetProtocolID() error {
	protocol := fr.GetService().GetProtocol()
	ok := true

	protocolID, err := strconv.ParseInt(protocol, 10, 64)
	if err != nil {
		protocolID, ok = protocolIDs[protocol]
	}

	if !ok || protocolID < 0 || protocolID > 255 {
		return common.ErrorBadRequestf("rule with invalid protocol: %s", protocol)
	}

	fr.Service.ProtocolID = protocolID
	return nil
}

// CheckEndpoints TODO
func (fr *FirewallRule) CheckEndpoints() error {
	endpoints := []*FirewallRuleEndpointType{
		fr.GetEndpoint1(),
		fr.GetEndpoint2(),
	}

	var count int
	for _, endpoint := range endpoints {
		if endpoint == nil {
			continue
		}

		if endpoint.GetAddressGroup() != "" {
			count++
		}
		if endpoint.GetAny() == true {
			count++
		}
		if endpoint.GetSubnet() != nil {
			count++
		}
		if len(endpoint.GetTags()) > 0 {
			count++
		}
		if endpoint.GetVirtualNetwork() != "" {
			count++
		}

		if count > 1 {
			return common.ErrorBadRequest("endpoint is limited to only one endpoint type at a time")
		}
	}

	return nil
}
