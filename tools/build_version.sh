#!/bin/sh
cat > pkg/version/gen_version_build.go << EOF
package version

var (
	// BuildTime contains UTC timestamp
	BuildTime = "$(TZ=UTC date --rfc-3339=ns)"
	// BuildUser contains user name
	BuildUser = "$(whoami)"
	// BuildHostname contains host name
	BuildHostname = "$(hostname)"
	// BuildID contains build ID like "5.1.0-504.el7"
	BuildID = "todo" // TODO: need to replace with build id
	// BuildNumber contains package owner name
	BuildNumber = "@contrail"
)
EOF
go fmt pkg/version/gen_version_build.go
