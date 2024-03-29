#!/bin/bash

# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

set -o nounset
set -o pipefail

TPRT_REPO_PATH="$(git rev-parse --show-toplevel)"

docker run --rm -t cytopia/yamllint --version
CONTAINER_NAME="tf_yamllint_$RANDOM"
docker run --name ${CONTAINER_NAME} -t -v "${TPRT_REPO_PATH}":/tanzu-plugin-runtime:ro cytopia/yamllint -s -c /tanzu-plugin-runtime/hack/check/.yamllintconfig.yaml /tanzu-plugin-runtime
EXIT_CODE=$(docker inspect ${CONTAINER_NAME} --format='{{.State.ExitCode}}')
docker rm -f ${CONTAINER_NAME} &> /dev/null

if [[ ${EXIT_CODE} == "0" ]]; then
  echo "yamllint passed!"
else
  echo "yamllint exit code ${EXIT_CODE}: YAML linting failed!"
  echo "Please fix the listed yamllint errors and verify using 'make yamllint'"
  exit "${EXIT_CODE}"
fi
