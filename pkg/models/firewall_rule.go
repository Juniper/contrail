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
	DefaultMatchTagType               = "application"
	FirewallPolicyTagNameGlobalPrefix = "global:"
)

var protocolIDs = map[string]int64{
	"any":  0,
	"icmp": 1,
	"tcp":  6,
	"udp":  17,
}

// CheckAssociatedRefsInSameScope checks scope of firewallRule refs
func (fr *FirewallRule) CheckAssociatedRefsInSameScope(fqName []string) error {

	var refs []basemodels.Reference
	for _, ref := range fr.GetAddressGroupRefs() {
		refs = append(refs, ref)
	}

	for _, ref := range fr.GetServiceGroupRefs() {
		refs = append(refs, ref)
	}

	for _, ref := range fr.GetVirtualNetworkRefs() {
		refs = append(refs, ref)
	}

	return basemodels.CheckAssociatedRefsInSameScope(fr, fqName, refs)
}

// AddDefaultMatchTag sets default matchTag if not defined in the request
func (fr *FirewallRule) AddDefaultMatchTag(fm *types.FieldMask) {
	if fr.GetMatchTags().GetTagList() == nil && (fm == nil ||
		basemodels.FieldMaskContains(fm, FirewallRuleFieldMatchTags)) {
		fr.MatchTags = &FirewallRuleMatchTagsType{
			TagList: []string{DefaultMatchTagType},
		}
	}
}

// GetProtocolID returns id based on service's protocol
func (fr *FirewallRule) GetProtocolID() (int64, error) {
	protocol := fr.GetService().GetProtocol()
	ok := true
	protocolID, err := strconv.ParseInt(protocol, 10, 64)
	if err != nil {
		protocolID, ok = protocolIDs[protocol]
	}

	if !ok || protocolID < 0 || protocolID > 255 {
		return 0, common.ErrorBadRequestf("rule with invalid protocol: %s", protocol)
	}

	return protocolID, nil
}

// CheckEndpoints validates endpoint_1 and endpoint_2
func (fr *FirewallRule) CheckEndpoints() error {

	//TODO Authorize only one condition per endpoint for the moment
	endpoints := []*FirewallRuleEndpointType{
		fr.GetEndpoint1(),
		fr.GetEndpoint2(),
	}

	for _, endpoint := range endpoints {
		if err := endpoint.ValidateEndpointType(); err != nil {
			return err
		}
	}

	//TODO check no ids present
	//TODO check endpoints exclusivity clause
	//TODO VN name in endpoints

	return nil
}

// GetTagFQName returns fqname based on a tag name and firewall rule
func (fr *FirewallRule) GetTagFQName(
	tagName string, parentType string, fqName []string,
) ([]string, error) {
	if !strings.Contains(tagName, "=") {
		return nil, common.ErrorNotFoundf("invalid tag name '%s'", tagName)
	}

	if strings.HasPrefix(tagName, FirewallPolicyTagNameGlobalPrefix) {
		return []string{tagName[7:]}, nil
	}

	name := append([]string(nil), fqName...)
	if parentType == KindPolicyManagement {
		return append(name[:len(fqName)-2], tagName), nil
	}

	name = append([]string(nil), fqName...)
	if parentType == KindProject {
		return append(name[:len(fqName)-1], tagName), nil
	}

	return nil, common.ErrorBadRequestf(
		"Firewall rule %s (%s) parent type '%s' is not supported as security resource scope",
		fr.GetUUID(),
		fqName,
		parentType,
	)
}

// GetEndpoints returns request and database endpoints properties of firewall rules
func (fr *FirewallRule) GetEndpoints(
	databaseFR *FirewallRule,
) ([]*FirewallRuleEndpointType, []*FirewallRuleEndpointType) {
	endpoints := []*FirewallRuleEndpointType{
		fr.GetEndpoint1(),
		fr.GetEndpoint2(),
	}

	dbEndpoints := []*FirewallRuleEndpointType{
		databaseFR.GetEndpoint1(),
		databaseFR.GetEndpoint2(),
	}

	return endpoints, dbEndpoints
}

// FirewallRuleTagRef stub to be removed when tag refs implemented
type FirewallRuleTagRef struct {
	UUID string
	To   []string
}

// GetTagRefs to be removed when tag refs implemented
func (fr *FirewallRule) GetTagRefs() []*FirewallRuleTagRef {
	return []*FirewallRuleTagRef{}
}
