package subscription

import (
	"context"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type PodSubscription struct {
	watchInterface           watch.Interface
	ClientSet                kubernetes.Interface
	Ctx                      context.Context
	Completion               chan bool
	ConfigMapSubscriptionRef *ConfigMapSubscription
}

func (p *PodSubscription) applyConfigMapChanges(pod *v1.Pod, event watch.EventType) {
	if pod == nil {
		return
	}
	if p.ConfigMapSubscriptionRef == nil {
		if p.ConfigMapSubscriptionRef.PlatformConfig == nil {
			updatedPod := pod.DeepCopy()
			if updatedPod.Annotations == nil {
				updatedPod.Annotations = make(map[string]string)
			}
			for _, annotation := range p.ConfigMapSubscriptionRef.PlatformConfig.Annotations {
				updatedPod.Annotations[annotation.Name] = annotation.Value
			}
			_, err := p.ClientSet.CoreV1().Pods(pod.Namespace).Update(p.Ctx, updatedPod, metav1.UpdateOptions{})
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (p *PodSubscription) Reconcile(object runtime.Object, event watch.EventType) error {
	rootSpan := opentracing.GlobalTracer().StartSpan("PodSubscription.Reconcile")
	defer rootSpan.Finish()

	pod := object.(*v1.Pod)
	log.WithFields(log.Fields{
		"pod":       pod.Name,
		"event":     event,
		"namespace": pod.Namespace,
	}).Info("PodSubscription event type %s for %s", event, pod.Name)

	switch event {
	case watch.Added:
		watchEventAdd := opentracing.GlobalTracer().StartSpan("PodSubscription.Reconcile.watchEventAdd", opentracing.ChildOf(rootSpan.Context()))
		defer watchEventAdd.Finish()
		p.applyConfigMapChanges(pod, event)
	case watch.Modified:
		watchEventModified := opentracing.GlobalTracer().StartSpan("PodSubscription.Reconcile.watchEventModified", opentracing.ChildOf(rootSpan.Context()))
		defer watchEventModified.Finish()
		p.applyConfigMapChanges(pod, event)
	case watch.Deleted:
		watchEventDeleted := opentracing.GlobalTracer().StartSpan("PodSubscription.Reconcile.watchEventDeleted", opentracing.ChildOf(rootSpan.Context()))
		defer watchEventDeleted.Finish()
		p.applyConfigMapChanges(pod, event)
	}
	return nil
}

func (p *PodSubscription) Subscribe() (watch.Interface, error) {
	var err error
	p.watchInterface, err = p.ClientSet.CoreV1().Pods("").Watch(p.Ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	log.Info("Started watch stream for PodSubscription")
	return p.watchInterface, nil
}

func (p *PodSubscription) IsComplete() <-chan bool {
	return p.Completion
}
