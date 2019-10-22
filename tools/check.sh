#!/usr/bin/env bash

go mod vendor
[[ -z $(git status --porcelain -- go.mod) ]] || { echo 'go.mod file has changed. Please update the go.mod file and commit it into repository.'; git diff -- go.mod; exit 1; }
[[ -z $(git status --porcelain -- vendor/) ]] || { echo "Project's vendor is not up to date. Please update it with 'go mod vendor' command."; exit 1; }
