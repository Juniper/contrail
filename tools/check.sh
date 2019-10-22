#!/usr/bin/env bash

go mod tidy
[[ -z $(git status --porcelain -- go.mod) ]] || { echo 'go.mod file has changed. Please update the go.mod file and commit it into repository.'; git diff -- go.mod; exit 1; }
