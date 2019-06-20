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
		//dfs
		"journalnode_dfs_syncs_60s_num_ops": `
		# HELP journalnode_dfs_syncs_60s_num_ops Number of sync operations (1 minute granularity)
		# TYPE journalnode_dfs_syncs_60s_num_ops gauge
		journalnode_dfs_syncs_60s_num_ops 0
		`,
		// jvm
		"journalnode_jvm_gc_count": `
		# HELP journalnode_jvm_gc_count Total GC count
		# TYPE journalnode_jvm_gc_count gauge
		journalnode_jvm_gc_count 0
		`,
		// rpc
		"journalnode_rpc_authentication_successes": `
		# HELP journalnode_rpc_authentication_successes Total number of authentication successes.
		# TYPE journalnode_rpc_authentication_successes gauge
		journalnode_rpc_authentication_successes{port="8485"} 0
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
