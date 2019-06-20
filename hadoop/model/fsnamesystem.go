package model

import "github.com/prometheus/client_golang/prometheus"

// FsNameSystem metrics
type FsNameSystem struct {
	MissingBlocks                   prometheus.Gauge
	ExpiredHeartbeats               prometheus.Gauge
	TransactionsSinceLastCheckpoint prometheus.Gauge
	TransactionsSinceLastLogRoll    prometheus.Gauge
	// 这里golint检查不通过，但是jmx导出的数据就是 Id,所以这里不修改
	LastWrittenTransactionId     prometheus.Gauge
	LastCheckpointTime           prometheus.Gauge
	CapacityTotal                prometheus.Gauge
	CapacityTotalGB              prometheus.Gauge
	CapacityUsed                 prometheus.Gauge
	CapacityUsedGB               prometheus.Gauge
	CapacityRemaining            prometheus.Gauge
	CapacityRemainingGB          prometheus.Gauge
	CapacityUsedNonDFS           prometheus.Gauge
	TotalLoad                    prometheus.Gauge
	SnapshottableDirectories     prometheus.Gauge
	Snapshots                    prometheus.Gauge
	BlocksTotal                  prometheus.Gauge
	FilesTotal                   prometheus.Gauge
	PendingReplicationBlocks     prometheus.Gauge
	UnderReplicatedBlocks        prometheus.Gauge
	CorruptBlocks                prometheus.Gauge
	ScheduledReplicationBlocks   prometheus.Gauge
	PendingDeletionBlocks        prometheus.Gauge
	ExcessBlocks                 prometheus.Gauge
	PostponedMisreplicatedBlocks prometheus.Gauge
	PendingDataNodeMessageCount  prometheus.Gauge
	MillisSinceLastLoadedEdits   prometheus.Gauge
	BlockCapacity                prometheus.Gauge
	StaleDataNodes               prometheus.Gauge
	TotalFiles                   prometheus.Gauge
}

// NewFsNameSystemExporter 返回 FsNameSystem 部分的exporter
func NewFsNameSystemExporter(namespace, context string) *FsNameSystem {
	return &FsNameSystem{
		MissingBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "missing_blocks",
			Help:      "Current number of missing blocks",
		}),
		ExpiredHeartbeats: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "expired_heartbeats",
			Help:      "Total number of expired heartbeats",
		}),
		TransactionsSinceLastCheckpoint: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "transactions_since_last_checkpoint",
			Help:      "Total number of transactions since last checkpoint",
		}),
		TransactionsSinceLastLogRoll: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "transactions_since_last_log_roll",
			Help:      "Total number of transactions since last edit log roll",
		}),
		LastWrittenTransactionId: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "last_written_transaction_id",
			Help:      "Last transaction ID written to the edit log",
		}),
		LastCheckpointTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "last_checkpoint_time",
			Help:      "Time in milliseconds since epoch of last checkpoint",
		}),
		CapacityTotal: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "capacity_total",
			Help:      "Current raw capacity of DataNodes in bytes",
		}),
		CapacityTotalGB: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "capacity_total_gb",
			Help:      "Current raw capacity of DataNodes in GB",
		}),
		CapacityUsed: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "capacity_used",
			Help:      "Current used capacity across all DataNodes in bytes",
		}),
		CapacityUsedGB: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "capacity_used_gb",
			Help:      "Current used capacity across all DataNodes in GB",
		}),
		CapacityRemaining: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "capacity_remaining",
			Help:      "Current remaining capacity in bytes",
		}),
		CapacityRemainingGB: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "capacity_remaining_gb",
			Help:      "Current remaining capacity in GB",
		}),
		CapacityUsedNonDFS: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "capacity_used_non_dfs",
			Help:      "Current space used by DataNodes for non DFS purposes in bytes",
		}),
		TotalLoad: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "total_load",
			Help:      "Current number of connections",
		}),
		SnapshottableDirectories: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "snapshottable_directories",
			Help:      "Current number of snapshottable directories",
		}),
		Snapshots: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "snapshots",
			Help:      "Current number of snapshots",
		}),
		BlocksTotal: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_total",
			Help:      "Current number of allocated blocks in the system",
		}),
		FilesTotal: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "files_total",
			Help:      "Current number of files and directories",
		}),
		PendingReplicationBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "pending_replication_blocks",
			Help:      "Current number of blocks pending to be replicated",
		}),
		UnderReplicatedBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "under_replicated_blocks",
			Help:      "Current number of blocks under replicated",
		}),
		CorruptBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "corrupt_blocks",
			Help:      "Current number of blocks with corrupt replicas",
		}),
		ScheduledReplicationBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "scheduled_replication_blocks",
			Help:      "Current number of blocks scheduled for replications",
		}),
		PendingDeletionBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "pending_deletion_blocks",
			Help:      "Current number of blocks pending deletion",
		}),
		ExcessBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "excess_blocks",
			Help:      "Current number of excess blocks",
		}),
		PostponedMisreplicatedBlocks: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "postponed_misreplicated_blocks",
			Help:      "(HA-only) Current number of blocks postponed to replicate",
		}),
		PendingDataNodeMessageCount: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "pending_data_node_message_count",
			Help:      "(HA-only) Current number of pending block-related messages for later processing in the standby NameNode",
		}),
		MillisSinceLastLoadedEdits: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "millis_since_last_loaded_edits",
			Help:      "(HA-only) Time in milliseconds since the last time standby NameNode load edit log. In active NameNode, set to 0",
		}),
		BlockCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_capacity",
			Help:      "Current number of block capacity",
		}),
		StaleDataNodes: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "stale_data_nodes",
			Help:      "Current number of DataNodes marked stale due to delayed heartbeat",
		}),
		TotalFiles: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "total_files",
			Help:      "Deprecated: Use FilesTotal instead",
		}),
	}
}

// Describe 通用方法
func (f *FsNameSystem) Describe(ch chan<- *prometheus.Desc) {
	f.MissingBlocks.Describe(ch)
	f.ExpiredHeartbeats.Describe(ch)
	f.TransactionsSinceLastCheckpoint.Describe(ch)
	f.TransactionsSinceLastLogRoll.Describe(ch)
	f.LastWrittenTransactionId.Describe(ch)
	f.LastCheckpointTime.Describe(ch)
	f.CapacityTotal.Describe(ch)
	f.CapacityTotalGB.Describe(ch)
	f.CapacityUsed.Describe(ch)
	f.CapacityUsedGB.Describe(ch)
	f.CapacityRemaining.Describe(ch)
	f.CapacityRemainingGB.Describe(ch)
	f.CapacityUsedNonDFS.Describe(ch)
	f.TotalLoad.Describe(ch)
	f.SnapshottableDirectories.Describe(ch)
	f.Snapshots.Describe(ch)
	f.BlocksTotal.Describe(ch)
	f.FilesTotal.Describe(ch)
	f.PendingReplicationBlocks.Describe(ch)
	f.UnderReplicatedBlocks.Describe(ch)
	f.CorruptBlocks.Describe(ch)
	f.ScheduledReplicationBlocks.Describe(ch)
	f.PendingDeletionBlocks.Describe(ch)
	f.ExcessBlocks.Describe(ch)
	f.PostponedMisreplicatedBlocks.Describe(ch)
	f.PendingDataNodeMessageCount.Describe(ch)
	f.MillisSinceLastLoadedEdits.Describe(ch)
	f.BlockCapacity.Describe(ch)
	f.StaleDataNodes.Describe(ch)
	f.TotalFiles.Describe(ch)

}

// Collect 通用方法
func (f *FsNameSystem) Collect(ch chan<- prometheus.Metric) {
	f.MissingBlocks.Collect(ch)
	f.ExpiredHeartbeats.Collect(ch)
	f.TransactionsSinceLastCheckpoint.Collect(ch)
	f.TransactionsSinceLastLogRoll.Collect(ch)
	f.LastWrittenTransactionId.Collect(ch)
	f.LastCheckpointTime.Collect(ch)
	f.CapacityTotal.Collect(ch)
	f.CapacityTotalGB.Collect(ch)
	f.CapacityUsed.Collect(ch)
	f.CapacityUsedGB.Collect(ch)
	f.CapacityRemaining.Collect(ch)
	f.CapacityRemainingGB.Collect(ch)
	f.CapacityUsedNonDFS.Collect(ch)
	f.TotalLoad.Collect(ch)
	f.SnapshottableDirectories.Collect(ch)
	f.Snapshots.Collect(ch)
	f.BlocksTotal.Collect(ch)
	f.FilesTotal.Collect(ch)
	f.PendingReplicationBlocks.Collect(ch)
	f.UnderReplicatedBlocks.Collect(ch)
	f.CorruptBlocks.Collect(ch)
	f.ScheduledReplicationBlocks.Collect(ch)
	f.PendingDeletionBlocks.Collect(ch)
	f.ExcessBlocks.Collect(ch)
	f.PostponedMisreplicatedBlocks.Collect(ch)
	f.PendingDataNodeMessageCount.Collect(ch)
	f.MillisSinceLastLoadedEdits.Collect(ch)
	f.BlockCapacity.Collect(ch)
	f.StaleDataNodes.Collect(ch)
	f.TotalFiles.Collect(ch)

}
