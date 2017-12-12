
#!/usr/bin/env bash

TOP=$(dirname "$0")

go test -race -cover $(go list ./... | grep -v /vendor/)