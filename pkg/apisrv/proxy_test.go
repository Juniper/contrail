package apisrv

import (
	"testing"

	"github.com/flosch/pongo2"
)

func TestProxyEndpoint(t *testing.T) {
	context := pongo2.Context{
		"cluster_a_keystone_private_url": "http://127.0.0.1:5000/v3",
		"cluster_a_keystone_public_url":  "http://127.0.0.1:35357/v3",
	}

	RunTestWithTemplate(t, "./test_data/test_endpoint.tmpl", context)
}
