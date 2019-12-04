all: deps generate test

deps:
	./tools/install_proto.sh

generate:
	go generate ./...

TESTPKGS=$(shell go list ./... | grep -vE "cmd/contrail|pkg/apisrv|pkg/bench|pkg/cmd|pkg/convert|pkg/db|pkg/endpoint|pkg/models|pkg/rbac|pkg/replication|pkg/services|pkg/sync|pkg/testutil/integration") ./pkg/services/baseservices ./pkg/models/basemodels ./cmd/contrailschema ./pkg/cmd/contrailschema ./pkg/db/basedb

test:
	# TODO(mblotniak): After fixing the package compilation and tests please add it to $$TESTPKGS above.
	go test $(TESTPKGS)

.PHONY: test generate
