package main

import (
	"flag"
	"log"
	"net/http"

	version "github.com/hashicorp/go-version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/willfantom/mininet-exporter/collector"
	"github.com/willfantom/mininet-exporter/mininet"
)

var (
	mininetAddress string
	Version        string = "0.1.0-alpha"
)

const (
	metricsPath = "/metrics"
)

func main() {
	if Version != "rolling" {
		if _, err := version.NewVersion(Version); err != nil {
			logrus.Fatalln("Invalid exporter version")
		}
	}

	flag.StringVar(&mininetAddress, "address", "http://localhost:8080", "Address in which the Mininet API can be accessed via (not including api prefix)")
	flag.Parse()

	client := mininet.NewClient(mininetAddress)

	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	prometheus.MustRegister(collector.NewPingCollector(client))

	logrus.Infoln("Mininet Exporter Starting...")
	logrus.Infoln("  - version: ", Version)

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
