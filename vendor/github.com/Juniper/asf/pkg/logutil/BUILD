package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "asf_logutils",
    srcs = ["logger.go"],
    importpath = "github.com/Juniper/asf/pkg/logutil",
    deps = [
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_mattn_go_shellwords//:go_default_library",
        "@com_github_yudai_gotty//server:go_default_library",
        "@com_github_yudai_gotty//utils:go_default_library",
        "@com_github_yudai_gotty//backend/localcommand:go_default_library",
    ],
)

