package main

import (
	"fmt"
	"net/http"

	version "github.com/hashicorp/go-version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

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
	logrus.Infoln("Mininet Exporter Starting...")
	logrus.Debugln("#️⃣	version: ", Version)

	if Version != "rolling" {
		if _, err := version.NewVersion(Version); err != nil {
			logrus.Fatalln("Invalid exporter version")
		}
	}

	client := mininet.NewClient(viper.GetString("MininetTarget"))
	registration(client)

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

	if err := http.ListenAndServe(viper.GetString("ServeAddress")+":"+fmt.Sprintf("%d", viper.GetInt("ServePort")), nil); err != nil {
		logrus.WithField("err msg", err.Error()).Fatalln("🆘	http server failed: exiting")
	}
}

func registration(client *mininet.Client) {
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	if viper.GetBool("PingAllTest") {
		prometheus.MustRegister(collector.NewPingCollector(client))
	}
}
