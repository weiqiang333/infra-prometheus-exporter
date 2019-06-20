package model

import "github.com/prometheus/client_golang/prometheus"

// DataNode metrics
type DataNode struct {
	BytesWritten                               prometheus.Gauge
	BytesRead                                  prometheus.Gauge
	BlocksWritten                              prometheus.Gauge
	BlocksRead                                 prometheus.Gauge
	BlocksReplicated                           prometheus.Gauge
	BlocksRemoved                              prometheus.Gauge
	BlocksVerified                             prometheus.Gauge
	BlockVerificationFailures                  prometheus.Gauge
	BlocksCached                               prometheus.Gauge
	BlocksUncached                             prometheus.Gauge
	ReadsFromLocalClient                       prometheus.Gauge
	ReadsFromRemoteClient                      prometheus.Gauge
	WritesFromLocalClient                      prometheus.Gauge
	WritesFromRemoteClient                     prometheus.Gauge
	BlocksGetLocalPathInfo                     prometheus.Gauge
	FsyncCount                                 prometheus.Gauge
	VolumeFailures                             prometheus.Gauge
	ReadBlockOpNumOps                          prometheus.Gauge
	ReadBlockOpAvgTime                         prometheus.Gauge
	WriteBlockOpNumOps                         prometheus.Gauge
	WriteBlockOpAvgTime                        prometheus.Gauge
	BlockChecksumOpNumOps                      prometheus.Gauge
	BlockChecksumOpAvgTime                     prometheus.Gauge
	CopyBlockOpNumOps                          prometheus.Gauge
	CopyBlockOpAvgTime                         prometheus.Gauge
	ReplaceBlockOpNumOps                       prometheus.Gauge
	ReplaceBlockOpAvgTime                      prometheus.Gauge
	HeartbeatsNumOps                           prometheus.Gauge
	HeartbeatsAvgTime                          prometheus.Gauge
	BlockReportsNumOps                         prometheus.Gauge
	BlockReportsAvgTime                        prometheus.Gauge
	CacheReportsNumOps                         prometheus.Gauge
	CacheReportsAvgTime                        prometheus.Gauge
	PacketAckRoundTripTimeNanosNumOps          prometheus.Gauge
	PacketAckRoundTripTimeNanosAvgTime         prometheus.Gauge
	FlushNanosNumOps                           prometheus.Gauge
	FlushNanosAvgTime                          prometheus.Gauge
	FsyncNanosNumOps                           prometheus.Gauge
	FsyncNanosAvgTime                          prometheus.Gauge
	SendDataPacketBlockedOnNetworkNanosNumOps  prometheus.Gauge
	SendDataPacketBlockedOnNetworkNanosAvgTime prometheus.Gauge
	SendDataPacketTransferNanosNumOps          prometheus.Gauge
	SendDataPacketTransferNanosAvgTime         prometheus.Gauge
}

// NewDataNodeExporter 返回 DataNode 部分的exporter
func NewDataNodeExporter(namespace, context string) *DataNode {
	return &DataNode{
		BytesWritten: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "bytes_written",
			Help:      "Total number of bytes written to DataNode",
		}),
		BytesRead: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "bytes_read",
			Help:      "Total number of bytes read from DataNode",
		}),
		BlocksWritten: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_written",
			Help:      "Total number of blocks written to DataNode",
		}),
		BlocksRead: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_read",
			Help:      "Total number of blocks read from DataNode",
		}),
		BlocksReplicated: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_replicated",
			Help:      "Total number of blocks replicated",
		}),
		BlocksRemoved: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_removed",
			Help:      "Total number of blocks removed",
		}),
		BlocksVerified: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_verified",
			Help:      "Total number of blocks verified",
		}),
		BlockVerificationFailures: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_verification_failures",
			Help:      "Total number of verifications failures",
		}),
		BlocksCached: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_cached",
			Help:      "Total number of blocks cached",
		}),
		BlocksUncached: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_uncached",
			Help:      "Total number of blocks cached",
		}),
		ReadsFromLocalClient: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "reads_from_local_client",
			Help:      "Total number of read operations from local client",
		}),
		ReadsFromRemoteClient: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "reads_from_remote_client",
			Help:      "Total number of read operations from remote client",
		}),
		WritesFromLocalClient: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "writes_from_local_client",
			Help:      "Total number of write operations from local client",
		}),
		WritesFromRemoteClient: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "writes_from_remote_client",
			Help:      "Total number of write operations from remote client",
		}),
		BlocksGetLocalPathInfo: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "blocks_get_local_path_info",
			Help:      "Total number of operations to get local path names of blocks",
		}),
		FsyncCount: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "fsync_count",
			Help:      "Total number of fsync",
		}),
		VolumeFailures: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "volume_failures",
			Help:      "Total number of volume failures occurred",
		}),
		ReadBlockOpNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "read_block_op_num_ops",
			Help:      "Total number of read operations",
		}),
		ReadBlockOpAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "read_block_op_avg_time",
			Help:      "Average time of read operations in milliseconds",
		}),
		WriteBlockOpNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "write_block_op_num_ops",
			Help:      "Total number of write operations",
		}),
		WriteBlockOpAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "write_block_op_avg_time",
			Help:      "Average time of write operations in milliseconds",
		}),
		BlockChecksumOpNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_checksum_op_num_ops",
			Help:      "Total number of blockChecksum operations",
		}),
		BlockChecksumOpAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_checksum_op_avg_time",
			Help:      "Average time of blockChecksum operations in milliseconds",
		}),
		CopyBlockOpNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "copy_block_op_num_ops",
			Help:      "Total number of block copy operations",
		}),
		CopyBlockOpAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "copy_block_op_avg_time",
			Help:      "Average time of block copy operations in milliseconds",
		}),
		ReplaceBlockOpNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "replace_block_op_num_ops",
			Help:      "Total number of block replace operations",
		}),
		ReplaceBlockOpAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "replace_block_op_avg_time",
			Help:      "Average time of block replace operations in milliseconds",
		}),
		HeartbeatsNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "heartbeats_num_ops",
			Help:      "Total number of heartbeats",
		}),
		HeartbeatsAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "heartbeats_avg_time",
			Help:      "Average heartbeat time in milliseconds",
		}),
		BlockReportsNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_reports_num_ops",
			Help:      "Total number of block report operations",
		}),
		BlockReportsAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "block_reports_avg_time",
			Help:      "Average time of block report operations in milliseconds",
		}),
		CacheReportsNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "cache_reports_mum_ops",
			Help:      "Total number of cache report operations",
		}),
		CacheReportsAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "cache_reports_avg_time",
			Help:      "Average time of cache report operations in milliseconds",
		}),
		PacketAckRoundTripTimeNanosNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "packet_ack_round_trip_time_manos_num_ops",
			Help:      "Total number of ack round trip",
		}),
		PacketAckRoundTripTimeNanosAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "packet_ack_round_trip_time_nanos_avg_time",
			Help:      "Average time from ack send to receive minus the downstream ack time in nanoseconds",
		}),
		FlushNanosNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "flush_nanos_num_ops",
			Help:      "Total number of flushes",
		}),
		FlushNanosAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "flush_nanos_avg_time",
			Help:      "Average flush time in nanoseconds",
		}),
		FsyncNanosNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "fsync_nanos_num_ops",
			Help:      "Total number of fsync",
		}),
		FsyncNanosAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "fsync_nanos_avg_time",
			Help:      "Average fsync time in nanoseconds",
		}),
		SendDataPacketBlockedOnNetworkNanosNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "send_data_packet_blocked_on_network_nanos_num_ops",
			Help:      "Total number of sending packets",
		}),
		SendDataPacketBlockedOnNetworkNanosAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "send_data_packet_blocked_on_network_nanos_avg_time",
			Help:      "Average waiting time of sending packets in nanoseconds",
		}),
		SendDataPacketTransferNanosNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "send_data_packet_transfer_nanos_num_ops",
			Help:      "Total number of sending packets",
		}),
		SendDataPacketTransferNanosAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "send_data_packet_transfer_nanos_avg_time",
			Help:      "Average transfer time of sending packets in nanoseconds",
		}),
	}
}

// Describe 通用方法
func (d *DataNode) Describe(ch chan<- *prometheus.Desc) {
	d.BytesWritten.Describe(ch)
	d.BytesRead.Describe(ch)
	d.BlocksWritten.Describe(ch)
	d.BlocksRead.Describe(ch)
	d.BlocksReplicated.Describe(ch)
	d.BlocksRemoved.Describe(ch)
	d.BlocksVerified.Describe(ch)
	d.BlockVerificationFailures.Describe(ch)
	d.BlocksCached.Describe(ch)
	d.BlocksUncached.Describe(ch)
	d.ReadsFromLocalClient.Describe(ch)
	d.ReadsFromRemoteClient.Describe(ch)
	d.WritesFromLocalClient.Describe(ch)
	d.WritesFromRemoteClient.Describe(ch)
	d.BlocksGetLocalPathInfo.Describe(ch)
	d.FsyncCount.Describe(ch)
	d.VolumeFailures.Describe(ch)
	d.ReadBlockOpNumOps.Describe(ch)
	d.ReadBlockOpAvgTime.Describe(ch)
	d.WriteBlockOpNumOps.Describe(ch)
	d.WriteBlockOpAvgTime.Describe(ch)
	d.BlockChecksumOpNumOps.Describe(ch)
	d.BlockChecksumOpAvgTime.Describe(ch)
	d.CopyBlockOpNumOps.Describe(ch)
	d.CopyBlockOpAvgTime.Describe(ch)
	d.ReplaceBlockOpNumOps.Describe(ch)
	d.ReplaceBlockOpAvgTime.Describe(ch)
	d.HeartbeatsNumOps.Describe(ch)
	d.HeartbeatsAvgTime.Describe(ch)
	d.BlockReportsNumOps.Describe(ch)
	d.BlockReportsAvgTime.Describe(ch)
	d.CacheReportsNumOps.Describe(ch)
	d.CacheReportsAvgTime.Describe(ch)
	d.PacketAckRoundTripTimeNanosNumOps.Describe(ch)
	d.PacketAckRoundTripTimeNanosAvgTime.Describe(ch)
	d.FlushNanosNumOps.Describe(ch)
	d.FlushNanosAvgTime.Describe(ch)
	d.FsyncNanosNumOps.Describe(ch)
	d.FsyncNanosAvgTime.Describe(ch)
	d.SendDataPacketBlockedOnNetworkNanosNumOps.Describe(ch)
	d.SendDataPacketBlockedOnNetworkNanosAvgTime.Describe(ch)
	d.SendDataPacketTransferNanosNumOps.Describe(ch)
	d.SendDataPacketTransferNanosAvgTime.Describe(ch)
}

// Collect 通用方法
func (d *DataNode) Collect(ch chan<- prometheus.Metric) {
	d.BytesWritten.Collect(ch)
	d.BytesRead.Collect(ch)
	d.BlocksWritten.Collect(ch)
	d.BlocksRead.Collect(ch)
	d.BlocksReplicated.Collect(ch)
	d.BlocksRemoved.Collect(ch)
	d.BlocksVerified.Collect(ch)
	d.BlockVerificationFailures.Collect(ch)
	d.BlocksCached.Collect(ch)
	d.BlocksUncached.Collect(ch)
	d.ReadsFromLocalClient.Collect(ch)
	d.ReadsFromRemoteClient.Collect(ch)
	d.WritesFromLocalClient.Collect(ch)
	d.WritesFromRemoteClient.Collect(ch)
	d.BlocksGetLocalPathInfo.Collect(ch)
	d.FsyncCount.Collect(ch)
	d.VolumeFailures.Collect(ch)
	d.ReadBlockOpNumOps.Collect(ch)
	d.ReadBlockOpAvgTime.Collect(ch)
	d.WriteBlockOpNumOps.Collect(ch)
	d.WriteBlockOpAvgTime.Collect(ch)
	d.BlockChecksumOpNumOps.Collect(ch)
	d.BlockChecksumOpAvgTime.Collect(ch)
	d.CopyBlockOpNumOps.Collect(ch)
	d.CopyBlockOpAvgTime.Collect(ch)
	d.ReplaceBlockOpNumOps.Collect(ch)
	d.ReplaceBlockOpAvgTime.Collect(ch)
	d.HeartbeatsNumOps.Collect(ch)
	d.HeartbeatsAvgTime.Collect(ch)
	d.BlockReportsNumOps.Collect(ch)
	d.BlockReportsAvgTime.Collect(ch)
	d.CacheReportsNumOps.Collect(ch)
	d.CacheReportsAvgTime.Collect(ch)
	d.PacketAckRoundTripTimeNanosNumOps.Collect(ch)
	d.PacketAckRoundTripTimeNanosAvgTime.Collect(ch)
	d.FlushNanosNumOps.Collect(ch)
	d.FlushNanosAvgTime.Collect(ch)
	d.FsyncNanosNumOps.Collect(ch)
	d.FsyncNanosAvgTime.Collect(ch)
	d.SendDataPacketBlockedOnNetworkNanosNumOps.Collect(ch)
	d.SendDataPacketBlockedOnNetworkNanosAvgTime.Collect(ch)
	d.SendDataPacketTransferNanosNumOps.Collect(ch)
	d.SendDataPacketTransferNanosAvgTime.Collect(ch)

}
