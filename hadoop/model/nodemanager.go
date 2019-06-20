package model

import "github.com/prometheus/client_golang/prometheus"

// NodeManager jmx中收集的NodeManager信息
type NodeManager struct {
	ContainersLaunched  prometheus.Gauge
	ContainersCompleted prometheus.Gauge
	ContainersFailed    prometheus.Gauge
	ContainersKilled    prometheus.Gauge
	ContainersIniting   prometheus.Gauge
	ContainersRunning   prometheus.Gauge
	AllocatedGB         prometheus.Gauge
	AllocatedContainers prometheus.Gauge
	AvailableGB         prometheus.Gauge
	AllocatedVCores     prometheus.Gauge
	AvailableVCores     prometheus.Gauge
}

// NewNodeManagerExporter 返回 NodeManager 部分的exporter
func NewNodeManagerExporter(namespace, context string) *NodeManager {
	return &NodeManager{
		ContainersLaunched: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "containers_launched",
			Help:      "Total number of launched containers",
		}),
		ContainersCompleted: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "containers_completed",
			Help:      "Total number of successfully completed containers ",
		}),
		ContainersFailed: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "containers_failed",
			Help:      "Total number of failed containers",
		}),
		ContainersKilled: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "containers_killed",
			Help:      "Total number of killed containers",
		}),
		ContainersIniting: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "containers_initing",
			Help:      "Current number of initializing containers",
		}),
		ContainersRunning: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "containers_running",
			Help:      "Current number of running containers",
		}),
		AllocatedGB: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "allocated_gb",
			Help:      "Current allocated memory in GB",
		}),
		AllocatedContainers: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "allocated_containers",
			Help:      "Current number of allocated containers",
		}),
		AvailableGB: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "available_gb",
			Help:      "Current available memory in GB",
		}),
		AllocatedVCores: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "allocated_vcores",
			Help:      "Current used vcores",
		}),
		AvailableVCores: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: context,
			Name:      "available_vcores",
			Help:      "Current available vcores",
		}),
	}
}

// Describe 通用方法
func (n *NodeManager) Describe(ch chan<- *prometheus.Desc) {
	n.ContainersLaunched.Describe(ch)
	n.ContainersCompleted.Describe(ch)
	n.ContainersFailed.Describe(ch)
	n.ContainersKilled.Describe(ch)
	n.ContainersIniting.Describe(ch)
	n.ContainersRunning.Describe(ch)
	n.AllocatedGB.Describe(ch)
	n.AllocatedContainers.Describe(ch)
	n.AvailableGB.Describe(ch)
	n.AllocatedVCores.Describe(ch)
	n.AvailableVCores.Describe(ch)
}

// Collect 通用方法
func (n *NodeManager) Collect(ch chan<- prometheus.Metric) {

	n.ContainersLaunched.Collect(ch)
	n.ContainersCompleted.Collect(ch)
	n.ContainersFailed.Collect(ch)
	n.ContainersKilled.Collect(ch)
	n.ContainersIniting.Collect(ch)
	n.ContainersRunning.Collect(ch)
	n.AllocatedGB.Collect(ch)
	n.AllocatedContainers.Collect(ch)
	n.AvailableGB.Collect(ch)
	n.AllocatedVCores.Collect(ch)
	n.AvailableVCores.Collect(ch)

}
