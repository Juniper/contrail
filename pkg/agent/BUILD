package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "agent",
    srcs = ["agent.go","backend.go","task.go","task_handler.go","watcher.go"],
    importpath = "github.com/Juniper/contrail/pkg/agent",
    deps = [
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
        "@com_github_pkg_errors//:go_default_library",
        "@in_gopkg_yaml_v2//:go_default_library",
        "@com_github_mattn_go_shellwords//:go_default_library",
        "@com_github_joho_godotenv//:go_default_library",
        "@com_github_flosch_pongo2//:go_default_library",
    ],
)
