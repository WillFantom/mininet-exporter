package mininet

import (
	"net/http"
	"net/url"
)

const (
	APIPrefix string = "/mn/api"
)

type Client struct {
	BaseURL    *url.URL
	httpClient http.Client
	Name       string
}

type PingResponse struct {
	Sender     string  `json:"sender,omitempty"`
	Target     string  `json:"target,omitempty"`
	Sent       int     `json:"sent,omitempty"`
	Received   int     `json:"received,omitempty"`
	RTTAverage float64 `json:"rtt_avg,omitempty"`
}

type IperfResponse struct {
	Server      string  `json:"server,omitempty"`
	Client      string  `json:"client,omitempty"`
	ServerSpeed float64 `json:"server_speed,omitempty"`
	ClientSpeed float64 `json:"client_speed,omitempty"`
}

type CustomCommandResponse struct {
	NodeName string `json:"node_name,omitempty"`
	Status   string `json:"status,omitempty"`
	ExitCode int    `json:"exit_code,omitempty"`
	StdOut   string `json:"stdout,omitempty"`
	StdErr   string `json:"stderr,omitempty"`
}
