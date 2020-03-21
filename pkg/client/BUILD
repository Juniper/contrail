package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "client",
    srcs = ["neutron.go"],
    importpath = "github.com/Juniper/contrail/pkg/client",
    deps = [
        "@com_github_pkg_errors//:go_default_library",
    ],
)
