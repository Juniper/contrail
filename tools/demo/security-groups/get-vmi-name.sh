#!/bin/bash

if [ -z "$1" ]; then
	echo "Usage: get-vmi-name.sh POD_NAME"
	exit 1
fi

POD_NAME=$1
curl "localhost:8082/virtual-machine-interfaces" | python -m json.tool | sed -n -e "s/^.*\"\(${POD_NAME}__.*\)\".*$/\1/p" | head -n 1
