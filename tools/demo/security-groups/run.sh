#!/bin/bash

if [ -z "$1" ]; then
	echo "Usage: run.sh POD_FULL_NAME"
	exit 1
fi

CONTRAIL_ADDR="localhost:8082"
POD=$1

cat > vmi_update.json.tmpl << "EOF"
{
    "virtual-machine-interface": {
        "fq_name": [
            "default-domain",
            "k8s-atom-pink",
            "{{vmi-name}}"
        ],
        "security_group_refs": [
            {
                "to": [
                    "default-domain",
                    "k8s-atom-pink",
                    "{{sg-name}}"
                ]
            }
        ]
    }
}
EOF

get_vmi_name() {
	local POD_NAME=$1

	curl "localhost:8082/virtual-machine-interfaces" | python -m json.tool | sed -n -e "s/^.*\"\(${POD_NAME}__.*\)\".*$/\1/p" | head -n 1
}

vmi_uuid_from_name() {
	echo "$1" | sed 's/.*__\(.*\)/\1/'
}

update_ref() {
	local VMI_NAME=$1
	local VMI_UUID; VMI_UUID=$(vmi_uuid_from_name "$VMI_NAME")
	local SG_NAME=$2

	echo "Setting ref from VMI $VMI_NAME to SG $SG_NAME..."
	sed "s/{{vmi-name}}/$VMI_NAME/; s/{{sg-name}}/$SG_NAME/;" vmi_update.json.tmpl > vmi_update.json
	echo "vmi_update.json:"
	cat vmi_update.json
	set -x
	curl -X PUT -H "Content-Type: application/json; charset=UTF-8" -d @vmi_update.json "$CONTRAIL_ADDR/virtual-machine-interface/$VMI_UUID"
	set +x
	echo "Reference updated."
}

create_sg() {
	local SG_JSON=$1
	echo "Creating security group..."
	echo "$SG_JSON:"
	cat "$SG_JSON"
	set -x
	curl -X POST -H "Content-Type: application/json; charset=UTF-8" -d "@$SG_JSON" $CONTRAIL_ADDR/security-groups
	set +x
	echo "Security group created."
}

deny() {
	local VMI_NAME=$1

	cat > sg.json << "EOF"
{
	"security-group": {
		"parent_type": "project",
		"id_perms": {
			"enable": true,
			"description": "Deny all",
			"user_visible": true
		},
		"fq_name": [
			"default-domain",
			"k8s-atom-pink",
			"deny-all"
		],
		"security_group_entries": {
			"policy_rule": []
		}
	}
}
EOF
	echo "Scenario: deny all traffic."
	read

	create_sg "sg.json"
	read

	update_ref "$VMI_NAME" "deny-all"
	read

	echo "All traffic from/to $VMI_NAME is now blocked."
	read
}

allow_tcp() {
	local VMI_NAME=$1

	cat > sg.json << "EOF"
{
	"security-group": {
		"parent_type": "project",
		"id_perms": {
			"enable": true,
			"description": "Allow TCP",
			"user_visible": true
		},
		"fq_name": [
			"default-domain",
			"k8s-atom-pink",
			"allow-tcp"
		],
		"security_group_entries": {
			"policy_rule": [
				{
					"direction": ">",
					"src_addresses": [ { "subnet": { "ip_prefix": "0.0.0.0", "ip_prefix_len": 0 } } ],
					"src_ports": [ { "end_port": 65535 } ],
					"ethertype": "IPv4",
					"protocol": "tcp",
					"dst_addresses": [ { "security_group": "local" } ],
					"dst_ports": [ { "end_port": 65535 } ]
				},
				{
					"direction": ">",
					"dst_addresses": [ { "subnet": { "ip_prefix": "0.0.0.0", "ip_prefix_len": 0 } } ],
					"dst_ports": [ { "end_port": 65535 } ],
					"ethertype": "IPv4",
					"protocol": "tcp",
					"src_addresses": [ { "security_group": "local" } ],
					"src_ports": [ { "end_port": 65535 } ]
				}
			]
		}
	}
}
EOF
	echo "Scenario: allow TCP traffic."
	read

	create_sg "sg.json"
	read

	update_ref "$VMI_NAME" "allow-tcp"
	read

	echo "Only TCP traffic from/to $VMI_NAME is now allowed."
	read
}

allow_icmp() {
	local VMI_NAME=$1

	cat > sg.json << "EOF"
{
	"security-group": {
		"parent_type": "project",
		"id_perms": {
			"enable": true,
			"description": "Allow ICMP",
			"user_visible": true
		},
		"fq_name": [
			"default-domain",
			"k8s-atom-pink",
			"allow-icmp"
		],
		"security_group_entries": {
			"policy_rule": [
				{
					"direction": ">",
					"src_addresses": [ { "subnet": { "ip_prefix": "0.0.0.0", "ip_prefix_len": 0 } } ],
					"src_ports": [ { "end_port": 65535 } ],
					"ethertype": "IPv4",
					"protocol": "icmp",
					"dst_addresses": [ { "security_group": "local" } ],
					"dst_ports": [ { "end_port": 65535 } ]
				},
				{
					"direction": ">",
					"dst_addresses": [ { "subnet": { "ip_prefix": "0.0.0.0", "ip_prefix_len": 0 } } ],
					"dst_ports": [ { "end_port": 65535 } ],
					"ethertype": "IPv4",
					"protocol": "icmp",
					"src_addresses": [ { "security_group": "local" } ],
					"src_ports": [ { "end_port": 65535 } ]
				}
			]
		}
	}
}
EOF
	echo "Scenario: allow ICMP traffic."
	read

	create_sg "sg.json"
	read

	update_ref "$VMI_NAME" "allow-icmp"
	read

	echo "Only ICMP traffic from/to $VMI_NAME is now allowed."
	read
}

default() {
	local VMI_NAME=$1

	echo "Scenario: allow all traffic."
	read

	update_ref "$1" "k8s-atom-pink-default-sg"

	echo "All traffic from/to $VMI_NAME is now allowed."
	read
}

VMI=$(get_vmi_name "$POD")
deny "$VMI"
allow_icmp "$VMI"
allow_tcp "$VMI"
default "$VMI"
