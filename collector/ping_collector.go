package collector

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	"github.com/willfantom/mininet-exporter/mininet"
)

type PingTest struct {
	Sender     string
	Target     string
	MaxRTT     float64
	MinRatio   float64
	testNumber int
}

type PingCollector struct {
	client          *mininet.Client
	tests           []PingTest
	PacketsSent     *prometheus.Desc
	PacketsReceived *prometheus.Desc
	RoundTripAvg    *prometheus.Desc
	Success         *prometheus.Desc
}

func NewPingCollector(client *mininet.Client, tests []PingTest) *PingCollector {
	logrus.Traceln("🛠	defining ping collector")
	specificNamespace := "ping"
	return &PingCollector{
		client: client,
		tests:  tests,
		PacketsSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, specificNamespace, "packets_sent"),
			"Number of ping packets sent",
			[]string{"sender", "target", "test", "test_number"},
			prometheus.Labels{
				"exporter_name": client.Name,
			},
		),
		PacketsReceived: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, specificNamespace, "packets_received"),
			"Number of ping packets received",
			[]string{"sender", "target", "test", "test_number"},
			prometheus.Labels{
				"exporter_name": client.Name,
			},
		),
		RoundTripAvg: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, specificNamespace, "rtt_average"),
			"The average round trip time of the ping",
			[]string{"sender", "target", "test", "test_number"},
			prometheus.Labels{
				"exporter_name": client.Name,
			},
		),
		Success: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, specificNamespace, "success"),
			"Was the ping test criteria met",
			[]string{"sender", "target", "test", "test_number"},
			prometheus.Labels{
				"exporter_name": client.Name,
			},
		),
	}
}

func (pc *PingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pc.PacketsSent
	ch <- pc.PacketsReceived
	ch <- pc.RoundTripAvg
	ch <- pc.Success
}

func (pc *PingCollector) Collect(ch chan<- prometheus.Metric) {
	logrus.Traceln("👀	collecting ping data...")

	for idx, test := range pc.tests {
		testName := test.Sender + "-" + test.Target
		pingData, err := pc.client.Ping(test.Sender, test.Target)
		if err != nil {
			logrus.WithField("message", err).Errorln("⚠️	could not run pingall on topology")
			continue
		}
		sender := pingData.Sender
		target := pingData.Target
		success := 0
		if pingData.Sent > 0 && pingData.Received > 0 {
			if float64(pingData.Received/pingData.Sent) >= test.MinRatio {
				if pingData.RTTAverage <= test.MaxRTT {
					success = 1
				}
			}
		}
		ch <- prometheus.MustNewConstMetric(
			pc.PacketsSent,
			prometheus.GaugeValue,
			float64(pingData.Sent),
			sender, target, testName, fmt.Sprintf("%d", idx),
		)

		ch <- prometheus.MustNewConstMetric(
			pc.PacketsReceived,
			prometheus.GaugeValue,
			float64(pingData.Received),
			sender, target, testName, fmt.Sprintf("%d", idx),
		)
		ch <- prometheus.MustNewConstMetric(
			pc.RoundTripAvg,
			prometheus.GaugeValue,
			float64(pingData.RTTAverage),
			sender, target, testName, fmt.Sprintf("%d", idx),
		)

		ch <- prometheus.MustNewConstMetric(
			pc.Success,
			prometheus.GaugeValue,
			float64(success),
			sender, target, testName, fmt.Sprintf("%d", idx),
		)

	}

	logrus.Traceln("✅	ping data collected")

}
