load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
  
http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.22.2/rules_go-v0.22.2.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.22.2/rules_go-v0.22.2.tar.gz",
    ],
    sha256 = "142dd33e38b563605f0d20e89d9ef9eda0fc3cb539a14be1bdb1350de2eda659",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.20.0/bazel-gazelle-v0.20.0.tar.gz",
    ],
    sha256 = "d8c45ee70ec39a57e7a05e5027c32b1576cc7f16d9dd37135b0eddde45cf1b10",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf/protoc-gen-gogofaster",
    tag = "v1.3.1",
)
go_repository(
    name = "com_github_pseudomuto_protoc_gen_doc",
    importpath = "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc",
    tag = "v1.3.1",
)
go_repository(
    name = "com_github_golang_mockgen",
    importpath = "github.com/golang/mock/mockgen",
    tag = "v1.4.3",
)
go_repository(
    name = "com_github_mitchellh_gox",
    importpath = "github.com/mitchellh/gox",
    tag = "v1.0.1",
)
go_repository(
    name = "com_github_golang_x_tool_cmd_goimports",
    importpath = "golang.org/x/tools/cmd/goimports",
    tag = "gopls/v0.3.4",
)
go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.4.2"
)
go_repository(
    name = "com_github_mattn_go_shellwords",
    importpath = "github.com/mattn/go-shellwords",
    tag = "v1.0.10"
)
go_repository(
    name = "com_github_yudai_gotty",
    importpath = "github.com/yudai/gotty",
    tag = "v2.0.0-alpha.3",
)
go_repository(
    name ="com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    tag = "v0.0.6"
)
go_repository(
    name ="com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    tag = "v1.6.2"
)
go_repository(
    name ="com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.9"
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    tag = "v1.0.5"
)
go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    commit = "914dc3f8dd7c463188c73fc47e9ced82a6e421ca"
)
go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    tag = "v1.2.2"
)
go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    tag = "v1.3.1"
)
go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    tag = "v1.55.0"
)
go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.8"
)
go_repository(
    name = "com_github_subosito_gotenv",
    importpath = "github.com/subosito/gotenv",
    tag = "v1.2.0"
)
go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    tag = "v1.1.0"
)
go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    tag = "v1.8.1"
)
go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    tag = "v1.2.2"
)
go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    commit = "06d492aade888ab8698aad35476286b7b555c961"
)
go_repository(
    name = "org_golang_x_text_transform",
    importpath = "golang.org/x/text",
    commit = "06d492aade888ab8698aad35476286b7b555c961"
)
go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    tag = "v0.9.1"
)
go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    commit = "d6ee6871f21dd95e76563a90f522ce9fe75443f8"
)
go_repository(
    name = "com_github_flosch_pongo2",
    importpath = "github.com/flosch/pongo2",
    commit = "bbf5a6c351f4d4e883daa40046a404d7553e0a00",
)
go_repository(
    name = "com_github_juju_errors",
    importpath = "github.com/juju/errors",
    commit = "d42613fe1ab9e303fc850e7a19fda2e8eeb6516e",
)
go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    tag	= "v1.2.0",
)
go_repository(
    name = "com_github_apparentlymart_go_cidr",
    importpath = "github.com/apparentlymart/go-cidr",
    tag = "v1.0.1"
)
http_archive(
    name = "rules_proto",
    strip_prefix = "rules_proto-218ffa7dfa5408492dc86c01ee637614f8695c45",
    urls = ["https://github.com/bazelbuild/rules_proto/archive/218ffa7dfa5408492dc86c01ee637614f8695c45.tar.gz",],
    sha256 = "2490dca4f249b8a9a3ab07bd1ba6eca085aaf8e45a734af92aad0c42d9dc7aaf"
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()
rules_proto_toolchains()

