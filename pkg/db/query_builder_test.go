package db

import (
	"net"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestStringIPv6(t *testing.T) {
	tests := []struct {
		ip     net.IP
		output string
	}{
		{
			ip:     net.ParseIP("10.0.0.1"),
			output: "::ffff:a00:1",
		},
		{
			ip:     net.ParseIP("0.0.1.1"),
			output: "::ffff:0:101",
		},
		{
			ip:     net.ParseIP("2001:db8:ac10:fe01::"),
			output: "2001:db8:ac10:fe01::",
		},
		{
			ip:     net.IP{},
			output: "",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, StringIPv6(tt.ip), tt.output)
	}
}
