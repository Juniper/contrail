package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
)

// EqualRule checks if rule contains same data as other rule.
func (m PolicyRuleType) EqualRule(other PolicyRuleType) bool {
	m.RuleUUID = ""
	m.LastModified = ""
	m.Created = ""

	other.RuleUUID = ""
	other.LastModified = ""
	other.Created = ""

	return reflect.DeepEqual(m, other)
}

var avaiableProtocols = []string{"any", "icmp", "tcp", "udp", "icmp6"}

var isAvailableProtocol = boolMap(avaiableProtocols)

// ValidateProtocol checks if protocol is valid rule protocol.
func (m *PolicyRuleType) ValidateProtocol() error {
	proto := m.GetProtocol()

	if isAvailableProtocol[proto] {
		return nil
	}

	number, err := strconv.Atoi(proto)
	if err != nil {
		return errutil.ErrorBadRequestf("Rule with invalid protocol: %v.", proto)
	}

	if number < 0 || number > 255 {
		return errutil.ErrorBadRequestf("Rule with invalid protocol: %v.", number)
	}

	return nil
}

// addressTypePair is a single combination of source and destination specifications from a PolicyRuleType.
type addressTypePair struct {
	policyRule                        *PolicyRuleType
	sourceAddress, destinationAddress *AddressType
	sourcePort, destinationPort       *PortType
}

func (pair *addressTypePair) isIngress() (bool, error) {
	switch {
	case pair.destinationAddress.isSecurityGroupLocal():
		return true, nil
	case pair.sourceAddress.isSecurityGroupLocal():
		return false, nil
	default:
		return false, neitherAddressIsLocal{
			sourceAddress:      pair.sourceAddress,
			destinationAddress: pair.destinationAddress,
		}
	}
}

type neitherAddressIsLocal struct {
	sourceAddress, destinationAddress *AddressType
}

func (err neitherAddressIsLocal) Error() string {
	return fmt.Sprintf("neither source nor destination address is local. Source address: %v. Destination address: %v",
		err.sourceAddress, err.destinationAddress)
}

func (m *PolicyRuleType) allAddressCombinations() (pairs []addressTypePair) {
	for _, sourceAddress := range m.SRCAddresses {
		for _, sourcePort := range m.SRCPorts {
			for _, destinationAddress := range m.DSTAddresses {
				for _, destinationPort := range m.DSTPorts {
					pairs = append(pairs, addressTypePair{
						policyRule: m,

						sourceAddress:      sourceAddress,
						sourcePort:         sourcePort,
						destinationAddress: destinationAddress,
						destinationPort:    destinationPort,
					})
				}
			}
		}
	}
	return pairs
}

var ipV6ProtocolStringToNumber = map[string]string{
	"icmp":  "58",
	"icmp6": "58",
	"tcp":   "6",
	"udp":   "17",
}

var ipV4ProtocolStringToNumber = map[string]string{
	"icmp":  "1",
	"icmp6": "58",
	"tcp":   "6",
	"udp":   "17",
}

// TODO: Generate this from the enum in the schema.
const ipv6Ethertype = "IPv6"

// ValidateSubnetsWithEthertype validates if every subnet
// within source and destination addresses matches rule ethertype.
func (m *PolicyRuleType) ValidateSubnetsWithEthertype() error {
	if m.GetEthertype() != "" {
		for _, addr := range m.GetSRCAddresses() {
			if err := addr.Subnet.ValidateWithEthertype(m.GetEthertype()); err != nil {
				return err
			}
		}
		for _, addr := range m.GetDSTAddresses() {
			if err := addr.Subnet.ValidateWithEthertype(m.GetEthertype()); err != nil {
				return err
			}
		}
	}
	return nil
}

// HasSecurityGroup returns true if any of addresses points at Security Group.
func (m *PolicyRuleType) HasSecurityGroup() bool {
	for _, addr := range m.GetSRCAddresses() {
		if addr.GetSecurityGroup() != "" {
			return true
		}
	}
	for _, addr := range m.GetDSTAddresses() {
		if addr.GetSecurityGroup() != "" {
			return true
		}
	}
	return false
}

// IsAnySecurityGroupAddrLocal returns true if at least one of addresses contains
// 'local' Security Group.
func (m *PolicyRuleType) IsAnySecurityGroupAddrLocal() bool {
	for _, addr := range m.GetSRCAddresses() {
		if addr.isSecurityGroupLocal() {
			return true
		}
	}
	for _, addr := range m.GetDSTAddresses() {
		if addr.isSecurityGroupLocal() {
			return true
		}
	}
	return false
}

// ACLProtocol returns the protocol in a format suitable for an AclRuleType.
func (m *PolicyRuleType) ACLProtocol() (string, error) {
	protocol := m.GetProtocol()
	ethertype := m.GetEthertype()

	if protocol == "" || protocol == "any" || isNumeric(protocol) {
		return protocol, nil
	}

	protocol, err := numericProtocolForEthertype(protocol, ethertype)
	if err != nil {
		return "", errors.Wrap(err, "failed to convert protocol for an ACL")
	}
	return protocol, nil
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func numericProtocolForEthertype(protocol, ethertype string) (numericProtocol string, err error) {
	var ok bool
	if ethertype == ipv6Ethertype {
		numericProtocol, ok = ipV6ProtocolStringToNumber[protocol]
	} else {
		numericProtocol, ok = ipV4ProtocolStringToNumber[protocol]
	}

	if !ok {
		return "", errors.Errorf("unknown protocol %q for ethertype %q", protocol, ethertype)
	}
	return numericProtocol, nil
}
