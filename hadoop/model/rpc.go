package model

import "github.com/prometheus/client_golang/prometheus"

// RPC metrics
type RPC struct {
	ReceivedBytes *prometheus.Desc
	SentBytes     *prometheus.Desc
	// golint 中建议RPC
	RpcQueueTimeNumOps         *prometheus.Desc
	RpcQueueTimeAvgTime        *prometheus.Desc
	RpcProcessingTimeNumOps    *prometheus.Desc
	RpcProcessingTimeAvgTime   *prometheus.Desc
	RpcAuthenticationFailures  *prometheus.Desc
	RpcAuthenticationSuccesses *prometheus.Desc
	RpcAuthorizationFailures   *prometheus.Desc
	RpcAuthorizationSuccesses  *prometheus.Desc
	NumOpenConnections         *prometheus.Desc
	CallQueueLength            *prometheus.Desc
}

// NewRPCExporter 返回 RPC Context 部分的exporter
func NewRPCExporter(namespace, context string) *RPC {
	return &RPC{
		ReceivedBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "received_bytes"),
			"Total number of received bytes .",
			[]string{"port"},
			nil,
		),
		SentBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "sent_bytes"),
			"Total number of sent bytes .",
			[]string{"port"},
			nil,
		),
		RpcQueueTimeNumOps: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_queue_time_num_ops"),
			"Total number of RPC calls .",
			[]string{"port"},
			nil,
		),
		RpcQueueTimeAvgTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_queue_time_avg_time"),
			"Average queue time in milliseconds .",
			[]string{"port"},
			nil,
		),

		RpcProcessingTimeNumOps: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_processing_time_num_ops"),
			"Total number of RPC calls (same to RpcQueueTimeNumOps).",
			[]string{"port"},
			nil,
		),
		RpcProcessingTimeAvgTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_processing_time_avg_time"),
			"Average Processing time in milliseconds.",
			[]string{"port"},
			nil,
		),
		RpcAuthenticationFailures: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_authentication_failures"),
			"Total number of authentication failures.",
			[]string{"port"},
			nil,
		),
		RpcAuthenticationSuccesses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_authentication_successes"),
			"Total number of authentication successes.",
			[]string{"port"},
			nil,
		),
		RpcAuthorizationFailures: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_authorization_failures"),
			"Total number of authorization failures.",
			[]string{"port"},
			nil,
		),
		RpcAuthorizationSuccesses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "rpc_authorization_successes"),
			"Total number of authorization successes.",
			[]string{"port"},
			nil,
		),
		NumOpenConnections: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "num_open_connections"),
			"Current number of open connections.",
			[]string{"port"},
			nil,
		),
		CallQueueLength: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, context, "call_queue_length"),
			"Current length of the call queue.",
			[]string{"port"},
			nil,
		),
	}
}

// Describe 通用方法
func (r *RPC) Describe(ch chan<- *prometheus.Desc) {
	ch <- r.ReceivedBytes
	ch <- r.SentBytes
	ch <- r.RpcQueueTimeNumOps
	ch <- r.RpcQueueTimeAvgTime
	ch <- r.RpcProcessingTimeNumOps
	ch <- r.RpcProcessingTimeAvgTime
	ch <- r.RpcAuthenticationFailures
	ch <- r.RpcAuthenticationSuccesses
	ch <- r.RpcAuthorizationFailures
	ch <- r.RpcAuthorizationSuccesses
	ch <- r.NumOpenConnections
	ch <- r.CallQueueLength
}
