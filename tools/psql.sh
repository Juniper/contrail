#!/bin/bash

psql -h localhost -U postgres -c "drop database contrail;"
psql -h localhost -U postgres -c "create database contrail;"
psql -h localhost -U root -d contrail < tools/init_psql.sql
