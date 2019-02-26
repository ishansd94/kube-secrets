#!/usr/bin/env bash
set -ex

DOCKER_BUILDKIT=1 docker build -t ${USERNAME}/${IMAGE} \
 --build-arg APP=${APP} --build-arg EXECUTABLE=${EXECUTABLE} --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" .