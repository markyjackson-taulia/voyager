#!/bin/bash

export CLOUD_PROVIDER=minikube
export CLOUD_CONFIG=
cat ./hack/deploy/voyager-without-rbac.yaml | envsubst | kubectl apply -f -
