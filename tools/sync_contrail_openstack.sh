#!/bin/bash

docker exec kolla_toolbox sh -c "source /var/lib/kolla/config_files/admin-openrc.sh && openstack project list -c ID -f csv | tail -n +2 | python -c 'import sys,uuid; print \"\n\".join([str(uuid.UUID(proj_id.strip().strip(\"\\\"\"))) for proj_id in sys.stdin])' | xargs -I PROJ_ID curl $(uname -n):8082/project/PROJ_ID"
