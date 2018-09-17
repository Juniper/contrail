package models

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// policyAddressPair is a single combination of source and destination specifications from a PolicyRuleType.
type policyAddressPair struct {
	policyRule                        *PolicyRuleType
	sourceAddress, destinationAddress *policyAddress
	sourcePort, destinationPort       *PortType
}

func (pair *policyAddressPair) isIngress() (bool, error) {
	switch {
	case pair.destinationAddress.isLocal():
		return true, nil
	case pair.sourceAddress.isLocal():
		return false, nil
	default:
		return false, neitherAddressIsLocal{
			sourceAddress:      pair.sourceAddress,
			destinationAddress: pair.destinationAddress,
		}
	}
}

// policyAddress is an address from a PolicyRuleType.
type policyAddress AddressType

func (policyAddress *policyAddress) isLocal() bool {
	return policyAddress.SecurityGroup == LocalSecurityGroup
}

type neitherAddressIsLocal struct {
	sourceAddress, destinationAddress *policyAddress
}

func (err neitherAddressIsLocal) Error() string {
	return fmt.Sprintf("neither source nor destination address is local. Source address: %v. Destination address: %v",
		err.sourceAddress, err.destinationAddress)
}

func (m *PolicyRuleType) allAddressCombinations() (pairs []policyAddressPair) {
	for _, sourceAddress := range m.SRCAddresses {
		for _, sourcePort := range m.SRCPorts {
			for _, destinationAddress := range m.DSTAddresses {
				for _, destinationPort := range m.DSTPorts {
					pairs = append(pairs, policyAddressPair{
						policyRule: m,

						sourceAddress:      (*policyAddress)(sourceAddress),
						sourcePort:         sourcePort,
						destinationAddress: (*policyAddress)(destinationAddress),
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
