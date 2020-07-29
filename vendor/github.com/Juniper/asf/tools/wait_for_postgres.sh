#!/bin/bash

until docker-compose exec postgres psql -c "\l" asf_test; do
    echo >&2 "$(date +%Y%m%dt%H%M%S) Command failed - sleeping"
    sleep 1
done
