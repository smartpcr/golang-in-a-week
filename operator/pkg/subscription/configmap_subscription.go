package subscription

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

var (
	platformConfigMapName                   = "platform-default-configmap"
	platformConfigMapNamespace              = "kube-system"
	prometheusPlatformConfigAnnotationCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "cnskunkworks_platform_config_annotation_count",
		Help: "This tell us the number of annotations with  the configmap",
	})
	prometheusPlatformConfigAvailibilityGuage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cnskunkworks_platformconfig_availibity",
		Help: "This tells whether a platform config is available",
	},
		[]string{"configmap_name", "namespace"})
)

type ConfigMapSubscription struct {
	watcherInterface    watch.Interface
	ClientSet           kubernetes.Interface
	Ctx                 context.Context
	Completion          chan bool
	PlatformConfig      *platformConfig
	platformConfigPhase string
}

type platformAnnotation struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type platformConfig struct {
	Annotations []platformAnnotation `yaml:"annotations"`
}

func isPlatformConfigMap(configMap *v1.ConfigMap) (bool, error) {
	if configMap == nil {
		return false, errors.New("configmap is nil")
	}
	if configMap.Name == platformConfigMapName && configMap.Namespace == platformConfigMapNamespace {
		return true, nil
	}
	return false, nil
}

func (c *ConfigMapSubscription) Reconcile(object runtime.Object, event watch.EventType) error {
	rootSpan := opentracing.GlobalTracer().StartSpan("ConfigMapSubscription.Recompile")
	defer rootSpan.Finish()

	configMap, ok := object.(*v1.ConfigMap)
	if !ok {
		return errors.New("object is not a configmap")
	}

	isPlatformConfigSpan := opentracing.GlobalTracer().StartSpan(
		"ConfigMapSubscription.Recompile.isPlatformConfig", opentracing.ChildOf(rootSpan.Context()))
	defer isPlatformConfigSpan.Finish()
	if ok, err := isPlatformConfigMap(configMap); !ok {
		if err != nil {
			klog.Error(err)
		}
		return err
	}

	klog.Info("ConfigMapSubscription event type %s for %s", event, configMap.Name)
	switch event {
	case watch.Added:
		watchEventAdd := opentracing.GlobalTracer().StartSpan(
			"ConfigMapSubscription.Recompile.watchEventAdd", opentracing.ChildOf(rootSpan.Context()))
		defer watchEventAdd.Finish()
		c.platformConfigPhase = string(event)
		rawDefaultsString := configMap.Data["platform-defaults"]
		var unmarshalledData platformConfig
		err := yaml.Unmarshal([]byte(rawDefaultsString), &unmarshalledData)
		if err != nil {
			klog.Error(err)
			return err
		}
		c.PlatformConfig = &unmarshalledData
		prometheusPlatformConfigAvailibilityGuage.WithLabelValues(configMap.Name, configMap.Namespace).Set(float64(1))
		prometheusPlatformConfigAnnotationCount.Set(float64(len(configMap.Annotations)))
	case watch.Deleted:
		watchEventDelete := opentracing.GlobalTracer().StartSpan(
			"ConfigMapSubscription.Recompile.watchEventDelete", opentracing.ChildOf(rootSpan.Context()))
		defer watchEventDelete.Finish()
		c.platformConfigPhase = string(event)
		c.PlatformConfig = nil
		prometheusPlatformConfigAvailibilityGuage.WithLabelValues(configMap.Name, configMap.Namespace).Set(float64(0))
		prometheusPlatformConfigAnnotationCount.Set(0)
	case watch.Modified:
		watchEventModify := opentracing.GlobalTracer().StartSpan(
			"ConfigMapSubscription.Recompile.watchEventModify", opentracing.ChildOf(rootSpan.Context()))
		defer watchEventModify.Finish()
		c.platformConfigPhase = string(event)
		rawDefaultsString := configMap.Data["platform-defaults"]
		var unmarshalledData platformConfig
		err := yaml.Unmarshal([]byte(rawDefaultsString), &unmarshalledData)
		if err != nil {
			klog.Error(err)
			return err
		}
		c.PlatformConfig = &unmarshalledData
		prometheusPlatformConfigAvailibilityGuage.WithLabelValues(configMap.Name, configMap.Namespace).Set(float64(1))
		prometheusPlatformConfigAnnotationCount.Set(float64(len(configMap.Annotations)))
	}
	return nil
}

func (c *ConfigMapSubscription) Subscribe() (watch.Interface, error) {
	var err error
	c.watcherInterface, err = c.ClientSet.CoreV1().ConfigMaps(platformConfigMapNamespace).Watch(c.Ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	log.Info("Started watch stream for ConfigMapSubscription")
	return c.watcherInterface, nil
}

func (c *ConfigMapSubscription) Unsubscribe() {
	c.watcherInterface.Stop()
}

func (c *ConfigMapSubscription) IsComplete() <-chan bool {
	return c.Completion
}
