package model

import "github.com/prometheus/client_golang/prometheus"

// JournalNode metrics
type JournalNode struct {
	Syncs60sNumOps                        prometheus.Gauge
	Syncs60s50thPercentileLatencyMicros   prometheus.Gauge
	Syncs60s75thPercentileLatencyMicros   prometheus.Gauge
	Syncs60s90thPercentileLatencyMicros   prometheus.Gauge
	Syncs60s95thPercentileLatencyMicros   prometheus.Gauge
	Syncs60s99thPercentileLatencyMicros   prometheus.Gauge
	Syncs300sNumOps                       prometheus.Gauge
	Syncs300s50thPercentileLatencyMicros  prometheus.Gauge
	Syncs300s75thPercentileLatencyMicros  prometheus.Gauge
	Syncs300s90thPercentileLatencyMicros  prometheus.Gauge
	Syncs300s95thPercentileLatencyMicros  prometheus.Gauge
	Syncs300s99thPercentileLatencyMicros  prometheus.Gauge
	Syncs3600sNumOps                      prometheus.Gauge
	Syncs3600s50thPercentileLatencyMicros prometheus.Gauge
	Syncs3600s75thPercentileLatencyMicros prometheus.Gauge
	Syncs3600s90thPercentileLatencyMicros prometheus.Gauge
	Syncs3600s95thPercentileLatencyMicros prometheus.Gauge
	Syncs3600s99thPercentileLatencyMicros prometheus.Gauge
	BatchesWritten                        prometheus.Gauge
	TxnsWritten                           prometheus.Gauge
	BytesWritten                          prometheus.Gauge
	BatchesWrittenWhileLagging            prometheus.Gauge
	LastWriterEpoch                       prometheus.Gauge
	CurrentLagTxns                        prometheus.Gauge
	// Id golint 检查不通过，但是jmx 导出的数据就是Id，所以这里不修改
	LastWrittenTxId   prometheus.Gauge
	LastPromisedEpoch prometheus.Gauge
}

// NewJournalNodeExporter 返回 JournalNode 部分的exporter
func NewJournalNodeExporter(namespace, context string) *JournalNode {
	return &JournalNode{
		Syncs60sNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_60s_num_ops",
			Help:      "Number of sync operations (1 minute granularity)",
		}),
		Syncs60s50thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_60s_50th_percentile_latency_micros",
			Help:      "The 50th percentile of sync latency in microseconds (1 minute granularity)",
		}),
		Syncs60s75thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_60s_75th_percentile_latency_micros",
			Help:      "The 75th percentile of sync latency in microseconds (1 minute granularity)",
		}),
		Syncs60s90thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_60s_90th_percentile_latency_micros",
			Help:      "The 90th percentile of sync latency in microseconds (1 minute granularity)",
		}),
		Syncs60s95thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_60s_95th_percentile_latency_micros",
			Help:      "The 95th percentile of sync latency in microseconds (1 minute granularity)",
		}),
		Syncs60s99thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_60s_99th_percentile_latency_micros",
			Help:      "The 99th percentile of sync latency in microseconds (1 minute granularity)",
		}),
		Syncs300sNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_300s_num_ops",
			Help:      "Number of sync operations (5 minutes granularity)",
		}),
		Syncs300s50thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_300s_50th_percentile_latency_micros",
			Help:      "The 50th percentile of sync latency in microseconds (5 minutes granularity)",
		}),
		Syncs300s75thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_300s_75th_percentile_latency_micros",
			Help:      "The 75th percentile of sync latency in microseconds (5 minutes granularity)",
		}),
		Syncs300s90thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_300s_90th_percentile_latency_micros",
			Help:      "The 90th percentile of sync latency in microseconds (5 minutes granularity)",
		}),
		Syncs300s95thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_300s_95th_percentile_latency_micros",
			Help:      "The 95th percentile of sync latency in microseconds (5 minutes granularity)",
		}),
		Syncs300s99thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_300s_99th_percentile_latency_micros",
			Help:      "The 99th percentile of sync latency in microseconds (5 minutes granularity)",
		}),
		Syncs3600sNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_3600s_num_ops",
			Help:      "Number of sync operations (1 hour granularity)",
		}),
		Syncs3600s50thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_3600s_50th_percentile_latency_micros",
			Help:      "The 50th percentile of sync latency in microseconds (1 hour granularity)",
		}),
		Syncs3600s75thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_3600s_75th_percentile_latency_micros",
			Help:      "The 75th percentile of sync latency in microseconds (1 hour granularity)",
		}),
		Syncs3600s90thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_3600s_90th_percentile_latency_micros",
			Help:      "The 90th percentile of sync latency in microseconds (1 hour granularity)",
		}),
		Syncs3600s95thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_3600s_95th_percentile_latency_micros",
			Help:      "The 95th percentile of sync latency in microseconds (1 hour granularity)",
		}),
		Syncs3600s99thPercentileLatencyMicros: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "syncs_3600s_99th_percentile_latency_micros",
			Help:      "The 99th percentile of sync latency in microseconds (1 hour granularity)",
		}),
		BatchesWritten: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "batches_written",
			Help:      "Total number of batches written since startup",
		}),
		TxnsWritten: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "txns_written",
			Help:      "Total number of transactions written since startup",
		}),
		BytesWritten: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "bytes_written",
			Help:      "Total number of bytes written since startup",
		}),
		BatchesWrittenWhileLagging: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "batches_written_while_lagging",
			Help:      "Total number of batches written where this node was lagging",
		}),
		LastWriterEpoch: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "last_writer_epoch",
			Help:      "Current writer’s epoch number",
		}),
		CurrentLagTxns: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "current_lag_txns",
			Help:      "The number of transactions that this JournalNode is lagging",
		}),
		LastWrittenTxId: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "last_written_tx_id",
			Help:      "The highest transaction id stored on this JournalNode",
		}),
		LastPromisedEpoch: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "last_promised_epoch",
			Help:      "The last epoch number which this node has promised not to accept any lower epoch, or 0 if no promises have been made",
		}),
	}
}

// Describe 通用方法
func (j *JournalNode) Describe(ch chan<- *prometheus.Desc) {
	j.Syncs60sNumOps.Describe(ch)
	j.Syncs60s50thPercentileLatencyMicros.Describe(ch)
	j.Syncs60s75thPercentileLatencyMicros.Describe(ch)
	j.Syncs60s90thPercentileLatencyMicros.Describe(ch)
	j.Syncs60s95thPercentileLatencyMicros.Describe(ch)
	j.Syncs60s99thPercentileLatencyMicros.Describe(ch)
	j.Syncs300sNumOps.Describe(ch)
	j.Syncs300s50thPercentileLatencyMicros.Describe(ch)
	j.Syncs300s75thPercentileLatencyMicros.Describe(ch)
	j.Syncs300s90thPercentileLatencyMicros.Describe(ch)
	j.Syncs300s95thPercentileLatencyMicros.Describe(ch)
	j.Syncs300s99thPercentileLatencyMicros.Describe(ch)
	j.Syncs3600sNumOps.Describe(ch)
	j.Syncs3600s50thPercentileLatencyMicros.Describe(ch)
	j.Syncs3600s75thPercentileLatencyMicros.Describe(ch)
	j.Syncs3600s90thPercentileLatencyMicros.Describe(ch)
	j.Syncs3600s95thPercentileLatencyMicros.Describe(ch)
	j.Syncs3600s99thPercentileLatencyMicros.Describe(ch)
	j.BatchesWritten.Describe(ch)
	j.TxnsWritten.Describe(ch)
	j.BytesWritten.Describe(ch)
	j.BatchesWrittenWhileLagging.Describe(ch)
	j.LastWriterEpoch.Describe(ch)
	j.CurrentLagTxns.Describe(ch)
	j.LastWrittenTxId.Describe(ch)
	j.LastPromisedEpoch.Describe(ch)

}

// Collect 通用方法
func (j *JournalNode) Collect(ch chan<- prometheus.Metric) {
	j.Syncs60sNumOps.Collect(ch)
	j.Syncs60s50thPercentileLatencyMicros.Collect(ch)
	j.Syncs60s75thPercentileLatencyMicros.Collect(ch)
	j.Syncs60s90thPercentileLatencyMicros.Collect(ch)
	j.Syncs60s95thPercentileLatencyMicros.Collect(ch)
	j.Syncs60s99thPercentileLatencyMicros.Collect(ch)
	j.Syncs300sNumOps.Collect(ch)
	j.Syncs300s50thPercentileLatencyMicros.Collect(ch)
	j.Syncs300s75thPercentileLatencyMicros.Collect(ch)
	j.Syncs300s90thPercentileLatencyMicros.Collect(ch)
	j.Syncs300s95thPercentileLatencyMicros.Collect(ch)
	j.Syncs300s99thPercentileLatencyMicros.Collect(ch)
	j.Syncs3600sNumOps.Collect(ch)
	j.Syncs3600s50thPercentileLatencyMicros.Collect(ch)
	j.Syncs3600s75thPercentileLatencyMicros.Collect(ch)
	j.Syncs3600s90thPercentileLatencyMicros.Collect(ch)
	j.Syncs3600s95thPercentileLatencyMicros.Collect(ch)
	j.Syncs3600s99thPercentileLatencyMicros.Collect(ch)
	j.BatchesWritten.Collect(ch)
	j.TxnsWritten.Collect(ch)
	j.BytesWritten.Collect(ch)
	j.BatchesWrittenWhileLagging.Collect(ch)
	j.LastWriterEpoch.Collect(ch)
	j.CurrentLagTxns.Collect(ch)
	j.LastWrittenTxId.Collect(ch)
	j.LastPromisedEpoch.Collect(ch)
}
