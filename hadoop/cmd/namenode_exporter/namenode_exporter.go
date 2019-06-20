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
	namespace = "namenode"
	jvm       = "jvm"
	dfs       = "dfs"
	rpc       = "rpc"
)

// Exporter  需要在 prometheus 中注册
type Exporter struct {
	jvmMetrics *model.JvmMetrics
	nameNode   *model.NameNode
	nameSystem *model.FsNameSystem
	rpc        *model.RPC
}

// NewExporter returns an initialized exporter.
func NewExporter() *Exporter {
	e := &Exporter{}
	e.nameNode = model.NewNameNodeExporter(namespace, dfs)
	e.jvmMetrics = model.NewJvmMetricsExporter(namespace, jvm)
	e.nameSystem = model.NewFsNameSystemExporter(namespace, dfs)
	e.rpc = model.NewRPCExporter(namespace, rpc)
	return e
}

// Describe describes all the metrics exported by the namenode exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.jvmMetrics.Describe(ch)
	e.nameNode.Describe(ch)
	e.nameSystem.Describe(ch)
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
		case nameDataMap["name"].(string) == hadoop.NameNodeActivity:
			e.nameNode.CreateFileOps.Set(nameDataMap["CreateFileOps"].(float64))
			e.nameNode.FilesCreated.Set(nameDataMap["FilesCreated"].(float64))
			e.nameNode.FilesAppended.Set(nameDataMap["FilesAppended"].(float64))
			e.nameNode.GetBlockLocations.Set(nameDataMap["GetBlockLocations"].(float64))
			e.nameNode.FilesRenamed.Set(nameDataMap["FilesRenamed"].(float64))
			e.nameNode.GetListingOps.Set(nameDataMap["GetListingOps"].(float64))
			e.nameNode.DeleteFileOps.Set(nameDataMap["DeleteFileOps"].(float64))
			e.nameNode.FilesDeleted.Set(nameDataMap["FilesDeleted"].(float64))
			e.nameNode.FileInfoOps.Set(nameDataMap["FileInfoOps"].(float64))
			e.nameNode.AddBlockOps.Set(nameDataMap["AddBlockOps"].(float64))
			e.nameNode.GetAdditionalDatanodeOps.Set(nameDataMap["GetAdditionalDatanodeOps"].(float64))
			e.nameNode.CreateSymlinkOps.Set(nameDataMap["CreateSymlinkOps"].(float64))
			e.nameNode.GetLinkTargetOps.Set(nameDataMap["GetLinkTargetOps"].(float64))
			e.nameNode.FilesInGetListingOps.Set(nameDataMap["FilesInGetListingOps"].(float64))
			e.nameNode.AllowSnapshotOps.Set(nameDataMap["AllowSnapshotOps"].(float64))
			e.nameNode.DisallowSnapshotOps.Set(nameDataMap["DisallowSnapshotOps"].(float64))
			e.nameNode.CreateSnapshotOps.Set(nameDataMap["CreateSnapshotOps"].(float64))
			e.nameNode.DeleteSnapshotOps.Set(nameDataMap["DeleteSnapshotOps"].(float64))
			e.nameNode.RenameSnapshotOps.Set(nameDataMap["RenameSnapshotOps"].(float64))
			e.nameNode.ListSnapshottableDirOps.Set(nameDataMap["ListSnapshottableDirOps"].(float64))
			e.nameNode.SnapshotDiffReportOps.Set(nameDataMap["SnapshotDiffReportOps"].(float64))
			e.nameNode.TransactionsNumOps.Set(nameDataMap["TransactionsNumOps"].(float64))
			e.nameNode.TransactionsAvgTime.Set(nameDataMap["TransactionsAvgTime"].(float64))
			e.nameNode.SyncsNumOps.Set(nameDataMap["SyncsNumOps"].(float64))
			e.nameNode.SyncsAvgTime.Set(nameDataMap["SyncsAvgTime"].(float64))
			e.nameNode.TransactionsBatchedInSync.Set(nameDataMap["TransactionsBatchedInSync"].(float64))
			e.nameNode.BlockReportNumOps.Set(nameDataMap["BlockReportNumOps"].(float64))
			e.nameNode.BlockReportAvgTime.Set(nameDataMap["BlockReportAvgTime"].(float64))
			e.nameNode.CacheReportNumOps.Set(nameDataMap["CacheReportNumOps"].(float64))
			e.nameNode.CacheReportAvgTime.Set(nameDataMap["CacheReportAvgTime"].(float64))
			e.nameNode.SafeModeTime.Set(nameDataMap["SafeModeTime"].(float64))
			e.nameNode.FsImageLoadTime.Set(nameDataMap["FsImageLoadTime"].(float64))
			e.nameNode.GetEditNumOps.Set(nameDataMap["GetEditNumOps"].(float64))
			e.nameNode.GetEditAvgTime.Set(nameDataMap["GetEditAvgTime"].(float64))
			e.nameNode.GetImageNumOps.Set(nameDataMap["GetImageNumOps"].(float64))
			e.nameNode.GetImageAvgTime.Set(nameDataMap["GetImageAvgTime"].(float64))
			e.nameNode.PutImageNumOps.Set(nameDataMap["PutImageNumOps"].(float64))
			e.nameNode.PutImageAvgTime.Set(nameDataMap["PutImageAvgTime"].(float64))
		case nameDataMap["name"].(string) == hadoop.NameNodeJvmMetrics:
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
		case nameDataMap["name"].(string) == hadoop.NameNodeFSNamesystem:
			e.nameSystem.MissingBlocks.Set(nameDataMap["MissingBlocks"].(float64))
			e.nameSystem.ExpiredHeartbeats.Set(nameDataMap["ExpiredHeartbeats"].(float64))
			e.nameSystem.TransactionsSinceLastCheckpoint.Set(nameDataMap["TransactionsSinceLastCheckpoint"].(float64))
			e.nameSystem.TransactionsSinceLastLogRoll.Set(nameDataMap["TransactionsSinceLastLogRoll"].(float64))
			e.nameSystem.LastWrittenTransactionId.Set(nameDataMap["LastWrittenTransactionId"].(float64))
			e.nameSystem.LastCheckpointTime.Set(nameDataMap["LastCheckpointTime"].(float64))
			e.nameSystem.CapacityTotal.Set(nameDataMap["CapacityTotal"].(float64))
			e.nameSystem.CapacityTotalGB.Set(nameDataMap["CapacityTotalGB"].(float64))
			e.nameSystem.CapacityUsed.Set(nameDataMap["CapacityUsed"].(float64))
			e.nameSystem.CapacityUsedGB.Set(nameDataMap["CapacityUsedGB"].(float64))
			e.nameSystem.CapacityRemaining.Set(nameDataMap["CapacityRemaining"].(float64))
			e.nameSystem.CapacityRemainingGB.Set(nameDataMap["CapacityRemainingGB"].(float64))
			e.nameSystem.CapacityUsedNonDFS.Set(nameDataMap["CapacityUsedNonDFS"].(float64))
			e.nameSystem.TotalLoad.Set(nameDataMap["TotalLoad"].(float64))
			e.nameSystem.SnapshottableDirectories.Set(nameDataMap["SnapshottableDirectories"].(float64))
			e.nameSystem.Snapshots.Set(nameDataMap["Snapshots"].(float64))
			e.nameSystem.BlocksTotal.Set(nameDataMap["BlocksTotal"].(float64))
			e.nameSystem.FilesTotal.Set(nameDataMap["FilesTotal"].(float64))
			e.nameSystem.PendingReplicationBlocks.Set(nameDataMap["PendingReplicationBlocks"].(float64))
			e.nameSystem.UnderReplicatedBlocks.Set(nameDataMap["UnderReplicatedBlocks"].(float64))
			e.nameSystem.CorruptBlocks.Set(nameDataMap["CorruptBlocks"].(float64))
			e.nameSystem.ScheduledReplicationBlocks.Set(nameDataMap["ScheduledReplicationBlocks"].(float64))
			e.nameSystem.PendingDeletionBlocks.Set(nameDataMap["PendingDeletionBlocks"].(float64))
			e.nameSystem.ExcessBlocks.Set(nameDataMap["ExcessBlocks"].(float64))
			e.nameSystem.PostponedMisreplicatedBlocks.Set(nameDataMap["PostponedMisreplicatedBlocks"].(float64))
			e.nameSystem.PendingDataNodeMessageCount.Set(nameDataMap["PendingDataNodeMessageCount"].(float64))
			e.nameSystem.MillisSinceLastLoadedEdits.Set(nameDataMap["MillisSinceLastLoadedEdits"].(float64))
			e.nameSystem.BlockCapacity.Set(nameDataMap["BlockCapacity"].(float64))
			e.nameSystem.StaleDataNodes.Set(nameDataMap["StaleDataNodes"].(float64))
			e.nameSystem.TotalFiles.Set(nameDataMap["TotalFiles"].(float64))
		case nameDataMap["name"].(string) == hadoop.NameNodeStatus:
			ch <- prometheus.MustNewConstMetric(e.nameNode.State, prometheus.GaugeValue, 0, nameDataMap["State"].(string))

		case strings.Contains(nameDataMap["name"].(string), hadoop.NameNodeRPC):
			// Each metrics record contains tags such as Hostname and port (number to which server is bound) as additional information along with metrics.
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
	e.nameNode.Collect(ch)
	e.jvmMetrics.Collect(ch)
	e.nameSystem.Collect(ch)
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
