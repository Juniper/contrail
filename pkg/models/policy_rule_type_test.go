package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIngress(t *testing.T) {
	testCases := []struct {
		name string
		policyAddressPair
		isIngress bool
		err       error
	}{
		{
			name: "specified security group to local security group",
			policyAddressPair: policyAddressPair{
				sourceAddress: &policyAddress{
					SecurityGroup: "default-domain:project-blue:default",
				},
				destinationAddress: &policyAddress{
					SecurityGroup: "local",
				},
			},
			isIngress: true,
		},
		{
			name: "local security group to all IPv4 addresses",
			policyAddressPair: policyAddressPair{
				sourceAddress: &policyAddress{
					SecurityGroup: "local",
				},
				destinationAddress: (*policyAddress)(AllIPv4Addresses()),
			},
			isIngress: false,
		},
		{
			name: "local security group to all IPv6 addresses",
			policyAddressPair: policyAddressPair{
				sourceAddress: &policyAddress{
					SecurityGroup: "local",
				},
				destinationAddress: (*policyAddress)(AllIPv6Addresses()),
			},
			isIngress: false,
		},
		{
			name: "both with local security group",
			policyAddressPair: policyAddressPair{
				sourceAddress: &policyAddress{
					SecurityGroup: "local",
				},
				destinationAddress: &policyAddress{
					SecurityGroup: "local",
				},
			},
			// https://github.com/Juniper/contrail-controller/blob/08f2b11d3/src/config/schema-transformer/config_db.py#L2030
			isIngress: true,
		},
		{
			name: "neither with local security group",
			policyAddressPair: policyAddressPair{
				sourceAddress:      &policyAddress{},
				destinationAddress: &policyAddress{},
			},
			err: neitherAddressIsLocal{
				sourceAddress:      &policyAddress{},
				destinationAddress: &policyAddress{},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			isIngress, err := tt.policyAddressPair.isIngress()
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.isIngress, isIngress)
			}
		})
	}
}

func TestIsLocal(t *testing.T) {
	testCases := []struct {
		name    string
		address *policyAddress
		is      bool
	}{
		{
			name: "local security group",
			address: &policyAddress{
				SecurityGroup: "local",
			},
			is: true,
		},
		{
			name: "specified security group",
			address: &policyAddress{
				SecurityGroup: "default-domain:project-blue:default",
			},
			is: false,
		},
		{
			name:    "all IPv4 addresses",
			address: (*policyAddress)(AllIPv4Addresses()),
			is:      false,
		},
		{
			name:    "all IPv6 addresses",
			address: (*policyAddress)(AllIPv6Addresses()),
			is:      false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.is, tt.address.isLocal())
		})
	}
}

func TestACLProtocol(t *testing.T) {
	testCases := []struct {
		name                string
		policyRule          *PolicyRuleType
		expectedACLProtocol string
		fails               bool
	}{
		{
			name: "any",
			policyRule: &PolicyRuleType{
				Protocol: "any",
			},
			expectedACLProtocol: "any",
		},

		{
			name: "not specified",
			policyRule: &PolicyRuleType{
				Protocol: "",
			},
			expectedACLProtocol: "",
		},

		{
			name: "already a number",
			policyRule: &PolicyRuleType{
				Protocol: "58",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "unknown IPv6 protocol",
			policyRule: &PolicyRuleType{
				Protocol:  "some unknown protocol",
				Ethertype: "IPv6",
			},
			fails: true,
		},

		{
			name: "unknown IPv4 protocol",
			policyRule: &PolicyRuleType{
				Protocol:  "some unknown protocol",
				Ethertype: "IPv4",
			},
			fails: true,
		},

		{
			name: "unknown ethertype and protocol",
			policyRule: &PolicyRuleType{
				Protocol:  "some unknown protocol",
				Ethertype: "some unknown ethertype",
			},
			fails: true,
		},

		{
			name: "icmp ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "icmp ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "1",
		},

		// The rest of the tests are the same for IPv6 and IPv4
		{
			name: "icmp6 ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp6",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "icmp6 ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "icmp6",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "58",
		},

		{
			name: "tcp ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "tcp",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "6",
		},

		{
			name: "tcp ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "tcp",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "6",
		},

		{
			name: "udp ipv6",
			policyRule: &PolicyRuleType{
				Protocol:  "udp",
				Ethertype: "IPv6",
			},
			expectedACLProtocol: "17",
		},

		{
			name: "udp ipv4",
			policyRule: &PolicyRuleType{
				Protocol:  "udp",
				Ethertype: "IPv4",
			},
			expectedACLProtocol: "17",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			aclProtocol, err := tt.policyRule.ACLProtocol()
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedACLProtocol, aclProtocol)
		})
	}
}
