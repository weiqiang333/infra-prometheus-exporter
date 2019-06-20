package hadoop

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
)

// Export 通用的metric 导出
func Export(namespace string) {

	log.Infoln("Starting "+namespace+" exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	http.Handle(MetricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>NameNode Exporter</title></head>
			<body>
			<h1>` + namespace + `</h1>
			<p><a href="` + MetricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Infoln("Listening on", ListenAddress)
	if err := http.ListenAndServe(ListenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
