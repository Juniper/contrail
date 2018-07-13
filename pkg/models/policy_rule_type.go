package models

import "fmt"

// policyAddressPair is a single combination of source and destination specifications from a PolicyRuleType.
type policyAddressPair struct {
	policyRule                        *PolicyRuleType
	sourceAddress, destinationAddress *policyAddress
	sourcePort, destinationPort       *PortType
}

func (pair policyAddressPair) isIngress() (bool, error) {
	switch {
	case pair.sourceAddress.isLocal() && pair.destinationAddress.isLocal():
		return true, nil
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

func (policyAddress policyAddress) isLocal() bool {
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
