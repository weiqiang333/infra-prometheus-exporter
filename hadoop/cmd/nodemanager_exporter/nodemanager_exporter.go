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
	namespace = "nodemanager"
	yarn      = "yarn"
	jvm       = "jvm"
	dfs       = "dfs"
	rpc       = "rpc"
)

// Exporter  需要在 prometheus 中注册
type Exporter struct {
	jvmMetrics  *model.JvmMetrics
	nodeManager *model.NodeManager
	rpc         *model.RPC
}

// NewExporter returns an initialized exporter.
func NewExporter() *Exporter {
	e := &Exporter{}
	e.jvmMetrics = model.NewJvmMetricsExporter(namespace, jvm)
	e.nodeManager = model.NewNodeManagerExporter(namespace, yarn)
	e.rpc = model.NewRPCExporter(namespace, rpc)
	return e
}

// Describe describes all the metrics exported by the namenode exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.jvmMetrics.Describe(ch)
	e.nodeManager.Describe(ch)
	e.rpc.Describe(ch)
}

// Collect fetches the statistics from the configured Namenode server, and
// delivers them as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	envelope, err := hadoop.GetJmxEnvelope()
	if err != nil {
		return
	}

	for _, nameDataMap := range envelope.Beans {

		switch {
		case nameDataMap["name"].(string) == hadoop.NodeManagerMetrics:
			e.nodeManager.ContainersLaunched.Set(nameDataMap["ContainersLaunched"].(float64))
			e.nodeManager.ContainersCompleted.Set(nameDataMap["ContainersCompleted"].(float64))
			e.nodeManager.ContainersFailed.Set(nameDataMap["ContainersFailed"].(float64))
			e.nodeManager.ContainersKilled.Set(nameDataMap["ContainersKilled"].(float64))
			e.nodeManager.ContainersIniting.Set(nameDataMap["ContainersIniting"].(float64))
			e.nodeManager.ContainersRunning.Set(nameDataMap["ContainersRunning"].(float64))
			e.nodeManager.AllocatedGB.Set(nameDataMap["AllocatedGB"].(float64))
			e.nodeManager.AllocatedContainers.Set(nameDataMap["AllocatedContainers"].(float64))
			e.nodeManager.AvailableGB.Set(nameDataMap["AvailableGB"].(float64))
			e.nodeManager.AllocatedVCores.Set(nameDataMap["AllocatedVCores"].(float64))
			e.nodeManager.AvailableVCores.Set(nameDataMap["AvailableVCores"].(float64))

		case nameDataMap["name"].(string) == hadoop.NodeManagerJvmMetrics:
			e.jvmMetrics.MemNonHeapUsedM.Set(nameDataMap["MemNonHeapUsedM"].(float64))
			e.jvmMetrics.MemNonHeapCommittedM.Set(nameDataMap["MemNonHeapCommittedM"].(float64))
			e.jvmMetrics.MemNonHeapMaxM.Set(nameDataMap["MemNonHeapMaxM"].(float64))
			e.jvmMetrics.MemHeapUsedM.Set(nameDataMap["MemHeapUsedM"].(float64))
			e.jvmMetrics.MemHeapCommittedM.Set(nameDataMap["MemHeapCommittedM"].(float64))
			e.jvmMetrics.MemHeapMaxM.Set(nameDataMap["MemHeapMaxM"].(float64))
			e.jvmMetrics.MemMaxM.Set(nameDataMap["MemMaxM"].(float64))
			e.jvmMetrics.ThreadsNew.Set(nameDataMap["ThreadsNew"].(float64))
			e.jvmMetrics.ThreadsRunnable.Set(nameDataMap["ThreadsRunnable"].(float64))
			e.jvmMetrics.ThreadsBlocked.Set(nameDataMap["ThreadsBlocked"].(float64))
			e.jvmMetrics.ThreadsWaiting.Set(nameDataMap["ThreadsWaiting"].(float64))
			e.jvmMetrics.ThreadsTimedWaiting.Set(nameDataMap["ThreadsTimedWaiting"].(float64))
			e.jvmMetrics.ThreadsTerminated.Set(nameDataMap["ThreadsTerminated"].(float64))
			e.jvmMetrics.GcCount.Set(nameDataMap["GcCount"].(float64))
			e.jvmMetrics.GcTimeMillis.Set(nameDataMap["GcTimeMillis"].(float64))
			e.jvmMetrics.LogFatal.Set(nameDataMap["LogFatal"].(float64))
			e.jvmMetrics.LogError.Set(nameDataMap["LogError"].(float64))
			e.jvmMetrics.LogWarn.Set(nameDataMap["LogWarn"].(float64))
			e.jvmMetrics.LogInfo.Set(nameDataMap["LogInfo"].(float64))
		case strings.Contains(nameDataMap["name"].(string), hadoop.NodeManagerRPC):
			ch <- prometheus.MustNewConstMetric(e.rpc.ReceivedBytes, prometheus.GaugeValue, nameDataMap["ReceivedBytes"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.SentBytes, prometheus.GaugeValue, nameDataMap["SentBytes"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcQueueTimeNumOps, prometheus.GaugeValue, nameDataMap["RpcQueueTimeNumOps"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcQueueTimeAvgTime, prometheus.GaugeValue, nameDataMap["RpcQueueTimeAvgTime"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcProcessingTimeNumOps, prometheus.GaugeValue, nameDataMap["RpcProcessingTimeNumOps"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcProcessingTimeAvgTime, prometheus.GaugeValue, nameDataMap["RpcProcessingTimeAvgTime"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcAuthenticationFailures, prometheus.GaugeValue, nameDataMap["RpcAuthenticationFailures"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcAuthenticationSuccesses, prometheus.GaugeValue, nameDataMap["RpcAuthenticationSuccesses"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcAuthorizationFailures, prometheus.GaugeValue, nameDataMap["RpcAuthorizationFailures"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.RpcAuthorizationSuccesses, prometheus.GaugeValue, nameDataMap["RpcAuthorizationSuccesses"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.NumOpenConnections, prometheus.GaugeValue, nameDataMap["NumOpenConnections"].(float64), nameDataMap["tag.port"].(string))
			ch <- prometheus.MustNewConstMetric(e.rpc.CallQueueLength, prometheus.GaugeValue, nameDataMap["CallQueueLength"].(float64), nameDataMap["tag.port"].(string))

		}
	}
	e.nodeManager.Collect(ch)
	e.jvmMetrics.Collect(ch)
}

func main() {
	flag.Parse()
	if hadoop.Help {
		flag.Usage()
		os.Exit(0)
	}
	exporter := NewExporter()
	prometheus.MustRegister(exporter)
	hadoop.Export(namespace)

}
