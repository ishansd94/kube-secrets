#!/usr/bin/env bash
set -ex


export USERNAME=emzian7
export IMAGE=kube-secrets
export APP="github.com/ishansd94/kube-secrets"
export EXECUTABLE="secrets"


./hack/build.sh

./hack/test.sh

./hack/deploy.sh
