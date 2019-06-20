package main

import (
	"flag"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/weiqiang333/infra-prometheus-exporter/hadoop"
	"github.com/weiqiang333/infra-prometheus-exporter/hadoop/model"
)

const (
	namespace = "regionserver"
	jvm       = "jvm"
)

func jvmset(e *model.JvmMetrics, jvmDataMap map[string]interface{}) {
	e.MemNonHeapUsedM.Set(jvmDataMap["MemNonHeapUsedM"].(float64))
	e.MemNonHeapCommittedM.Set(jvmDataMap["MemNonHeapCommittedM"].(float64))
	e.MemNonHeapMaxM.Set(jvmDataMap["MemNonHeapMaxM"].(float64))
	e.MemHeapUsedM.Set(jvmDataMap["MemHeapUsedM"].(float64))
	e.MemHeapCommittedM.Set(jvmDataMap["MemHeapCommittedM"].(float64))
	e.MemHeapMaxM.Set(jvmDataMap["MemHeapMaxM"].(float64))
	e.MemMaxM.Set(jvmDataMap["MemMaxM"].(float64))
	e.ThreadsNew.Set(jvmDataMap["ThreadsNew"].(float64))
	e.ThreadsRunnable.Set(jvmDataMap["ThreadsRunnable"].(float64))
	e.ThreadsBlocked.Set(jvmDataMap["ThreadsBlocked"].(float64))
	e.ThreadsWaiting.Set(jvmDataMap["ThreadsWaiting"].(float64))
	e.ThreadsTimedWaiting.Set(jvmDataMap["ThreadsTimedWaiting"].(float64))
	e.ThreadsTerminated.Set(jvmDataMap["ThreadsTerminated"].(float64))
	e.GcCount.Set(jvmDataMap["GcCount"].(float64))
	e.GcTimeMillis.Set(jvmDataMap["GcTimeMillis"].(float64))
	e.LogFatal.Set(jvmDataMap["LogFatal"].(float64))
	e.LogError.Set(jvmDataMap["LogError"].(float64))
	e.LogWarn.Set(jvmDataMap["LogWarn"].(float64))
	e.LogInfo.Set(jvmDataMap["LogInfo"].(float64))
}

func main() {
	flag.Parse()
	if hadoop.Help {
		flag.Usage()
		os.Exit(0)
	}

	exporter := model.NewJvmMetricsExporter(namespace, jvm)

	envelope, err := hadoop.GetJmxEnvelope()
	if err != nil {
		return
	}

	for _, hmasterDataMap := range envelope.Beans {
		switch {
		case strings.Contains(hmasterDataMap["name"].(string), hadoop.JvmMetrics):
			jvmset(exporter, hmasterDataMap)
		}
	}

	prometheus.MustRegister(exporter)
	hadoop.Export(namespace)
}
