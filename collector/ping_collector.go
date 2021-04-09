package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	"github.com/willfantom/mininet-exporter/mininet"
)

type PingCollector struct {
	client          *mininet.Client
	PacketsSent     *prometheus.Desc
	PacketsReceived *prometheus.Desc
	RoundTripAvg    *prometheus.Desc
}

func NewPingCollector(client *mininet.Client) *PingCollector {
	specificNamespace := "pingall"
	return &PingCollector{
		client: client,
		PacketsSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, specificNamespace, "packets_sent"),
			"Number of ping packets sent",
			[]string{"sender", "target"},
			nil,
		),
		PacketsReceived: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, specificNamespace, "packets_received"),
			"Number of ping packets received",
			[]string{"sender", "target"},
			nil,
		),
		RoundTripAvg: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, specificNamespace, "rtt_average"),
			"The average round trip time of the ping",
			[]string{"sender", "target"},
			nil,
		),
	}
}

func (pc *PingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pc.PacketsSent
	ch <- pc.PacketsReceived
	ch <- pc.RoundTripAvg
}

func (pc *PingCollector) Collect(ch chan<- prometheus.Metric) {

	pings, err := pc.client.PingAll()
	if err != nil {
		logrus.WithField("message", err).Errorln("could not run pingall on topology")
		return
	}

	for _, pingData := range pings {

		sender := pingData.Sender
		target := pingData.Target

		ch <- prometheus.MustNewConstMetric(
			pc.PacketsSent,
			prometheus.GaugeValue,
			float64(pingData.Sent),
			sender, target,
		)

		ch <- prometheus.MustNewConstMetric(
			pc.PacketsReceived,
			prometheus.GaugeValue,
			float64(pingData.Received),
			sender, target,
		)

		ch <- prometheus.MustNewConstMetric(
			pc.RoundTripAvg,
			prometheus.GaugeValue,
			float64(pingData.RTTAverage),
			sender, target,
		)

	}

}
