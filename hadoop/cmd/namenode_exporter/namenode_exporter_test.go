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
		"namenode_rpc_num_open_connections": `
		# HELP namenode_rpc_num_open_connections Current number of open connections
		# TYPE namenode_rpc_num_open_connections gauge
		namenode_rpc_num_open_connections 201
		`,
		"namenode_dfs_state": `
		# HELP namenode_dfs_state Indicate namenode state (0 - standby, 1 - active).
		# TYPE namenode_dfs_state gauge
		namenode_dfs_state 0
		`,
		// 下面这段测试不过
		"namenode_rpc_authorization_failures": `
		# HELP namenode_rpc_authorization_failures Total number of authorization failures
		# TYPE namenode_rpc_authorization_failures gauge
		namenode_rpc_num_open_connections 0
		`,
	}
	for k, v := range excepts {
		err := testutil.CollectAndCompare(e, strings.NewReader(v), k)
		if err != nil {
			t.Fatal(err)
		}
	}

}
