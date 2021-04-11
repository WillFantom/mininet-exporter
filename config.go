package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/willfantom/mininet-exporter/collector"
)

type Config struct {
	Debug         bool
	Trace         bool
	ServeAddress  string `mapstructure:"path_map"`
	ServePort     int
	MininetTarget string
	PingAllTest   bool
	PingTests     []collector.PingTest
}

var (
	ExporterConfig   Config
	DefaultServePort int = 9881
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: false,
	})
	logrus.SetLevel(logrus.DebugLevel)
	setDefaultConfiguration()
	readConfiguration()
	if !viper.GetBool("Debug") {
		logrus.SetLevel(logrus.InfoLevel)
	}
	if viper.GetBool("Trace") {
		logrus.SetLevel(logrus.TraceLevel)
	}
}

func setDefaultConfiguration() {
	viper.SetDefault("Debug", false)
	viper.SetDefault("Trace", false)
	viper.SetDefault("ServeAddress", "0.0.0.0")
	viper.SetDefault("ServePort", DefaultServePort)
	viper.SetDefault("MininetTarget", "http://localhost:8080")
	viper.SetDefault("PingAllTest", false)
}

func readConfiguration() {
	viper.SetConfigName("tests")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/config/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Debugln("⚠️	no config file found: using defaults")
		} else {
			logrus.WithField("err msg", err.Error()).Fatalln("🆘	could not read in given configuration file: exiting...")
		}
	}
	viper.Unmarshal(&ExporterConfig)
}
