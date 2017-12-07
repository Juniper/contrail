
#!/usr/bin/env bash

TOP=$(dirname "$0")

./$TOP/reset_db.sh
go test -race -cover $(go list ./... | grep -v /vendor/)