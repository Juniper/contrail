#!/bin/bash

PASSWORD=contrail123
docker rm -f contrail_postgres
docker rm -f contrail_mysql
docker run -d --name contrail_postgres -p 5432:5432 -e "POSTGRES_USER=root" -e "POSTGRES_DB=$PASSWORD" circleci/postgres:10.3-alpine
docker run -d --name contrail_mysql -p 3306:3306  -e "MYSQL_ROOT_PASSWORD=$PASSWORD"  circleci/mysql:5.7