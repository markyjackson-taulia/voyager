package main

import (
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	hpe "github.com/appscode/haproxy_exporter/exporter"
	"github.com/appscode/pat"
	"github.com/appscode/voyager/api"
	"github.com/orcaman/concurrent-map"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/tamalsaha/go-oneliners"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PathParamAPIGroup  = ":apiGroup"
	PathParamNamespace = ":namespace"
	PathParamName      = ":name"
	QueryParamPodIP    = "pod"
)

var (
	selectedServerMetrics map[int]*prometheus.GaugeVec

	registerers = cmap.New() // URL.path => *prometheus.Registry
)

func DeleteRegistry(w http.ResponseWriter, r *http.Request) {
	registerers.Remove(r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func ExportMetrics(w http.ResponseWriter, r *http.Request) {
	params, found := pat.FromContext(r.Context())
	if !found {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}
	apiGroup := params.Get(PathParamAPIGroup)
	if apiGroup == "" {
		http.Error(w, "Missing parameter:"+PathParamAPIGroup, http.StatusBadRequest)
		return
	}
	namespace := params.Get(PathParamNamespace)
	if namespace == "" {
		http.Error(w, "Missing parameter:"+PathParamNamespace, http.StatusBadRequest)
		return
	}
	name := params.Get(PathParamName)
	if name == "" {
		http.Error(w, "Missing parameter:"+PathParamName, http.StatusBadRequest)
		return
	}
	podIP := r.URL.Query().Get(QueryParamPodIP)
	if podIP == "" {
		podIP = "127.0.0.1"
	}
	oneliners.FILE("apiGroup", apiGroup, "namespace", namespace, "name", name, "podIP", podIP)

	switch apiGroup {
	case "extensions":
		var reg *prometheus.Registry
		if val, ok := registerers.Get(r.URL.Path); ok {
			reg = val.(*prometheus.Registry)
		} else {
			reg = prometheus.NewRegistry()
			if absent := registerers.SetIfAbsent(r.URL.Path, reg); !absent {
				r2, _ := registerers.Get(r.URL.Path)
				reg = r2.(*prometheus.Registry)
			} else {
				log.Infof("Configuring exporter for standard ingress %s in namespace %s", name, namespace)
				ingress, err := kubeClient.ExtensionsV1beta1().Ingresses(namespace).Get(name, metav1.GetOptions{})
				if kerr.IsNotFound(err) {
					http.NotFound(w, r)
					return
				} else if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				engress, err := api.NewEngressFromIngress(ingress)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				scrapeURL, err := getScrapeURL(engress, podIP)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				exporter, err := hpe.NewExporter(scrapeURL, selectedServerMetrics, haProxyTimeout)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				reg.MustRegister(exporter)
				reg.MustRegister(version.NewCollector("haproxy_exporter"))
			}
		}
		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
		return
	case api.GroupName:
		oneliners.FILE()
		var reg *prometheus.Registry
		if val, ok := registerers.Get(r.URL.Path); ok {
			oneliners.FILE()
			reg = val.(*prometheus.Registry)
		} else {
			oneliners.FILE()
			reg = prometheus.NewRegistry()
			if absent := registerers.SetIfAbsent(r.URL.Path, reg); !absent {
				oneliners.FILE()
				r2, _ := registerers.Get(r.URL.Path)
				reg = r2.(*prometheus.Registry)
			} else {
				log.Infof("Configuring exporter for appscode ingress %s in namespace %s", name, namespace)
				engress, err := extClient.Ingresses(namespace).Get(name)
				oneliners.FILE(err)
				if kerr.IsNotFound(err) {
					http.NotFound(w, r)
					return
				} else if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				oneliners.FILE()
				scrapeURL, err := getScrapeURL(engress, podIP)
				oneliners.FILE(err)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				oneliners.FILE(scrapeURL)
				exporter, err := hpe.NewExporter(scrapeURL, selectedServerMetrics, haProxyTimeout)
				oneliners.FILE(err)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				reg.MustRegister(exporter)
				reg.MustRegister(version.NewCollector("haproxy_exporter"))
			}
		}
		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}

func getScrapeURL(r *api.Ingress, podIP string) (string, error) {
	if !r.Stats() {
		return "", errors.New("Stats not exposed")
	}
	if r.StatsSecretName() == "" {
		return fmt.Sprintf("http://%s:%d?stats;csv", podIP, r.StatsPort()), nil
	}
	secret, err := kubeClient.CoreV1().Secrets(r.Namespace).Get(r.StatsSecretName(), metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	userName := string(secret.Data["username"])
	passWord := string(secret.Data["password"])
	return fmt.Sprintf("http://%s:%s@%s:%d?stats;csv", userName, passWord, podIP, r.StatsPort()), nil
}
