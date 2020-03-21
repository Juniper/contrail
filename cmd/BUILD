load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "contrailschema",
    srcs = ["contrail/main.go"],
    deps = ["//vendor/github.com/Juniper/asf/pkg/logutil:asf_logutils"]
)
