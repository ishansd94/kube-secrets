#!/usr/bin/env bash
set -ex

USERNAME=emzian7
IMAGE=kube-secrets

DOCKER_BUILDKIT=0 docker build -t ${USERNAME}/${IMAGE} --build-arg APP="github.com/ishansd94/kube-secrets" --build-arg EXECUTABLE="secret" .
docker push ${USERNAME}/${IMAGE}