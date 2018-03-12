variable "server_address" {
	description = "Contrail server address"
}

variable "username" {
	description = "Username used for authentication via keystone"
}

variable "tenant_name" {
	description = "Tenant to be used"
}

variable "password" {
	description = "Password used for authentication via keystone"
}

provider "contrail" {
	server = "${var.server_address}"
	auth_url = "http://${var.server_address}:5000/v2.0/"
	username = "${var.username}"
	tenant_name = "${var.tenant_name}"
	password = "${var.password}"
}

resource "contrail_virtual_network" "spocknet" {
	name = "spocknet"
	display_name = "spocknet"
}

resource "contrail_virtual_network_refs" "spockrefs" {
	uuid = "${contrail_virtual_network.spocknet.id}"
	network_ipam_refs {
		to = "${contrail_network_ipam.spock_ipam.id}"
		attr {
			ipam_subnets {
				subnet_name = "spock_subnet"
				subnet {
					ip_prefix = "10.10.20.0"
					ip_prefix_len = 24
				}
			}
		}
	}
}

resource "contrail_network_ipam" "spock_ipam" {
	name = "spock_ipam"
	#dns_nameservers = ["1.1.1.1", "2.2.2.2", "3.3.3.3", "8.8.8.8"]
}

output "spocknet-id" {
	value = "${contrail_virtual_network.spocknet.id}"
}
