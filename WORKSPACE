workspace(name = "distroless_sh_bug")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")

go_repositories()

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    commit = "df9d21334be0d45b6995f5f46024a3d2ea22eca9",
)

load(
  "@io_bazel_rules_docker//docker:docker.bzl",
  "docker_repositories", "docker_pull"
)
docker_repositories()

docker_pull(
    name = "base",
    registry = "gcr.io",
    repository = "distroless/base",
    # 'tag' is also supported, but digest is encouraged for reproducibility.
    digest = "sha256:06fcd3edcfeefe13b82fa8bdb9e3f4fa3bf4c7e8fe997bee0230e392f77d0e04",
)
