#!/bin/bash

docker rm -f some-keystone

docker run --name some-keystone \
    -v `pwd`/keystone/apache2:/etc/apache2/sites-available/ \
    -v `pwd`/keystone/etc:/etc/keystone \
    -v `pwd`/keystone/scripts:/tmp \
    -p 5000:5000 \
    -d \
    openstackhelm/keystone:newton \
    bash /tmp/start.sh

sleep 10 

docker exec some-keystone bash /tmp/init.sh
