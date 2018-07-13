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
