package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/willfantom/mininet-exporter/collector"
	"github.com/willfantom/mininet-exporter/mininet"
)

var (
	mininetAddress string
	Port           int = 9881
)

const (
	metricsPath = "/metrics"
)

func main() {
	flag.StringVar(&mininetAddress, "address", "http://localhost:8080", "Address in which the Mininet API can be accessed via (not including api prefix)")
	flag.Parse()

	client := mininet.NewClient(mininetAddress)

	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	prometheus.MustRegister(collector.NewPingCollector(client))

	logrus.Infoln("Mininet Exporter Starting...")

	handler := promhttp.Handler()
	http.Handle(metricsPath, handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>Mininet Exporter</title></head>
		<body>
		<h1>Mininet Exporter</h1>
		<p><a href="` + metricsPath + `">Metrics</a></p>
		</body>
		</html>`))
	})

	err := http.ListenAndServe("0.0.0.0:9881", nil)
	if err != nil {
		log.Fatal(err)
	}
}
