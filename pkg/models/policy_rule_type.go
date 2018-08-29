package models

import (
	"fmt"
	"strconv"
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
	return policyAddress.SecurityGroup == "local"
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

type unknownProtocol struct {
	protocol, ethertype string
}

func (err unknownProtocol) Error() string {
	return fmt.Sprintf("unknown protocol %v for ethertype %v", err.protocol, err.ethertype)
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

const IPv6Ethertype = "IPv6"

func (m *PolicyRuleType) ACLProtocol() (string, error) {
	protocol := m.GetProtocol()

	if protocol == "" || protocol == "any" || isNumeric(protocol) {
		return protocol, nil
	}

	var numericProtocol string
	var ok bool
	if m.GetEthertype() == IPv6Ethertype {
		numericProtocol, ok = ipV6ProtocolStringToNumber[protocol]
	} else {
		numericProtocol, ok = ipV4ProtocolStringToNumber[protocol]
	}

	if !ok {
		return "", unknownProtocol{
			protocol:  protocol,
			ethertype: m.GetEthertype(),
		}
	}
	return numericProtocol, nil
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
