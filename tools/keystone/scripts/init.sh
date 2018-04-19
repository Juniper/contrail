#!/bin/bash

set -ex

export OS_TOKEN=ADMIN_TOKEN
export OS_URL=http://localhost:35357/v3
export OS_IDENTITY_API_VERSION=3

openstack service create \
  --name keystone --description "OpenStack Identity" identity

openstack endpoint create --region RegionOne \
  identity public http://some-keystone:5000/v3

openstack endpoint create --region RegionOne \
  identity internal http://some-keystone:5000/v3

openstack endpoint create --region RegionOne \
  identity admin http://some-keystone:35357/v3

openstack domain create --description "Default Domain" default
openstack project create --domain default \
  --description "Admin Project" admin

openstack user create --domain default \
  --password contrail123 admin

openstack role create admin
openstack role add --project admin --user admin admin

openstack project create --domain default \
  --description "Service Project" service

openstack project create --domain default \
  --description "Demo Project" demo
openstack user create --domain default \
  --password contrail123 demo

openstack role create user
openstack role add --project demo --user demo user