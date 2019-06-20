package model

import "github.com/prometheus/client_golang/prometheus"

// ClusterMetrics
type ClusterMetrics struct {
	NumActiveNMs           prometheus.Gauge
	NumDecommissioningNMs  prometheus.Gauge
	NumDecommissionedNMs   prometheus.Gauge
	NumShutdownNMs         prometheus.Gauge
	NumLostNMs             prometheus.Gauge
	NumUnhealthyNMs        prometheus.Gauge
	NumRebootedNMs         prometheus.Gauge
	AMLaunchDelayNumOps    prometheus.Gauge
	AMLaunchDelayAvgTime   prometheus.Gauge
	AMRegisterDelayNumOps  prometheus.Gauge
	AMRegisterDelayAvgTime prometheus.Gauge
}

// NewClusterMetricsExporter 返回 ClusterMetrics 部分的exporter
func NewClusterMetricsExporter(namespace, context string) *ClusterMetrics {
	return &ClusterMetrics{
		NumActiveNMs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "num_active_nms",
			Help:      "Current number of active NodeManagers",
		}),
		NumDecommissioningNMs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "num_decommissioning_nms",
			Help:      "Current number of NodeManagers being decommissioned",
		}),
		NumDecommissionedNMs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "num_decommissioned_nms",
			Help:      "Current number of decommissioned NodeManagers",
		}),
		NumShutdownNMs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "num_shutdown_nms",
			Help:      "Current number of NodeManagers shut down gracefully. Note that this does not count NodeManagers that are forcefully killed.",
		}),
		NumLostNMs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "num_lost_nms",
			Help:      "Current number of lost NodeManagers for not sending heartbeats.",
		}),
		NumUnhealthyNMs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "NumUnhealthyNMs",
			Help:      "Current number of unhealthy NodeManagers",
		}),
		NumRebootedNMs: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "NumRebootedNMs",
			Help:      "Current number of rebooted NodeManagers",
		}),
		AMLaunchDelayNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "AMLaunchDelayNumOps",
			Help:      "Total number of AMs launched",
		}),
		AMLaunchDelayAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "AMLaunchDelayAvgTime",
			Help:      "Average time in milliseconds RM spends to launch AM containers after the AM container is allocated",
		}),
		AMRegisterDelayNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "AMRegisterDelayNumOps",
			Help:      "Total number of AMs registered",
		}),
		AMRegisterDelayAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "AMRegisterDelayAvgTime",
			Help:      "Average time in milliseconds AM spends to register with RM after the AM container gets launched",
		}),
	}
}

// Describe 通用方法
func (c *ClusterMetrics) Describe(ch chan<- *prometheus.Desc) {

	c.NumActiveNMs.Describe(ch)
	c.NumDecommissioningNMs.Describe(ch)
	c.NumDecommissionedNMs.Describe(ch)
	c.NumShutdownNMs.Describe(ch)
	c.NumLostNMs.Describe(ch)
	c.NumUnhealthyNMs.Describe(ch)
	c.NumRebootedNMs.Describe(ch)
	c.AMLaunchDelayNumOps.Describe(ch)
	c.AMLaunchDelayAvgTime.Describe(ch)
	c.AMRegisterDelayNumOps.Describe(ch)
	c.AMRegisterDelayAvgTime.Describe(ch)

}

// Collect 通用方法
func (c *ClusterMetrics) Collect(ch chan<- prometheus.Metric) {
	c.NumActiveNMs.Collect(ch)
	c.NumDecommissioningNMs.Collect(ch)
	c.NumDecommissionedNMs.Collect(ch)
	c.NumShutdownNMs.Collect(ch)
	c.NumLostNMs.Collect(ch)
	c.NumUnhealthyNMs.Collect(ch)
	c.NumRebootedNMs.Collect(ch)
	c.AMLaunchDelayNumOps.Collect(ch)
	c.AMLaunchDelayAvgTime.Collect(ch)
	c.AMRegisterDelayNumOps.Collect(ch)
	c.AMRegisterDelayAvgTime.Collect(ch)
}

// QueueMetrics
type QueueMetrics struct {
	Running_0                    *prometheus.Desc
	Running_60                   *prometheus.Desc
	Running_300                  *prometheus.Desc
	Running_1440                 *prometheus.Desc
	AppsSubmitted                *prometheus.Desc
	AppsRunning                  *prometheus.Desc
	AppsPending                  *prometheus.Desc
	AppsCompleted                *prometheus.Desc
	AppsKilled                   *prometheus.Desc
	AppsFailed                   *prometheus.Desc
	AllocatedMB                  *prometheus.Desc
	AllocatedVCores              *prometheus.Desc
	AllocatedContainers          *prometheus.Desc
	AggregateContainersAllocated *prometheus.Desc
	AggregateContainersReleased  *prometheus.Desc
	AvailableMB                  *prometheus.Desc
	AvailableVCores              *prometheus.Desc
	PendingMB                    *prometheus.Desc
	PendingVCores                *prometheus.Desc
	PendingContainers            *prometheus.Desc
	ReservedMB                   *prometheus.Desc
	ReservedVCores               *prometheus.Desc
	ReservedContainers           *prometheus.Desc
	ActiveUsers                  *prometheus.Desc
	ActiveApplications           *prometheus.Desc
	FairShareMB                  *prometheus.Desc
	FairShareVCores              *prometheus.Desc
	MinShareMB                   *prometheus.Desc
	MinShareVCores               *prometheus.Desc
	MaxShareMB                   *prometheus.Desc
	MaxShareVCores               *prometheus.Desc
}

// NewQueueMetricsExporter 返回QueueMetrics 的exporter
func NewQueueMetricsExporter(namespace, context string) *QueueMetrics {

	return &QueueMetrics{

		Running_0: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "running_0"),
			"Current number of running applications whose elapsed time are less than 60 minutes.",
			[]string{"queue", "user"},
			nil,
		),
		Running_60: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "running_60"),
			"Current number of running applications whose elapsed time are between 60 and 300 minutes.",
			[]string{"queue", "user"},
			nil,
		),
		Running_300: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "running_300"),
			"Current number of running applications whose elapsed time are between 300 and 1440 minutes.",
			[]string{"queue", "user"},
			nil,
		),
		Running_1440: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "running_1440"),
			"Current number of running applications elapsed time are more than 1440 minutes.",
			[]string{"queue", "user"},
			nil,
		),
		AppsSubmitted: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "apps_submitted"),
			"Total number of submitted applications.",
			[]string{"queue", "user"},
			nil,
		),
		AppsRunning: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "apps_running"),
			"Current number of running applications.",
			[]string{"queue", "user"},
			nil,
		),
		AppsPending: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "apps_pending"),
			"Current number of applications that have not yet been assigned by any containers.",
			[]string{"queue", "user"},
			nil,
		),
		AppsCompleted: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "apps_completed"),
			"Total number of completed applications.",
			[]string{"queue", "user"},
			nil,
		),
		AppsKilled: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "apps_killed"),
			"Total number of killed applications.",
			[]string{"queue", "user"},
			nil,
		),
		AppsFailed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "apps_failed"),
			"Total number of failed applications.",
			[]string{"queue", "user"},
			nil,
		),
		AllocatedMB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "allocated_mb"),
			"Current allocated memory in MB.",
			[]string{"queue", "user"},
			nil,
		),
		AllocatedVCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "allocated_vcores"),
			"Current allocated CPU in virtual cores.",
			[]string{"queue", "user"},
			nil,
		),
		AllocatedContainers: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "allocated_containers"),
			"Current number of allocated containers.",
			[]string{"queue", "user"},
			nil,
		),
		AggregateContainersAllocated: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "aggregate_containers_allocated"),
			"Total number of allocated containers.",
			[]string{"queue", "user"},
			nil,
		),
		AggregateContainersReleased: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "aggregate_containers_released"),
			"Total number of released containers.",
			[]string{"queue", "user"},
			nil,
		),
		AvailableMB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "available_mb"),
			"Current available memory in MB.",
			[]string{"queue", "user"},
			nil,
		),
		AvailableVCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "available_vcores"),
			"Current available CPU in virtual cores.",
			[]string{"queue", "user"},
			nil,
		),
		PendingMB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "pending_mb"),
			"Current memory requests in MB that are pending to be fulfilled by the scheduler.",
			[]string{"queue", "user"},
			nil,
		),
		PendingVCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "pending_vcores"),
			"Current CPU requests in virtual cores that are pending to be fulfilled by the scheduler.",
			[]string{"queue", "user"},
			nil,
		),
		PendingContainers: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "pending_containers"),
			"Current number of containers that are pending to be fulfilled by the scheduler.",
			[]string{"queue", "user"},
			nil,
		),
		ReservedMB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "reserved_mb"),
			"Current reserved memory in MB.",
			[]string{"queue", "user"},
			nil,
		),
		ReservedVCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "reserved_vcores"),
			"Current reserved CPU in virtual cores.",
			[]string{"queue", "user"},
			nil,
		),
		ReservedContainers: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "reserved_containers"),
			"Current number of reserved containers.",
			[]string{"queue", "user"},
			nil,
		),
		ActiveUsers: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "active_users"),
			"Current number of active users.",
			[]string{"queue", "user"},
			nil,
		),
		ActiveApplications: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "active_applications"),
			"Current number of active applications.",
			[]string{"queue", "user"},
			nil,
		),
		FairShareMB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "fair_share_mb"),
			"(FairScheduler only) Current fair share of memory in MB.",
			[]string{"queue", "user"},
			nil,
		),
		FairShareVCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "fair_share_vcores"),
			"(FairScheduler only) Current fair share of CPU in virtual cores.",
			[]string{"queue", "user"},
			nil,
		),
		MinShareMB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "min_share_mb"),
			"(FairScheduler only) Minimum share of memory in MB.",
			[]string{"queue", "user"},
			nil,
		),
		MinShareVCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "min_share_vcores"),
			"(FairScheduler only) Minimum share of CPU in virtual cores.",
			[]string{"queue", "user"},
			nil,
		),
		MaxShareMB: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "max_share_mb"),
			"(FairScheduler only) Maximum share of memory in MB.",
			[]string{"queue", "user"},
			nil,
		),
		MaxShareVCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "max_share_vcores"),
			"(FairScheduler only) Maximum share of CPU in virtual cores.",
			[]string{"queue", "user"},
			nil,
		),
	}
}

// Describe 通用方法
func (q *QueueMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- q.Running_0
	ch <- q.Running_60
	ch <- q.Running_300
	ch <- q.Running_1440
	ch <- q.AppsSubmitted
	ch <- q.AppsRunning
	ch <- q.AppsPending
	ch <- q.AppsCompleted
	ch <- q.AppsKilled
	ch <- q.AppsFailed
	ch <- q.AllocatedMB
	ch <- q.AllocatedVCores
	ch <- q.AllocatedContainers
	ch <- q.AggregateContainersAllocated
	ch <- q.AggregateContainersReleased
	ch <- q.AvailableMB
	ch <- q.AvailableVCores
	ch <- q.PendingMB
	ch <- q.PendingVCores
	ch <- q.PendingContainers
	ch <- q.ReservedMB
	ch <- q.ReservedVCores
	ch <- q.ReservedContainers
	ch <- q.ActiveUsers
	ch <- q.ActiveApplications
	ch <- q.FairShareMB
	ch <- q.FairShareVCores
	ch <- q.MinShareMB
	ch <- q.MinShareVCores
	ch <- q.MaxShareMB
	ch <- q.MaxShareVCores

}
