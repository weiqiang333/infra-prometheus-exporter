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
	namespace = "resourcemanager"
	yarn      = "yarn"
	jvm       = "jvm"
	rpc       = "rpc"
)

// Exporter  需要在 prometheus 中注册
type Exporter struct {
	jvmMetrics     *model.JvmMetrics
	clusterMetrics *model.ClusterMetrics
	queueMetrics   *model.QueueMetrics
	rpc            *model.RPC
}

// NewExporter returns an initialized exporter.
func NewExporter() *Exporter {
	e := &Exporter{}
	e.jvmMetrics = model.NewJvmMetricsExporter(namespace, jvm)
	e.clusterMetrics = model.NewClusterMetricsExporter(namespace, yarn)
	e.queueMetrics = model.NewQueueMetricsExporter(namespace, yarn)
	e.rpc = model.NewRPCExporter(namespace, rpc)
	return e
}

// Describe describes all the metrics exported by the namenode exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.jvmMetrics.Describe(ch)
	e.clusterMetrics.Describe(ch)
	e.queueMetrics.Describe(ch)
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

		case nameDataMap["name"].(string) == hadoop.ResourceManagerJvmMetrics:
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
		case strings.Contains(nameDataMap["name"].(string), hadoop.ResourceManagerRPC):
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
		case nameDataMap["name"].(string) == hadoop.ResourceManagerClusterMetrics:
			e.clusterMetrics.NumActiveNMs.Set(nameDataMap["NumActiveNMs"].(float64))
			e.clusterMetrics.NumDecommissionedNMs.Set(nameDataMap["NumDecommissionedNMs"].(float64))
			e.clusterMetrics.NumLostNMs.Set(nameDataMap["NumLostNMs"].(float64))
			e.clusterMetrics.NumUnhealthyNMs.Set(nameDataMap["NumUnhealthyNMs"].(float64))
			e.clusterMetrics.NumRebootedNMs.Set(nameDataMap["NumRebootedNMs"].(float64))
		case strings.Contains(nameDataMap["name"].(string), hadoop.ResourceManagerQueueMetrics) && nameDataMap["tag.Queue"] != nil && nameDataMap["tag.User"] != nil:
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.Running_0, prometheus.GaugeValue, nameDataMap["running_0"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.Running_60, prometheus.GaugeValue, nameDataMap["running_60"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.Running_300, prometheus.GaugeValue, nameDataMap["running_300"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.Running_1440, prometheus.GaugeValue, nameDataMap["running_1440"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AppsSubmitted, prometheus.GaugeValue, nameDataMap["AppsSubmitted"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AppsRunning, prometheus.GaugeValue, nameDataMap["AppsRunning"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AppsPending, prometheus.GaugeValue, nameDataMap["AppsPending"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AppsCompleted, prometheus.GaugeValue, nameDataMap["AppsCompleted"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AppsKilled, prometheus.GaugeValue, nameDataMap["AppsKilled"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AppsFailed, prometheus.GaugeValue, nameDataMap["AppsFailed"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AllocatedMB, prometheus.GaugeValue, nameDataMap["AllocatedMB"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AllocatedVCores, prometheus.GaugeValue, nameDataMap["AllocatedVCores"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AllocatedContainers, prometheus.GaugeValue, nameDataMap["AllocatedContainers"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AggregateContainersAllocated, prometheus.GaugeValue, nameDataMap["AggregateContainersAllocated"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AggregateContainersReleased, prometheus.GaugeValue, nameDataMap["AggregateContainersReleased"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AvailableMB, prometheus.GaugeValue, nameDataMap["AvailableMB"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.AvailableVCores, prometheus.GaugeValue, nameDataMap["AvailableVCores"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.PendingMB, prometheus.GaugeValue, nameDataMap["PendingMB"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.PendingVCores, prometheus.GaugeValue, nameDataMap["PendingVCores"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.PendingContainers, prometheus.GaugeValue, nameDataMap["PendingContainers"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.ReservedMB, prometheus.GaugeValue, nameDataMap["ReservedMB"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.ReservedVCores, prometheus.GaugeValue, nameDataMap["ReservedVCores"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.ReservedContainers, prometheus.GaugeValue, nameDataMap["ReservedContainers"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.ActiveUsers, prometheus.GaugeValue, nameDataMap["ActiveUsers"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))
			ch <- prometheus.MustNewConstMetric(e.queueMetrics.ActiveApplications, prometheus.GaugeValue, nameDataMap["ActiveApplications"].(float64), nameDataMap["tag.Queue"].(string), nameDataMap["tag.User"].(string))

		}
	}

	e.jvmMetrics.Collect(ch)
	e.clusterMetrics.Collect(ch)
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
