#!/usr/bin/env bash

reset_psql() {
    drop_psql
    create_psql
    initialize_psql
}

drop_psql(){
    echo "Dropping database contrail_test"
    docker exec contrail_psql bash -c \
        "PGPASSWORD=contrail123 psql -Uroot -d postgres -c \"drop database if exists contrail_test;\""
}

create_psql(){
    echo "Creating new database"
    docker exec contrail_psql bash -c \
        "PGPASSWORD=contrail123 psql -Uroot -d postgres -c \"create database contrail_test;\""
}

initialize_psql(){
    echo "Initializing database with gen_init_psql.sql"
    docker exec contrail_psql bash -c \
        "PGPASSWORD=contrail123 psql -Uroot -d contrail_test -q --file /tools/gen_init_psql.sql"
    docker exec contrail_psql bash -c \
        "PGPASSWORD=contrail123 psql -Uroot -d contrail_test -q --file /tools/init_psql.sql"
}

reset_psql
