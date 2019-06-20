package hadoop

import (
	"flag"
)

var (
	//ListenAddress exporter 的监听地址端口
	ListenAddress string
	//MetricsPath /metrics
	MetricsPath string
	//JmxURL 监听的jmx地址
	JmxURL string
	//Help -h 显示帮助
	Help bool
)

func init() {
	flag.StringVar(&ListenAddress, "web.listen-address", ":9100", "Address on which to expose metrics and web interface.")
	flag.StringVar(&MetricsPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.StringVar(&JmxURL, "web.jmx-url", "http://cnhm0:50070/jmx", "JMX URL")
	flag.BoolVar(&Help, "h", false, "exporter help")
}
