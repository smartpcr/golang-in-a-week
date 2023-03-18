package main

import (
	"context"
	"flag"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"math/rand"
	"net/http"
	"operator/pkg/runtime"
	"operator/pkg/subscription"
	"time"
)

var (
	minWatchTimeout = 5 * time.Minute
	timeoutSeconds  = int64(minWatchTimeout.Seconds() * (rand.Float64() + 1.0))
	masterURL       string
	kubeconfig      string
	addr            = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

func main() {
	flag.Parse()

	// metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(*addr, nil)
	}()

	// tracing
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"operator",
		config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	// logs
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	// run ...
	log.Info("Got watcher client...")

	kubeCfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	log.Info("Building config from flags...")

	defaultKubernetesClient, err := kubernetes.NewForConfig(kubeCfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// context
	context := context.TODO()

	configMapSubscription := &subscription.ConfigMapSubscription{
		ClientSet:  defaultKubernetesClient,
		Ctx:        context,
		Completion: make(chan bool),
	}
	podSubscription := &subscription.PodSubscription{
		ClientSet:                defaultKubernetesClient,
		Ctx:                      context,
		Completion:               make(chan bool),
		ConfigMapSubscriptionRef: configMapSubscription,
	}

	if err := runtime.RunLoop([]subscription.ISubscription{
		configMapSubscription,
		podSubscription,
	}); err != nil {
		log.Fatalf("Error running loop: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
