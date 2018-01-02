load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/fishy/s3bucket",
)

go_library(
    name = "go_default_library",
    srcs = ["s3.go"],
    importpath = "github.com/fishy/s3bucket",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_aws_aws_sdk_go//aws:go_default_library",
        "@com_github_aws_aws_sdk_go//aws/awserr:go_default_library",
        "@com_github_aws_aws_sdk_go//aws/session:go_default_library",
        "@com_github_aws_aws_sdk_go//service/s3:go_default_library",
        "@com_github_aws_aws_sdk_go//service/s3/s3manager:go_default_library",
        "@com_github_fishy_fsdb//bucket:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    size = "small",
    srcs = ["s3_test.go"],
    importpath = "github.com/fishy/s3bucket_test",
    deps = [":go_default_library"],
)
