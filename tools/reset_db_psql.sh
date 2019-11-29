#!/usr/bin/env bash

reset_psql() {
    TOOLSDIR=$(dirname "$0")
    PROJECT="contrail"
    DOCKER_COMPOSE="docker-compose -f $TOOLSDIR/patroni/docker-compose.yml -p $PROJECT"
    SUCCESS_MSG="ERROR, SQLSTATE: no results to fetch"

    drop_psql_db
    create_psql_db
    initialize_psql_db
}

drop_psql_db() {
    echo "Dropping database contrail_test"
    response=$(NETWORKNAME=${PROJECT} ${DOCKER_COMPOSE} exec -T dbnode bash -c \
        "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"drop database if exists contrail_test;\" testcluster")
    if [[ "${response:20}" = "$SUCCESS_MSG" ]]; then
        echo "Dropped database contrail_test"
    else
        echo "Error while dropping database ${response:20}"
        exit 1
    fi
}

create_psql_db() {
    echo "Creating new database"
    response=$(NETWORKNAME=${PROJECT} ${DOCKER_COMPOSE} exec -T dbnode bash -c \
        "PGPASSWORD=contrail123 patronictl query -Uroot -d postgres -c \"create database contrail_test;\" testcluster")
    if [[ "${response:20}" = "$SUCCESS_MSG" ]]; then
        echo "Created database contrail_test"
    else
        echo "Error while creating database contrail_test ${response:20}"
        return 1
    fi
}

initialize_psql_db() {
    echo "Initializing database with gen_init_psql.sql"
    response=$(NETWORKNAME=${PROJECT} ${DOCKER_COMPOSE} exec -T dbnode bash -c \
        "PGPASSWORD=contrail123 patronictl query -Uroot -d contrail_test --file /tools/gen_init_psql.sql testcluster")
    if [[ "${response:20}" = "$SUCCESS_MSG" ]]; then
        echo "Initialized database contrail_test with gen_init_psql.sql"
    else
        echo "Error while initializing database contrail_test with gen_init_psql.sql: ${response:20}"
        return 1
    fi

    echo "Initializing database with init_psql.sql"
    response=$(NETWORKNAME=${PROJECT} ${DOCKER_COMPOSE} exec -T dbnode bash -c \
        "PGPASSWORD=contrail123 patronictl query -Uroot -d contrail_test --file /tools/init_psql.sql testcluster")
    if [[ "${response:20}" = "$SUCCESS_MSG" ]]; then
        echo "Initialized database contrail_test with init_psql.sql"
    else
        echo "Error while initializing database contrail_test with init_psql.sql: ${response:20}"
        return 1
    fi
}

reset_psql
