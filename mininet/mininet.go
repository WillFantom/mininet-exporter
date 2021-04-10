package mininet

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/sirupsen/logrus"
)

func NewClient(address string) *Client {

	baseURL, err := url.Parse(address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"given address": address,
			"message":       err.Error(),
		}).Fatalln("Could not create Mininet API client")
	}
	logrus.Infoln("⚙️	mininet target: ", baseURL.String())

	var c Client

	c.BaseURL = baseURL
	c.httpClient = http.Client{}

	rand.Seed(time.Now().UnixNano())
	c.Name = petname.Generate(1, "")

	return &c
}

func (c *Client) PingAll() (map[string]PingResponse, error) {

	path := &url.URL{Path: APIPrefix + "/pingall"}
	fullURL := c.BaseURL.ResolveReference(path)
	request, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"full url": fullURL.String(),
			"message":  err.Error(),
		}).Errorln("Could not create ping all request")
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"full url": fullURL.String(),
			"message":  err.Error(),
		}).Errorln("Failed ping all request")
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logrus.WithFields(logrus.Fields{
			"full url":        fullURL.String(),
			"response status": response.StatusCode,
			"message":         err.Error(),
		}).Warnln("Failed ping all request")
		return nil, err
	}

	var responses map[string]PingResponse
	if err := json.NewDecoder(response.Body).Decode(&responses); err != nil {
		logrus.WithFields(logrus.Fields{
			"full url":        fullURL.String(),
			"response status": response.StatusCode,
			"message":         err.Error(),
		}).Errorln("Failed to parse ping all response")
		return nil, err
	}

	return responses, nil
}
