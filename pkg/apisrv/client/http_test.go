package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTP(t *testing.T) {
	tests := []struct {
		desc     string
		url      string
		protocol string
	}{{
		desc:     "https client test",
		url:      "http://fake_endpoint/",
		protocol: "http",
	}, {
		desc:     "https client test",
		url:      "https://fake_endpoint/",
		protocol: "https",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			h := NewHTTP(tt.url, tt.url, "fake_client", "fake_password", true, nil)
			h.Init()
			assert.Equal(t, tt.url+"foo", h.getURL("foo"), "getURL failed")
			assert.Equal(t, tt.url, h.AuthURL, "invalid auth URL")
			assert.Equal(t, tt.protocol, h.getProtocol(), "getProtocol failed")
		})
	}

}
