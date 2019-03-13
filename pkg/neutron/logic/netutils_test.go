package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIPPrefixAndPrefixLen(t *testing.T) {

	tests := []struct {
		name                 string
		input                string
		outputIPPrefix       string
		outputIPNetworPrefix string
		outputIPPrefixLen    int64
		expectError          bool
	}{
		{"good ipv4 1", "137.105.215.54/23", "137.105.215.54", "137.105.214.0", 23, false},
		{"good ipv4 2", "3.37.8.253/13", "3.37.8.253", "3.32.0.0", 13, false},
		{"good ipv6 1", "FE80:0000:0000:0000:0202:B3FF:FE1E:8329/64", "fe80::202:b3ff:fe1e:8329", "fe80::", 64, false},
		{"good ipv6 2", "2001:DB8:ABCD:12::/96", "2001:db8:abcd:12::", "2001:db8:abcd:12::", 96, false},
		{"good ipv6 3 no mask", "2001:DB8:ABCD:12::", "", "", 0, true},
		{"zero ipv4 address", "0.0.0.0/0", "0.0.0.0", "0.0.0.0", 0, false},
		{"zero ipv6 address", "::/0", "::", "::", 0, false},
		{"empty ip address", "", "", "", 0, true},
		{"wrong ip - invalid separator", "137.105.215.54:23", "", "", 0, true},
		{"wrong ip - invalid mask", "137.105.215.54/23a", "", "", 0, true},
		{"wrong ip - random chars", "abcd", "", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipPrefix, ipNetwork, ipPrefixLen, err := getIPPrefixAndPrefixLen(tt.input)

			if err != nil && !tt.expectError {
				t.Errorf("Expected no error but got: '%+v'", err)
			}

			assert.EqualValues(t, tt.outputIPPrefix, ipPrefix, "Error wrong IP address")
			assert.EqualValues(t, tt.outputIPNetworPrefix, ipNetwork, "Error wrong IP Network address")
			assert.EqualValues(t, tt.outputIPPrefixLen, ipPrefixLen, "Error wrong IP Prefix length")
		})
	}
}

func TestDecodeIP(t *testing.T) {
	tests := []struct {
		name                 string
		input                string
		outputIPPrefix       string
		outputIPNetworPrefix string
		outputIPPrefixLen    int64
		outputVersion        int8
		expectError          bool
	}{
		{"good ipv4 mask", "137.105.215.54/23", "137.105.215.54", "137.105.214.0", 23, ipV4, false},
		{"good ipv4 nomask", "3.37.8.253", "3.37.8.253", "0.0.0.0", 0, ipV4, false},
		{"good ipv6 mask", "FE80:0000:0000:0000:0202:B3FF:FE1E:8329/64",
			"fe80::202:b3ff:fe1e:8329", "fe80::", 64, ipV6, false},
		{"good ipv6 nomask", "2001:DB8:ABCD:12::", "2001:db8:abcd:12::", "::", 0, ipV6, false},
		{"wrong ip - invalid separator", "137.105.215.54:23", "", "", 0, 0, true},
		{"wrong ip - invalid mask", "137.105.215.54/23a", "", "", 0, 0, true},
		{"wrong ip - multiple separator", "137.105.215.54/23/4", "", "", 0, 0, true},
		{"wrong ip - random chars", "abcd", "", "", 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip, err := decodeIP(tt.input)

			if err != nil && !tt.expectError {
				t.Errorf("Expected no error but got: '%+v'", err)
			}

			assert.EqualValues(t, tt.outputIPPrefix, ip.IP, "Error wrong IP address")
			assert.EqualValues(t, tt.outputIPNetworPrefix, ip.netIP, "Error wrong IP Network address")
			assert.EqualValues(t, tt.outputIPPrefixLen, ip.prefixLen, "Error wrong IP Prefix length")
			assert.EqualValues(t, tt.outputVersion, ip.ver, "Error wrong IP version")
		})
	}
}
