workspace(name = "distroless_sh_bug")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")

go_repositories()

git_repository(
    name = "distroless",
    remote = "https://github.com/GoogleCloudPlatform/distroless.git",
    commit = "9ea474449e4e22aaa952cc1f253c52880c0cc220",
)

# Required for the @distroless external to work correctly.
git_repository(
    name = "runtimes_common",
    commit = "b91f0dedf1ff6c53899b0f64d29993c3c65b73b4",
    remote = "https://github.com/GoogleCloudPlatform/runtimes-common.git",
)

# Required to build //hello:use_ext with the @distroless external repository.
load(
    "@distroless//package_manager:package_manager.bzl",
    "package_manager_repositories",
    "dpkg_src",
    "dpkg",
)

package_manager_repositories()

dpkg_src(
    name = "debian_jessie",
    arch = "amd64",
    distro = "jessie",
    url = "http://deb.debian.org",
)

dpkg_src(
    name = "debian_jessie_backports",
    arch = "amd64",
    distro = "jessie-backports",
    url = "http://deb.debian.org",
)

# For the glibc base image.
dpkg(
    name = "libc6",
    source = "@debian_jessie//file:Packages.json",
)

dpkg(
    name = "ca-certificates",
    source = "@debian_jessie//file:Packages.json",
)

dpkg(
    name = "openssl",
    source = "@debian_jessie//file:Packages.json",
)

dpkg(
    name = "libssl1.0.0",
    source = "@debian_jessie//file:Packages.json",
)

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
