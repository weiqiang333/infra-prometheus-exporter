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
	namespace = "journalnode"
	jvm       = "jvm"
	dfs       = "dfs"
	rpc       = "rpc"
)

// Exporter  需要在 prometheus 中注册
type Exporter struct {
	jvmMetrics  *model.JvmMetrics
	journalNode *model.JournalNode
	rpc         *model.RPC
}

// NewExporter returns an initialized exporter.
func NewExporter() *Exporter {
	e := &Exporter{}
	e.jvmMetrics = model.NewJvmMetricsExporter(namespace, jvm)
	e.journalNode = model.NewJournalNodeExporter(namespace, dfs)
	e.rpc = model.NewRPCExporter(namespace, rpc)
	return e
}

// Describe describes all the metrics exported by the namenode exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.jvmMetrics.Describe(ch)
	e.journalNode.Describe(ch)
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
		case strings.Contains(nameDataMap["name"].(string), hadoop.JournalNodeRPC):
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

		case strings.Contains(nameDataMap["name"].(string), hadoop.JournalNode):
			e.journalNode.Syncs60sNumOps.Set(nameDataMap["Syncs60sNumOps"].(float64))
			e.journalNode.Syncs60s50thPercentileLatencyMicros.Set(nameDataMap["Syncs60s50thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs60s75thPercentileLatencyMicros.Set(nameDataMap["Syncs60s75thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs60s90thPercentileLatencyMicros.Set(nameDataMap["Syncs60s90thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs60s95thPercentileLatencyMicros.Set(nameDataMap["Syncs60s95thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs60s99thPercentileLatencyMicros.Set(nameDataMap["Syncs60s99thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs300sNumOps.Set(nameDataMap["Syncs300sNumOps"].(float64))
			e.journalNode.Syncs300s50thPercentileLatencyMicros.Set(nameDataMap["Syncs300s50thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs300s75thPercentileLatencyMicros.Set(nameDataMap["Syncs300s75thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs300s90thPercentileLatencyMicros.Set(nameDataMap["Syncs300s90thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs300s95thPercentileLatencyMicros.Set(nameDataMap["Syncs300s95thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs300s99thPercentileLatencyMicros.Set(nameDataMap["Syncs300s99thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs3600sNumOps.Set(nameDataMap["Syncs3600sNumOps"].(float64))
			e.journalNode.Syncs3600s50thPercentileLatencyMicros.Set(nameDataMap["Syncs3600s50thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs3600s75thPercentileLatencyMicros.Set(nameDataMap["Syncs3600s75thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs3600s90thPercentileLatencyMicros.Set(nameDataMap["Syncs3600s90thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs3600s95thPercentileLatencyMicros.Set(nameDataMap["Syncs3600s95thPercentileLatencyMicros"].(float64))
			e.journalNode.Syncs3600s99thPercentileLatencyMicros.Set(nameDataMap["Syncs3600s99thPercentileLatencyMicros"].(float64))
			e.journalNode.BatchesWritten.Set(nameDataMap["BatchesWritten"].(float64))
			e.journalNode.TxnsWritten.Set(nameDataMap["TxnsWritten"].(float64))
			e.journalNode.BytesWritten.Set(nameDataMap["BytesWritten"].(float64))
			e.journalNode.BatchesWrittenWhileLagging.Set(nameDataMap["BatchesWrittenWhileLagging"].(float64))
			e.journalNode.LastWriterEpoch.Set(nameDataMap["LastWriterEpoch"].(float64))
			e.journalNode.CurrentLagTxns.Set(nameDataMap["CurrentLagTxns"].(float64))
			e.journalNode.LastWrittenTxId.Set(nameDataMap["LastWrittenTxId"].(float64))
			e.journalNode.LastPromisedEpoch.Set(nameDataMap["LastPromisedEpoch"].(float64))

		case nameDataMap["name"].(string) == hadoop.JournalNodeJvmMetrics:
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

		}

	}
	e.journalNode.Collect(ch)
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
