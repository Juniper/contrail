package models

import (
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/gogo/protobuf/types"
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

// TagTypeIDs contains tag-type names with corresponding ids
var TagTypeIDs = map[string]int64{ //TODO move to TagType file
	"label":       0,
	"application": 1,
	"tier":        2,
	"deployment":  3,
	"site":        4,
}

// CheckAssociatedRefsInSameScope checks scope of firewallRule refs
// that global scoped firewall rule cannot reference scoped resources
func (fr *FirewallRule) CheckAssociatedRefsInSameScope(
	databaseFR *FirewallRule,
) error {
	if databaseFR == nil && len(fr.GetFQName()) != 2 {
		return nil
	}

	if databaseFR != nil && len(databaseFR.GetFQName()) != 2 {
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

// CheckServiceProperties checks for existence of service and serviceGroupRefs property
func (fr *FirewallRule) CheckServiceProperties(databaseFR *FirewallRule) error {
	serviceGroupRefs := fr.GetServiceGroupRefs()
	if serviceGroupRefs == nil {
		serviceGroupRefs = databaseFR.GetServiceGroupRefs()
	}

	service := fr.GetService()
	if service == nil {
		service = databaseFR.GetService()
	}

	if service == nil && len(serviceGroupRefs) == 0 {
		return common.ErrorBadRequest("firewall Rule requires at least 'service' property or Service Group references(s)")
	}

	if service != nil && len(serviceGroupRefs) > 0 {
		return common.ErrorBadRequest(
			"firewall Rule cannot have both defined 'service' property and Service Group reference(s)",
		)
	}

	return nil
}

// AddDefaultMatchTag sets default matchTag if not defined in the request
func (fr *FirewallRule) AddDefaultMatchTag(fm *types.FieldMask) {
	if fr.GetMatchTags() == nil &&
		basemodels.FieldMaskContains(fm, FirewallRuleFieldMatchTags) {
		fr.MatchTags = &FirewallRuleMatchTagsType{
			TagList: []string{DefaultMatchTagType},
		}
	}
}

// SetProtocolID sets protocolID based on protocol property
func (fr *FirewallRule) SetProtocolID(fm *types.FieldMask) error {
	if !basemodels.FieldMaskContains(fm, FirewallRuleFieldService) {
		return nil
	}

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

// CheckEndpoints checks if endpoint properties refer to only one endpoint type
func (fr *FirewallRule) CheckEndpoints() error {
	endpoints := []*FirewallRuleEndpointType{
		fr.GetEndpoint1(),
		fr.GetEndpoint2(),
	}

	for _, endpoint := range endpoints {
		if endpoint == nil {
			continue
		}

		count := 0
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
