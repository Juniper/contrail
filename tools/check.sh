#!/usr/bin/env bash

go mod tidy
[[ -z $(git status --porcelain -- go.mod) ]] || { echo 'go.mod file has changed. Please update the go.mod file and commit it into repository.'; git diff -- go.mod; exit 1; }
[[ -z $(git status --porcelain -- go.sum) ]] || { echo 'go.sum file has changed. Please update the go.sum file and commit it into repository.'; git diff -- go.sum; exit 1; }
