#!/bin/bash

set -ex

keystone-manage bootstrap \
--bootstrap-password contrail123 \
--bootstrap-service-name keystone \
--bootstrap-admin-url http://some-keystone:35357 \
--bootstrap-public-url http://some-keystone:5000 \
--bootstrap-internal-url http://some-keystone:5000 \
--bootstrap-project-name admin \
--bootstrap-role-name admin \
--bootstrap-region-id RegionOne