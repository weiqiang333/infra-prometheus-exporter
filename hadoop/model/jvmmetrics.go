package model

import "github.com/prometheus/client_golang/prometheus"

// JvmMetrics jvm 的一的参数指标
type JvmMetrics struct {
	MemNonHeapUsedM                 prometheus.Gauge
	MemNonHeapCommittedM            prometheus.Gauge
	MemNonHeapMaxM                  prometheus.Gauge
	MemHeapUsedM                    prometheus.Gauge
	MemHeapCommittedM               prometheus.Gauge
	MemHeapMaxM                     prometheus.Gauge
	MemMaxM                         prometheus.Gauge
	ThreadsNew                      prometheus.Gauge
	ThreadsRunnable                 prometheus.Gauge
	ThreadsBlocked                  prometheus.Gauge
	ThreadsWaiting                  prometheus.Gauge
	ThreadsTimedWaiting             prometheus.Gauge
	ThreadsTerminated               prometheus.Gauge
	GcCountParNew                   prometheus.Gauge
	GcTimeMillisParNew              prometheus.Gauge
	GcCountConcurrentMarkSweep      prometheus.Gauge
	GcTimeMillisConcurrentMarkSweep prometheus.Gauge
	GcInfo                          prometheus.Gauge
	GcCount                         prometheus.Gauge
	GcTimeMillis                    prometheus.Gauge
	LogFatal                        prometheus.Gauge
	LogError                        prometheus.Gauge
	LogWarn                         prometheus.Gauge
	LogInfo                         prometheus.Gauge
	GcNumWarnThresholdExceeded      prometheus.Gauge
	GcNumInfoThresholdExceeded      prometheus.Gauge
	GcTotalExtraSleepTime           prometheus.Gauge
}

// NewJvmMetricsExporter 返回 JvmMetrics 部分的exporter
func NewJvmMetricsExporter(namespace string, context string) *JvmMetrics {
	return &JvmMetrics{
		MemNonHeapUsedM: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "mem_non_heap_used",
			Help:      "Current non-heap memory used in MB",
		}),
		MemNonHeapCommittedM: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "mem_non_heap_committed",
			Help:      "Current non-heap memory committed in MB",
		}),
		MemNonHeapMaxM: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "mem_non_heap_max",
			Help:      "Max non-heap memory size in MB",
		}),
		MemHeapUsedM: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "mem_heap_used",
			Help:      "Current heap memory used in MB",
		}),
		MemHeapCommittedM: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "mem_heap_committed",
			Help:      "Current heap memory committed in MB",
		}),
		MemHeapMaxM: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "mem_heap_max",
			Help:      "Max heap memory size in MB",
		}),
		MemMaxM: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "mem_max",
			Help:      "Max memory size in MB",
		}),
		ThreadsNew: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "threads_new",
			Help:      "Current number of NEW threads",
		}),
		ThreadsRunnable: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "threads_runnable",
			Help:      "Current number of RUNNABLE threads",
		}),
		ThreadsBlocked: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "threads_blocked",
			Help:      "Current number of BLOCKED threads",
		}),
		ThreadsWaiting: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "threads_waiting",
			Help:      "Current number of WAITING threads",
		}),
		ThreadsTimedWaiting: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "threads_timed_waiting",
			Help:      "Current number of TIMED_WAITING threads",
		}),
		ThreadsTerminated: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "threads_terminated",
			Help:      "Current number of TERMINATED threads",
		}),
		GcCountParNew: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_count_par_new",
			Help:      "New 's Total GC count",
		}),
		GcTimeMillisParNew: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_time_millis_par_new",
			Help:      "New 's Total GC time in msec",
		}),
		GcCountConcurrentMarkSweep: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_count_concurrent_mark_sweep",
			Help:      "Total GC count Concurrent",
		}),
		GcTimeMillisConcurrentMarkSweep: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_time_millis_concurrent_mark_sweep",
			Help:      "Concurrent Total GC time in msec",
		}),
		GcInfo: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_info",
			Help:      "Total GC count and GC time in msec, grouped by the kind of GC.  ex.) GcCountPS Scavenge=6, GCTimeMillisPS Scavenge=40, GCCountPS MarkSweep=0, GCTimeMillisPS MarkSweep=0",
		}),
		GcCount: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_count",
			Help:      "Total GC count",
		}),
		GcTimeMillis: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_time_millis",
			Help:      "Total GC time in msec",
		}),
		LogFatal: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "log_fatal",
			Help:      "Total number of FATAL logs",
		}),
		LogError: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "log_error",
			Help:      "Total number of ERROR logs",
		}),
		LogWarn: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "log_warn",
			Help:      "Total number of WARN logs",
		}),
		LogInfo: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "log_info",
			Help:      "Total number of INFO logs",
		}),
		GcNumWarnThresholdExceeded: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_num_warn_threshold_exceeded",
			Help:      "Number of times that the GC warn threshold is exceeded",
		}),
		GcNumInfoThresholdExceeded: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_num_info_threshold_exceeded",
			Help:      "Number of times that the GC info threshold is exceeded",
		}),
		GcTotalExtraSleepTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "gc_total_extra_sleep_time",
			Help:      "Total GC extra sleep time in msec",
		}),
	}
}

// Describe 通用方法
func (j *JvmMetrics) Describe(ch chan<- *prometheus.Desc) {
	j.MemNonHeapUsedM.Describe(ch)
	j.MemNonHeapCommittedM.Describe(ch)
	j.MemNonHeapMaxM.Describe(ch)
	j.MemHeapUsedM.Describe(ch)
	j.MemHeapCommittedM.Describe(ch)
	j.MemHeapMaxM.Describe(ch)
	j.MemMaxM.Describe(ch)
	j.ThreadsNew.Describe(ch)
	j.ThreadsRunnable.Describe(ch)
	j.ThreadsBlocked.Describe(ch)
	j.ThreadsWaiting.Describe(ch)
	j.ThreadsTimedWaiting.Describe(ch)
	j.ThreadsTerminated.Describe(ch)
	j.GcInfo.Describe(ch)
	j.GcCount.Describe(ch)
	j.GcTimeMillis.Describe(ch)
	j.LogFatal.Describe(ch)
	j.LogError.Describe(ch)
	j.LogWarn.Describe(ch)
	j.LogInfo.Describe(ch)
	j.GcNumWarnThresholdExceeded.Describe(ch)
	j.GcNumInfoThresholdExceeded.Describe(ch)
	j.GcTotalExtraSleepTime.Describe(ch)
}

// Collect 通用方法
func (j *JvmMetrics) Collect(ch chan<- prometheus.Metric) {
	j.MemNonHeapUsedM.Collect(ch)
	j.MemNonHeapCommittedM.Collect(ch)
	j.MemNonHeapMaxM.Collect(ch)
	j.MemHeapUsedM.Collect(ch)
	j.MemHeapCommittedM.Collect(ch)
	j.MemHeapMaxM.Collect(ch)
	j.MemMaxM.Collect(ch)
	j.ThreadsNew.Collect(ch)
	j.ThreadsRunnable.Collect(ch)
	j.ThreadsBlocked.Collect(ch)
	j.ThreadsWaiting.Collect(ch)
	j.ThreadsTimedWaiting.Collect(ch)
	j.ThreadsTerminated.Collect(ch)
	j.GcInfo.Collect(ch)
	j.GcCount.Collect(ch)
	j.GcTimeMillis.Collect(ch)
	j.LogFatal.Collect(ch)
	j.LogError.Collect(ch)
	j.LogWarn.Collect(ch)
	j.LogInfo.Collect(ch)
	j.GcNumWarnThresholdExceeded.Collect(ch)
	j.GcNumInfoThresholdExceeded.Collect(ch)
	j.GcTotalExtraSleepTime.Collect(ch)
}
