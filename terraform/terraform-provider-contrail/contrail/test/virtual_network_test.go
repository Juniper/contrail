package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestVirtualNetwork_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMPublicIpDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAzureRMVPublicIpStatic_basic,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMPublicIpExists("azurerm_public_ip.test"),
				),
			},
		},
	})
}

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}
