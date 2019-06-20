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
	namespace = "datanode"
	jvm       = "jvm"
	dfs       = "dfs"
	rpc       = "rpc"
)

// Exporter  需要在 prometheus 中注册
type Exporter struct {
	jvmMetrics *model.JvmMetrics
	dataNode   *model.DataNode
	rpc        *model.RPC
}

// NewExporter returns an initialized exporter.
func NewExporter() *Exporter {
	e := &Exporter{}
	e.dataNode = model.NewDataNodeExporter(namespace, dfs)
	e.jvmMetrics = model.NewJvmMetricsExporter(namespace, jvm)
	e.rpc = model.NewRPCExporter(namespace, rpc)
	return e
}

// Describe describes all the metrics exported by the namenode exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.jvmMetrics.Describe(ch)
	e.dataNode.Describe(ch)
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
		case nameDataMap["name"].(string) == hadoop.DataNodeActivity:
			e.dataNode.BytesWritten.Set(nameDataMap["BytesWritten"].(float64))
			e.dataNode.BytesRead.Set(nameDataMap["BytesRead"].(float64))
			e.dataNode.BlocksWritten.Set(nameDataMap["BlocksWritten"].(float64))
			e.dataNode.BlocksRead.Set(nameDataMap["BlocksRead"].(float64))
			e.dataNode.BlocksReplicated.Set(nameDataMap["BlocksReplicated"].(float64))
			e.dataNode.BlocksRemoved.Set(nameDataMap["BlocksRemoved"].(float64))
			e.dataNode.BlocksVerified.Set(nameDataMap["BlocksVerified"].(float64))
			e.dataNode.BlockVerificationFailures.Set(nameDataMap["BlockVerificationFailures"].(float64))
			e.dataNode.BlocksCached.Set(nameDataMap["BlocksCached"].(float64))
			e.dataNode.BlocksUncached.Set(nameDataMap["BlocksUncached"].(float64))
			e.dataNode.ReadsFromLocalClient.Set(nameDataMap["ReadsFromLocalClient"].(float64))
			e.dataNode.ReadsFromRemoteClient.Set(nameDataMap["ReadsFromRemoteClient"].(float64))
			e.dataNode.WritesFromLocalClient.Set(nameDataMap["WritesFromLocalClient"].(float64))
			e.dataNode.WritesFromRemoteClient.Set(nameDataMap["WritesFromRemoteClient"].(float64))
			e.dataNode.BlocksGetLocalPathInfo.Set(nameDataMap["BlocksGetLocalPathInfo"].(float64))
			e.dataNode.FsyncCount.Set(nameDataMap["FsyncCount"].(float64))
			e.dataNode.VolumeFailures.Set(nameDataMap["VolumeFailures"].(float64))
			e.dataNode.ReadBlockOpNumOps.Set(nameDataMap["ReadBlockOpNumOps"].(float64))
			e.dataNode.ReadBlockOpAvgTime.Set(nameDataMap["ReadBlockOpAvgTime"].(float64))
			e.dataNode.WriteBlockOpNumOps.Set(nameDataMap["WriteBlockOpNumOps"].(float64))
			e.dataNode.WriteBlockOpAvgTime.Set(nameDataMap["WriteBlockOpAvgTime"].(float64))
			e.dataNode.BlockChecksumOpNumOps.Set(nameDataMap["BlockChecksumOpNumOps"].(float64))
			e.dataNode.BlockChecksumOpAvgTime.Set(nameDataMap["BlockChecksumOpAvgTime"].(float64))
			e.dataNode.CopyBlockOpNumOps.Set(nameDataMap["CopyBlockOpNumOps"].(float64))
			e.dataNode.CopyBlockOpAvgTime.Set(nameDataMap["CopyBlockOpAvgTime"].(float64))
			e.dataNode.ReplaceBlockOpNumOps.Set(nameDataMap["ReplaceBlockOpNumOps"].(float64))
			e.dataNode.ReplaceBlockOpAvgTime.Set(nameDataMap["ReplaceBlockOpAvgTime"].(float64))
			e.dataNode.HeartbeatsNumOps.Set(nameDataMap["HeartbeatsNumOps"].(float64))
			e.dataNode.HeartbeatsAvgTime.Set(nameDataMap["HeartbeatsAvgTime"].(float64))
			e.dataNode.BlockReportsNumOps.Set(nameDataMap["BlockReportsNumOps"].(float64))
			e.dataNode.BlockReportsAvgTime.Set(nameDataMap["BlockReportsAvgTime"].(float64))
			e.dataNode.CacheReportsNumOps.Set(nameDataMap["CacheReportsNumOps"].(float64))
			e.dataNode.CacheReportsAvgTime.Set(nameDataMap["CacheReportsAvgTime"].(float64))
			e.dataNode.PacketAckRoundTripTimeNanosNumOps.Set(nameDataMap["PacketAckRoundTripTimeNanosNumOps"].(float64))
			e.dataNode.PacketAckRoundTripTimeNanosAvgTime.Set(nameDataMap["PacketAckRoundTripTimeNanosAvgTime"].(float64))
			e.dataNode.FlushNanosNumOps.Set(nameDataMap["FlushNanosNumOps"].(float64))
			e.dataNode.FlushNanosAvgTime.Set(nameDataMap["FlushNanosAvgTime"].(float64))
			e.dataNode.FsyncNanosNumOps.Set(nameDataMap["FsyncNanosNumOps"].(float64))
			e.dataNode.FsyncNanosAvgTime.Set(nameDataMap["FsyncNanosAvgTime"].(float64))
			e.dataNode.SendDataPacketBlockedOnNetworkNanosNumOps.Set(nameDataMap["SendDataPacketBlockedOnNetworkNanosNumOps"].(float64))
			e.dataNode.SendDataPacketBlockedOnNetworkNanosAvgTime.Set(nameDataMap["SendDataPacketBlockedOnNetworkNanosAvgTime"].(float64))
			e.dataNode.SendDataPacketTransferNanosNumOps.Set(nameDataMap["SendDataPacketTransferNanosNumOps"].(float64))
			e.dataNode.SendDataPacketTransferNanosAvgTime.Set(nameDataMap["SendDataPacketTransferNanosAvgTime"].(float64))

		case nameDataMap["name"].(string) == hadoop.DataNodeJvmMetrics:
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
			e.jvmMetrics.GcNumWarnThresholdExceeded.Set(nameDataMap["GcNumWarnThresholdExceeded"].(float64))
			e.jvmMetrics.GcNumInfoThresholdExceeded.Set(nameDataMap["GcNumInfoThresholdExceeded"].(float64))
			e.jvmMetrics.GcTotalExtraSleepTime.Set(nameDataMap["GcTotalExtraSleepTime"].(float64))
		case strings.Contains(nameDataMap["name"].(string), hadoop.DataNodeRPC):
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
	e.dataNode.Collect(ch)
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
