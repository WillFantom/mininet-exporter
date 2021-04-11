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
	Sent       int     `json:"sent"`
	Received   int     `json:"received"`
	RTTAverage float64 `json:"rtt_avg,omitempty"`
}

type IperfResponse struct {
	Server      string  `json:"server,omitempty"`
	Client      string  `json:"client,omitempty"`
	ServerSpeed float64 `json:"server_speed"`
	ClientSpeed float64 `json:"client_speed"`
}

type CustomCommandResponse struct {
	NodeName string `json:"node_name,omitempty"`
	Status   string `json:"status"`
	ExitCode int    `json:"exit_code"`
	StdOut   string `json:"stdout,omitempty"`
	StdErr   string `json:"stderr,omitempty"`
}
