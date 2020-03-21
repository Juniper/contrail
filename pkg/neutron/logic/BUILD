package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "logic",
    srcs = ["errors.go",
            "floating_ip.go", 
            "netutils.go", 
            "project.go", 
            "security_group.go", 
            "field_mask.go", 
            "port.go", 
            "request.go", 
            "security_group_rule.go", 
            "filters.go", 
            "host_prefixes.go", 
            "network.go", 
            "port_response.go", 
            "resource.go", 
            "uuid.go", 
            "network_response.go", 
            "router.go", 
            "subnet.go", 
    ],
    importpath = "github.com/Juniper/contrail/pkg/neutron/logic",
    deps = [
        "@com_github_pkg_errors//:go_default_library",
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_gogo_protobuf//types:go_default_library",
        "@com_github_satori_go_uuid//:go_default_library",
        "@com_github_apparentlymart_go_cidr//cidr:go_default_library",
    ],
)
