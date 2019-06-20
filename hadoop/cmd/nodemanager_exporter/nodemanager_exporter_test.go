package main

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestExporter_Collect(t *testing.T) {
	e := NewExporter()
	t.Log(e)

	excepts := map[string]string{
		//yarn
		"nodemanager_yarn_containers_launched": `
		# HELP nodemanager_yarn_containers_launched Total number of launched containers
		# TYPE nodemanager_yarn_containers_launched gauge
		nodemanager_yarn_containers_launched 0
		`,
		// jvm
		"nodemanager_jvm_threads_new": `
		# HELP nodemanager_jvm_threads_new Current number of NEW threads
		# TYPE nodemanager_jvm_threads_new gauge
		nodemanager_jvm_threads_new 0
		`,
		// rpc
		"nodemanager_rpc_call_queue_length": `
		# HELP nodemanager_rpc_call_queue_length Current length of the call queue.
		# TYPE nodemanager_rpc_call_queue_length gauge
		nodemanager_rpc_call_queue_length{port="33651"} 0
		nodemanager_rpc_call_queue_length{port="8040"} 0
		`,
		// 下面这段测试不过
		//"namenode_rpc_authorization_failures": `
		//# HELP namenode_rpc_authorization_failures Total number of authorization failures
		//# TYPE namenode_rpc_authorization_failures gauge
		//namenode_rpc_num_open_connections 0
		//`,
	}
	for k, v := range excepts {
		err := testutil.CollectAndCompare(e, strings.NewReader(v), k)
		if err != nil {
			t.Fatal(err)
		}
	}
}
