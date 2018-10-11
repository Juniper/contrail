#!/usr/bin/env bash

# "dep ensure -dry-run" does not fail when inconsistencies are found
# TODO: use "dep check" when it passes on CI
dep ensure -dry-run
