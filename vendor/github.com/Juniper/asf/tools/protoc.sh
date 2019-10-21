#!/bin/bash

# Wrapper script for protoc. If the binary is missing it runs the insallation script.

BASEDIR=$(dirname "$0")
PROTOC_BIN=$BASEDIR/../bin/protoc
[[ -f "$PROTOC_BIN" ]] || $BASEDIR/install_proto.sh

exec $PROTOC_BIN "$@"
