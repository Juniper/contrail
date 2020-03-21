package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cmd_contrail",
    srcs = ["run.go"],
    importpath = "github.com/Juniper/contrail/pkg/cmd/contrail",
    deps = [
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
    ],
)
