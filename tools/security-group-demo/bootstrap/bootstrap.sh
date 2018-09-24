#!/bin/bash
set -x

if [ -z "$1" ] || [ -z "$2" ] || [ -z "$3" ]; then
	echo "Usage: bootstrap.sh CONTRAIL_ADDR NAMESPACE DIRECTORY"
	exit 1
fi

CONTRAIL_ADDR=$1
NAMESPACE=$2
DIRECTORY=$3
PROJECT="k8s-$NAMESPACE"

bootstrap() {
	sed "s/{{addr}}/$CONTRAIL_ADDR/; s/{{proj-name}}/$PROJECT/; s/{{ns-name}}/$NAMESPACE/;" "$1.tmpl" > "$DIRECTORY/$1" && chmod +x "$DIRECTORY/$1"
}

mkdir -p "$DIRECTORY"
bootstrap "run.sh"
bootstrap "setup.sh"
bootstrap "pod.sh"
bootstrap "get-vmi-name.sh"
