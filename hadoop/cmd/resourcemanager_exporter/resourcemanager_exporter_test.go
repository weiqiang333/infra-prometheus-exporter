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
		"resourcemanager_yarn_active_users": `
		# HELP resourcemanager_yarn_active_users Current number of active users.
		# TYPE resourcemanager_yarn_active_users gauge
		resourcemanager_yarn_active_users{queue="root",user="apps"} 0
		resourcemanager_yarn_active_users{queue="root",user="dr.who"} 0
		resourcemanager_yarn_active_users{queue="root.apps",user="apps"} 0
		`,
		// jvm
		"resourcemanager_jvm_log_fatal": `
		# HELP resourcemanager_jvm_log_fatal Total number of FATAL logs
		# TYPE resourcemanager_jvm_log_fatal gauge
		resourcemanager_jvm_log_fatal 0
		`,
		// rpc
		"resourcemanager_rpc_authorization_failures": `
		# HELP resourcemanager_rpc_authorization_failures Total number of authorization failures.
		# TYPE resourcemanager_rpc_authorization_failures gauge
		resourcemanager_rpc_authorization_failures{port="8030"} 0
		resourcemanager_rpc_authorization_failures{port="8031"} 0
		resourcemanager_rpc_authorization_failures{port="8032"} 0
		resourcemanager_rpc_authorization_failures{port="8033"} 0
		`,
	}
	for k, v := range excepts {
		err := testutil.CollectAndCompare(e, strings.NewReader(v), k)
		if err != nil {
			t.Fatal(err)
		}
	}
}
