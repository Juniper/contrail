#!/bin/bash

#docker pull circleci/postgres:10.2-alpine
#docker run -e "POSTGRES_USER=root" -d -p 127.0.0.1:5432:5432 --name psql -e "POSTGRES_DB=contrail" circleci/postgres:10.2-alpine

psql -h localhost -U postgres -c "drop database contrail;"
psql -h localhost -U postgres -c "create database contrail;"
psql -h localhost -U root -d contrail < tools/init.sql
