#!/bin/bash

export CLOUD_PROVIDER=aws
export CLOUD_CONFIG=
export INGRESS_CLASS=
cat ./hack/deploy/voyager-without-rbac.yaml | envsubst | kubectl apply -f -
