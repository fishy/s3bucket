workspace(name = "com_github_fishy_s3bucket")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.8.1/rules_go-0.8.1.tar.gz",
    sha256 = "90bb270d0a92ed5c83558b2797346917c46547f6f7103e648941ecdb6b9d0e72",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")
go_rules_dependencies()
go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.8/bazel-gazelle-0.8.tar.gz",
    sha256 = "e3dadf036c769d1f40603b86ae1f0f90d11837116022d9b06e4cd88cae786676",
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

go_repository(
    name = "com_github_fishy_fsdb",
    importpath = "github.com/fishy/fsdb",
    commit = "206200ff9529d88b9b00ecc4a970e0537c09eed1",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    commit = "f98ff3505c17e44d349af899eb85e6b37a7953db",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    commit = "dd801d4f4ce7ac746e7e7b4489d2fa600b3b096b",
)

go_repository(
    name = "com_github_go_ini_ini",
    importpath = "github.com/go-ini/ini",
    commit = "32e4c1e6bc4e7d0d8451aa6b75200d19e37a536a",
)
