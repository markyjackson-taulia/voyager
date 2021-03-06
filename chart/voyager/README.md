# Voyager
[Voyager by AppsCode](https://github.com/appscode/voyager) - Secure Ingress Controller for Kubernetes

## TL;DR;

```bash
$ helm install chart/voyager
```

## Introduction

This chart bootstraps an [ingress controller](https://github.com/appscode/voyager) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.


## Prerequisites

- Kubernetes 1.3+

## Installing the Chart
To install the chart with the release name `my-release`:
```bash
$ helm install --name my-release chart/voyager
```
The command deploys Voyager Controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release`:

```bash
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following tables lists the configurable parameters of the Voyager chart and their default values.


| Parameter              | Description                                                   | Default            |
| ---------------------- | ------------------------------------------------------------- | ------------------ |
| `image`                |  Container image to run                                       | `appscode/voyager` |
| `imageTag`             |  Image tag of container                                       | `3.1.1`            |
| `cloudProvider`        |  Name of cloud provider                                       | ``                 |
| `cloudConfig`          |  Path to cloud config                                         | ``                 |
| `logLevel`             |  Log level for operator                                       | `3`                |
| `persistence.enabled`  |  Enable mounting cloud config                                 | `false`            |
| `persistence.hostPath` |  Host mount path for cloud config                             | `/etc/kubernetes`  |
| `rbac.install`         | install required rbac service account, roles and rolebindings | `false`            |
| `rbac.apiVersion`      | rbac api version `v1alpha1|v1beta1`                           | `v1beta1`          |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```bash
$ helm install --name my-release --set image.tag=v0.2.1 chart/voyager
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm install --name my-release --values values.yaml chart/voyager
```

## RBAC
By default the chart will not install the recommended RBAC roles and rolebindings.

You need to have the following parameter on the api server. See the following document for how to enable [RBAC](https://kubernetes.io/docs/admin/authorization/rbac/)

```
--authorization-mode=RBAC
```

To determine if your cluster supports RBAC, run the the following command:

```console
$ kubectl api-versions | grep rbac
```

If the output contains "alpha" and/or "beta", you can may install the chart with RBAC enabled (see below).

### Enable RBAC role/rolebinding creation

To enable the creation of RBAC resources (On clusters with RBAC). Do the following:

```console
$ helm install --name my-release chart/voyager --set rbac.install=true
```

### Changing RBAC manifest apiVersion

By default the RBAC resources are generated with the "v1beta1" apiVersion. To use "v1alpha1" do the following:

```console
$ helm install --name my-release chart/voyager --set rbac.install=true,rbac.apiVersion=v1alpha1
```
