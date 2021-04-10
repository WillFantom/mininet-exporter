package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	"github.com/willfantom/mininet-exporter/mininet"
)

type PingAllCollector struct {
	client            *mininet.Client
	PacketsSent       *prometheus.Desc
	PacketsReceived   *prometheus.Desc
	ReceivedSentRatio *prometheus.Desc
	RoundTripAvg      *prometheus.Desc
}

func NewPingAllCollector(client *mininet.Client) *PingAllCollector {
	logrus.Traceln("🛠	defining pingall collector")
	specificNamespace := "pingall"
	return &PingAllCollector{
		client: client,
		PacketsSent: prometheus.NewDesc(
			prometheus.BuildFQName(getNamespace(client), specificNamespace, "packets_sent"),
			"Number of ping packets sent",
			[]string{"sender", "target"},
			nil,
		),
		PacketsReceived: prometheus.NewDesc(
			prometheus.BuildFQName(getNamespace(client), specificNamespace, "packets_received"),
			"Number of ping packets received",
			[]string{"sender", "target"},
			nil,
		),
		ReceivedSentRatio: prometheus.NewDesc(
			prometheus.BuildFQName(getNamespace(client), specificNamespace, "packets_loss_ratio"),
			"Number of ping packets being received against being sent",
			[]string{"sender", "target"},
			nil,
		),
		RoundTripAvg: prometheus.NewDesc(
			prometheus.BuildFQName(getNamespace(client), specificNamespace, "rtt_average"),
			"The average round trip time of the ping",
			[]string{"sender", "target"},
			nil,
		),
	}
}

func (pc *PingAllCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pc.PacketsSent
	ch <- pc.PacketsReceived
	ch <- pc.RoundTripAvg
}

func (pc *PingAllCollector) Collect(ch chan<- prometheus.Metric) {
	logrus.Traceln("👀	collecting pingall data...")

	pings, err := pc.client.PingAll()
	if err != nil {
		logrus.WithField("message", err).Errorln("⚠️	could not run pingall on topology")
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
			pc.ReceivedSentRatio,
			prometheus.CounterValue,
			float64(pingData.Received/pingData.Sent),
			sender, target,
		)

		ch <- prometheus.MustNewConstMetric(
			pc.RoundTripAvg,
			prometheus.GaugeValue,
			float64(pingData.RTTAverage),
			sender, target,
		)

	}
	logrus.Traceln("✅	pingall data collected")

}
