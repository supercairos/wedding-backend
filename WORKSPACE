# Load the http ruleset and expose the http_archive rule
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Download rules_go ruleset.
# Bazel makes a https call and downloads the zip file, and then
# checks the sha.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
    ],
)

# Download the bazel_gazelle ruleset.
http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

# Download the rules_docker ruleset.
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)

# Load rules_go ruleset and expose the toolchain and dep rules.
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# the line below instructs gazelle to save the go dependency definitions
# in the deps.bzl file. Located under '//'.
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

############################################################
# Define your own dependencies here using go_repository.   #
############################################################

# The following line defines the symbol go_dependencies from the deps.bzl file.
# Having the deps in that file, helps the WORKSPACE file stay less
# cluttered.  The library symbol go_dependencies is then added to
# the envionment. The line below calls that function.
load("//:deps.bzl", "go_dependencies")

# The next comment line includes a macro that gazelle reads.
# This macro tells Gazelle to look for repository rules in a macro in a .bzl file,
# and allows Gazelle to find the correct file to maintain the Go dependencies.
# Then the line after the comment calls go_dependencies(), and that funcation
# contains calls to various go_repository rules.

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

# go_rules_dependencies is a function that registers external dependencies
# needed by the Go rules.
# https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#go_rules_dependencies
go_rules_dependencies()

# The next rule installs the Go toolchains. The Go version is specified
# using the version parameter. This rule will download the Go SDK.
# https://github.com/bazelbuild/rules_go/blob/master/go/toolchains.rst#go_register_toolchains
go_register_toolchains(version = "1.19.5")

# The following call configured the gazelle dependencies, Go environment and Go SDK.
gazelle_dependencies()

############################################################
# Define your own dependencies here using go_repository.   #
############################################################

# load the container ruleset and expose the container_repositories rule
load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

# The container_repositories rule downloads the container dependencies.
container_repositories()

# load the go_image ruleset and expose the go_image_repositories rule
load("@io_bazel_rules_docker//go:image.bzl", go_image_repositories = "repositories")

# The go_image_repositories rule downloads the go_image dependencies.
go_image_repositories()
