#!/bin/bash

export CLOUD_PROVIDER=gke
export CLOUD_CONFIG=
export INGRESS_CLASS=voyager
cat ./hack/deploy/voyager-without-rbac.yaml | envsubst | kubectl apply -f -
