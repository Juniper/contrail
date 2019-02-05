package collector

/* to use add to config.yml a section:
collector:
  enable: true
  url: http://collectorproxyhost:port/url
  queuesize: 32 # optional
*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	defaultTimeout         = 30 * time.Second
	defaultMaxMessageCount = 16
)

type message struct {
	SandeshType string      `json:"sandesh_type"`
	Payload     interface{} `json:"payload"`
}

// Collector represent interface to Json2Sandesh proxy
type Collector struct {
	enabled bool
	url     string
	client  *http.Client
	queue   chan struct{}
}

// NewCollectorFromConfig makes a collector
func NewCollectorFromConfig() *Collector {
	collector := &Collector{}
	if !viper.GetBool("collector.enable") {
		return collector
	}
	collector.enabled = true
	collector.url = viper.GetString("collector.url")
	if collector.url == "" {
		logrus.Fatal("collector.url not specified in configuration file")
		return nil
	}
	collector.client = &http.Client{
		Timeout: defaultTimeout,
	}
	maxMessageCount := viper.GetInt("collector.queue_size")
	if maxMessageCount == 0 {
		maxMessageCount = defaultMaxMessageCount
	}
	collector.queue = make(chan struct{}, maxMessageCount)
	return collector
}

func (collector *Collector) sendMessage(m *message) {
	if !collector.enabled {
		return
	}
	go collector.trySendMessage(m)
}

func (collector *Collector) trySendMessage(m *message) {
	select {
	case collector.queue <- struct{}{}:
		defer func() { <-collector.queue }()
	default:
		logrus.Warn("collector message queue is full")
		return
	}
	if err := collector.postMessage(m); err != nil {
		logrus.WithError(err).Warn("send message to collector failed")
	}
}

func (collector *Collector) postMessage(m *message) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", collector.url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	resp, err := collector.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint: errcheck
	if resp.StatusCode != http.StatusOK {
		return errors.New("http status " + resp.Status)
	}
	return nil
}
