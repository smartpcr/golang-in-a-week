package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const ServiceName = "open-telemetry-test"

type Receiver struct {
	Endpoint string `yaml:"endpoint"`
}

type Configuration struct {
	Receiver Receiver `yaml:"receiver"`
}

var Config Configuration

func init() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	if _, err := os.Stat("./config.yaml"); os.IsNotExist(err) {
		log.Fatal("Config file does not exist!")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})

	//// Set default values
	//viper.SetDefault("loki.endpoint", "http://loki.local:3100/api/prom/push")
	//viper.SetDefault("prometheus.endpoint", "http://prometheus.local:9090")
	//viper.SetDefault("jaeger.trace_endpoint", "http://jaeger.local:14250/api/traces")
	//viper.SetDefault("jaeger.prometheus_endpoint", "http://jaeger.local:16686/metrics")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	// Print to verify
	fmt.Println("Open Telemetry Receiver Endpoint:", Config.Receiver.Endpoint)
}
