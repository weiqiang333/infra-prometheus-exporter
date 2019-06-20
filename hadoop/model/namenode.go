package model

import "github.com/prometheus/client_golang/prometheus"

// NameNode JMX中收集的NameNode 信息
type NameNode struct {
	// NameNode status
	State                     *prometheus.Desc
	CreateFileOps             prometheus.Gauge
	FilesCreated              prometheus.Gauge
	FilesAppended             prometheus.Gauge
	GetBlockLocations         prometheus.Gauge
	FilesRenamed              prometheus.Gauge
	GetListingOps             prometheus.Gauge
	DeleteFileOps             prometheus.Gauge
	FilesDeleted              prometheus.Gauge
	FileInfoOps               prometheus.Gauge
	AddBlockOps               prometheus.Gauge
	GetAdditionalDatanodeOps  prometheus.Gauge
	CreateSymlinkOps          prometheus.Gauge
	GetLinkTargetOps          prometheus.Gauge
	FilesInGetListingOps      prometheus.Gauge
	AllowSnapshotOps          prometheus.Gauge
	DisallowSnapshotOps       prometheus.Gauge
	CreateSnapshotOps         prometheus.Gauge
	DeleteSnapshotOps         prometheus.Gauge
	RenameSnapshotOps         prometheus.Gauge
	ListSnapshottableDirOps   prometheus.Gauge
	SnapshotDiffReportOps     prometheus.Gauge
	TransactionsNumOps        prometheus.Gauge
	TransactionsAvgTime       prometheus.Gauge
	SyncsNumOps               prometheus.Gauge
	SyncsAvgTime              prometheus.Gauge
	TransactionsBatchedInSync prometheus.Gauge
	BlockReportNumOps         prometheus.Gauge
	BlockReportAvgTime        prometheus.Gauge
	CacheReportNumOps         prometheus.Gauge
	CacheReportAvgTime        prometheus.Gauge
	SafeModeTime              prometheus.Gauge
	FsImageLoadTime           prometheus.Gauge
	GetEditNumOps             prometheus.Gauge
	GetEditAvgTime            prometheus.Gauge
	GetImageNumOps            prometheus.Gauge
	GetImageAvgTime           prometheus.Gauge
	PutImageNumOps            prometheus.Gauge
	PutImageAvgTime           prometheus.Gauge
}

// NewNameNodeExporter 返回 NameNode 部分的exporter
func NewNameNodeExporter(namespace, context string) *NameNode {
	return &NameNode{
		State: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "state"),
			"Indicate namenode state .",
			[]string{"status"},
			nil,
		),
		CreateFileOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "create_file_ops",
			Help:      "Total number of files creat",
		}),
		FilesCreated: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "files_created",
			Help:      "Total number of files and directories created by create or mkdir operations",
		}),
		FilesAppended: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "files_appended",
			Help:      "Total number of files appended",
		}),
		GetBlockLocations: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "get_block_locations",
			Help:      "Total number of getBlockLocations operations",
		}),
		FilesRenamed: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "files_renamed",
			Help:      "Total number of rename operations (NOT number of files/dirs renamed)",
		}),
		GetListingOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "get_listing_ops",
			Help:      "Total number of directory listing operations",
		}),
		DeleteFileOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "delete_file_ops",
			Help:      "Total number of delete operations",
		}),
		FilesDeleted: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "files_deleted",
			Help:      "Total number of files and directories deleted by delete or rename operations",
		}),
		FileInfoOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "file_info_ops",
			Help:      "Total number of getFileInfo and getLinkFileInfo operations",
		}),
		AddBlockOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "add_block_ops",
			Help:      "Total number of addBlock operations succeeded",
		}),
		GetAdditionalDatanodeOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "get_additional_datanode_ops",
			Help:      "Total number of getAdditionalDatanode operations",
		}),
		CreateSymlinkOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "create_symlink_ops",
			Help:      "Total number of createSymlink operations",
		}),
		GetLinkTargetOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "getLink_target_ops",
			Help:      "Total number of getLinkTarget operations",
		}),
		FilesInGetListingOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "files_in_get_listing_ops",
			Help:      "Total number of files and directories listed by directory listing operations",
		}),
		AllowSnapshotOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "allow_snapshot_ops",
			Help:      "Total number of allowSnapshot operations",
		}),
		DisallowSnapshotOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "disallow_snapshot_ops",
			Help:      "Total number of disallowSnapshot operations",
		}),
		CreateSnapshotOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "create_snapshot_ops",
			Help:      "Total number of createSnapshot operations",
		}),
		DeleteSnapshotOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "delete_snapshot_ops",
			Help:      "Total number of deleteSnapshot operations",
		}),
		RenameSnapshotOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "rename_snapshot_ops",
			Help:      "Total number of renameSnapshot operations",
		}),
		ListSnapshottableDirOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "list_snapshottable_dir_ops",
			Help:      "Total number of snapshottableDirectoryStatus operations",
		}),
		SnapshotDiffReportOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "snapshot_diff_report_ops",
			Help:      "Total number of getSnapshotDiffReport operations",
		}),
		TransactionsNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "transactions_num_ops",
			Help:      "Total number of Journal transactions",
		}),
		TransactionsAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "transactions_avg_time",
			Help:      "Average time of Journal transactions in milliseconds",
		}),
		SyncsNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_num_ops",
			Help:      "Total number of Journal syncs",
		}),
		SyncsAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_avg_time",
			Help:      "Average time of Journal syncs in milliseconds",
		}),
		TransactionsBatchedInSync: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "transactions_batched_in_sync",
			Help:      "Total number of Journal transactions batched in sync",
		}),
		BlockReportNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_report_num_ops",
			Help:      "Total number of processing block reports from DataNode",
		}),
		BlockReportAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_report_avg_time",
			Help:      "Average time of processing block reports in milliseconds",
		}),
		CacheReportNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "cache_report_num_ops",
			Help:      "Total number of processing cache reports from DataNode",
		}),
		CacheReportAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "cache_report_avg_time",
			Help:      "Average time of processing cache reports in milliseconds",
		}),
		SafeModeTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "safe_mode_time",
			Help:      "The interval between FSNameSystem starts and the last time safemode leaves in milliseconds.  (sometimes not equal to the time in SafeMode, see HDFS-5156)",
		}),
		FsImageLoadTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "fs_image_load_time",
			Help:      "Time loading FS Image at startup in milliseconds",
		}),
		GetEditNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "get_edit_num_ops",
			Help:      "Total number of edits downloads from SecondaryNameNode",
		}),
		GetEditAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "get_edit_avg_time",
			Help:      "Average edits download time in milliseconds",
		}),
		GetImageNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "get_image_num_ops",
			Help:      "Total number of fsimage downloads from SecondaryNameNode",
		}),
		GetImageAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "get_image_avg_time",
			Help:      "Average fsimage download time in milliseconds",
		}),
		PutImageNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "put_image_num_ops",
			Help:      "Total number of fsimage uploads to SecondaryNameNode",
		}),
		PutImageAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "put_image_avg_time",
			Help:      "Average fsimage upload time in milliseconds",
		}),
	}
}

// Describe 通用方法
func (n *NameNode) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.State
	// namenode dfs context
	n.CreateFileOps.Describe(ch)
	n.FilesCreated.Describe(ch)
	n.FilesAppended.Describe(ch)
	n.GetBlockLocations.Describe(ch)
	n.FilesRenamed.Describe(ch)
	n.GetListingOps.Describe(ch)
	n.DeleteFileOps.Describe(ch)
	n.FilesDeleted.Describe(ch)
	n.FileInfoOps.Describe(ch)
	n.AddBlockOps.Describe(ch)
	n.GetAdditionalDatanodeOps.Describe(ch)
	n.CreateSymlinkOps.Describe(ch)
	n.GetLinkTargetOps.Describe(ch)
	n.FilesInGetListingOps.Describe(ch)
	n.AllowSnapshotOps.Describe(ch)
	n.DisallowSnapshotOps.Describe(ch)
	n.CreateSnapshotOps.Describe(ch)
	n.DeleteSnapshotOps.Describe(ch)
	n.RenameSnapshotOps.Describe(ch)
	n.ListSnapshottableDirOps.Describe(ch)
	n.SnapshotDiffReportOps.Describe(ch)
	n.TransactionsNumOps.Describe(ch)
	n.TransactionsAvgTime.Describe(ch)
	n.SyncsNumOps.Describe(ch)
	n.SyncsAvgTime.Describe(ch)
	n.TransactionsBatchedInSync.Describe(ch)
	n.BlockReportNumOps.Describe(ch)
	n.BlockReportAvgTime.Describe(ch)
	n.CacheReportNumOps.Describe(ch)
	n.CacheReportAvgTime.Describe(ch)
	n.SafeModeTime.Describe(ch)
	n.FsImageLoadTime.Describe(ch)
	n.GetEditNumOps.Describe(ch)
	n.GetEditAvgTime.Describe(ch)
	n.GetImageNumOps.Describe(ch)
	n.GetImageAvgTime.Describe(ch)
	n.PutImageNumOps.Describe(ch)
	n.PutImageAvgTime.Describe(ch)
}

// Collect 通用方法
func (n *NameNode) Collect(ch chan<- prometheus.Metric) {
	// namenode dfs context
	n.CreateFileOps.Collect(ch)
	n.FilesCreated.Collect(ch)
	n.FilesAppended.Collect(ch)
	n.GetBlockLocations.Collect(ch)
	n.FilesRenamed.Collect(ch)
	n.GetListingOps.Collect(ch)
	n.DeleteFileOps.Collect(ch)
	n.FilesDeleted.Collect(ch)
	n.FileInfoOps.Collect(ch)
	n.AddBlockOps.Collect(ch)
	n.GetAdditionalDatanodeOps.Collect(ch)
	n.CreateSymlinkOps.Collect(ch)
	n.GetLinkTargetOps.Collect(ch)
	n.FilesInGetListingOps.Collect(ch)
	n.AllowSnapshotOps.Collect(ch)
	n.DisallowSnapshotOps.Collect(ch)
	n.CreateSnapshotOps.Collect(ch)
	n.DeleteSnapshotOps.Collect(ch)
	n.RenameSnapshotOps.Collect(ch)
	n.ListSnapshottableDirOps.Collect(ch)
	n.SnapshotDiffReportOps.Collect(ch)
	n.TransactionsNumOps.Collect(ch)
	n.TransactionsAvgTime.Collect(ch)
	n.SyncsNumOps.Collect(ch)
	n.SyncsAvgTime.Collect(ch)
	n.TransactionsBatchedInSync.Collect(ch)
	n.BlockReportNumOps.Collect(ch)
	n.BlockReportAvgTime.Collect(ch)
	n.CacheReportNumOps.Collect(ch)
	n.CacheReportAvgTime.Collect(ch)
	n.SafeModeTime.Collect(ch)
	n.FsImageLoadTime.Collect(ch)
	n.GetEditNumOps.Collect(ch)
	n.GetEditAvgTime.Collect(ch)
	n.GetImageNumOps.Collect(ch)
	n.GetImageAvgTime.Collect(ch)
	n.PutImageNumOps.Collect(ch)
	n.PutImageAvgTime.Collect(ch)
}
