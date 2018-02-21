#!/usr/bin/env bash

goimports -w -d $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*pkg/generated/models/*") &> /dev/null