all: deps generate test

deps:
	./tools/install_proto.sh

generate:
	go generate ./...

# TODO(dfurman): use go test ./...
# TODO(dfurman): remove or fix the package compilation and tests for following packages:
# cmd/contrailcli
# cmd/contrailutil
# pkg/bench
# pkg/cmd/contrailcli
# pkg/cmd/contrailutil
# pkg/db/cache
# pkg/convert
# pkg/db
# pkg/db/cassandra
# pkg/db/etcd
# pkg/endpoint
# pkg/models
# pkg/rbac TODO(dfurman): this definitely needs to be fixed
# pkg/services
# pkg/sync
# pkg/sync/replication
# pkg/testutil/integration
# pkg/testutil/integration/etcd

TESTPKGS=./cmd/contrailschema \
	./pkg/apiserver \
	./pkg/auth \
	./pkg/client \
	./pkg/cmd/contrailschema \
	./pkg/config \
	./pkg/db/basedb \
	./pkg/errutil \
	./pkg/fileutil \
	./pkg/fileutil/template \
	./pkg/format \
	./pkg/httputil \
	./pkg/keystone \
	./pkg/logutil \
	./pkg/models/basemodels \
	./pkg/osutil \
	./pkg/proxy \
	./pkg/retry \
	./pkg/schema \
	./pkg/services/baseservices \
	./pkg/testutil

test:
	# TODO(mblotniak): After fixing the package compilation and tests please add it to $$TESTPKGS above.
	go test $(TESTPKGS)

.PHONY: test generate
