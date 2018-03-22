#!/usr/bin/env bash

TOP=$(dirname "$0")

psql -h localhost -U postgres -c "drop database contrail_test"
psql -h localhost -U postgres -c "create database contrail_test"
psql -h localhost -U postgres contrail_test < $TOP/init_psql.sql