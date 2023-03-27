#!/usr/bin/env bash
target=$1
bazel build --strip=never -c dbg "${target}"
outs=$(bazel cquery --strip=never -c dbg --output=files "${target}")
n=${#outs[@]}
if [[ "$n" -gt 1 ]]; then
  echo "too many outputs, not sure which one to debug (is this necessary?)"
fi
command="${outs[0]}"
dlv exec "${command}" --log --log-output=dap --headless --listen=127.0.0.1:50034 --api-version=2