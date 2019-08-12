#!/bin/bash
source /multicloud_tor_config/bin/activate
ansible-playbook -i /tmp/contrail_cluster/test_cluster_uuid/multi-cloud/inventories/inventory.yml ansible-multicloud/tor/playbooks/deploy_and_run_all.yml
