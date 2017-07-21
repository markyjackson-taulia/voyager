```console
$ kubectl create -f ./docs/examples/monitoring/demo-0.yaml
namespace "demo" created
deployment "prometheus-operator" created

$ kubectl get pods -n demo
NAME                                  READY     STATUS    RESTARTS   AGE
prometheus-operator-449376836-mq4p3   1/1       Running   0          1m
```

```console
$ kubectl create -f ./docs/examples/monitoring/demo-1.yaml
prometheus "prometheus" created
service "prometheus" created

$ kubectl get pods -n demo --watch
NAME                                  READY     STATUS    RESTARTS   AGE
prometheus-operator-449376836-mq4p3   1/1       Running   0          1m
prometheus-prometheus-0   0/2       ContainerCreating   0         6s
prometheus-prometheus-0   1/2       Running   0         25s
prometheus-prometheus-0   2/2       Running   0         26s
^C⏎
```

```console
$ ./hack/deploy/minikube.sh

$ kubectl get pods -l app=voyager --all-namespaces --watch
NAMESPACE     NAME                                READY     STATUS    RESTARTS   AGE
kube-system   voyager-operator-2464855905-gfdlq   1/1       Running   0          22s
^C⏎
```



```console
$ kubectl delete deployment voyager-operator -n kube-system
deployment "voyager-operator" deleted
$ kubectl delete svc voyager-operator -n kube-system
service "voyager-operator" deleted
```
