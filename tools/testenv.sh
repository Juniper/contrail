#!/bin/bash

set -e

PASSWORD=contrail123
docker network create contrail || echo > /dev/null
docker rm -f contrail_postgres  || echo > /dev/null
docker rm -f contrail_mysql || echo > /dev/null
docker run -d --name contrail_postgres --net contrail -p 5432:5432 -e "POSTGRES_USER=root" -e "POSTGRES_DB=$PASSWORD" circleci/postgres:10.3-alpine
docker run -d --name contrail_mysql --net contrail -p 3306:3306  -e "MYSQL_ROOT_PASSWORD=$PASSWORD"  circleci/mysql:5.7

echo "Waiting for mysql"
until docker exec contrail_mysql mysql -uroot -p"$PASSWORD" -e "show status" &> /dev/null
do
  printf "."
  sleep 1
done
echo "done"