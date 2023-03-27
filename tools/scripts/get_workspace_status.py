from __future__ import print_function
import subprocess
import sys


def execute_subprocess(path, args=[], ignore_errors=False):
    p = subprocess.Popen(args, cwd=path, stdout=subprocess.PIPE)
    (out, err) = p.communicate()
    if p.returncode != 0 and not ignore_errors:
        sys.exit(p.returncode)
    if out is None:
        return None
    return out.decode("ascii").strip()


def get_git_commit(path):
    return execute_subprocess(path, ["git", "rev-parse", "HEAD"])


def get_git_tag(path):
    return execute_subprocess(path, ["git", "describe", "--tags", "--exact-match", "HEAD"], True)
    # return execute_subprocess(path, ["git", "name-rev", "--tags", "--name-only", "$(git rev-parse HEAD)"])


def get_git_branch(path):
    return execute_subprocess(path, ["git", "rev-parse", "--abbrev-ref", "HEAD"])


def is_git_dirty(path):
    return execute_subprocess(path, ["git", "status", "-s"])


def main():
    git_tag = get_git_tag(".")
    git_commit = get_git_commit(".")
    git_branch = get_git_branch(".")
    build_tag = git_branch.replace("/", "_") if git_branch else "unknown"
    if git_tag:
        build_tag = git_tag
    if not git_tag:
        git_tag = "-"

    print("Setting STABLE_SCM_TAG = {}, STABLE_SCM_COMMIT = {}, STABLE_SCM_BRANCH = {} and STABLE_BUILD_TAG = {}".format(
        git_tag, git_commit, git_branch, build_tag
    ), file=sys.stderr)

    print("STABLE_SCM_TAG {}".format(git_tag))
    print("STABLE_SCM_COMMIT {}".format(git_commit))
    print("STABLE_SCM_BRANCH {}".format(git_branch))
    print("STABLE_BUILD_TAG {}".format(build_tag))


if __name__ == "__main__":
    main()
