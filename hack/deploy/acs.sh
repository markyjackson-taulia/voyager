#!/bin/bash

export CLOUD_PROVIDER=azure
export CLOUD_CONFIG=/etc/kubernetes/azure.json
export INGRESS_CLASS=
cat ./hack/deploy/voyager-without-rbac.yaml | envsubst | kubectl apply -f -
