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
	"net/http"
	"time"

	"github.com/pkg/errors"
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

type sender interface {
	sendMessage(m *message)
}

// Collector represent interface to Json2Sandesh proxy
type Collector struct {
	sender
}

type noSender struct{}

func (*noSender) sendMessage(m *message) {}

type proxySender struct {
	url    string
	client *http.Client
	queue  chan struct{}
}

// NewCollectorFromConfig makes a collector
func NewCollectorFromConfig() (*Collector, error) {
	if !viper.GetBool("collector.enable") {
		logrus.Warn("collector is disabled")
		return &Collector{
			sender: &noSender{},
		}, nil
	}
	s, err := newProxySenderFromConfig()
	if err != nil {
		return nil, err
	}
	return &Collector{
		sender: s,
	}, nil
}

func newProxySenderFromConfig() (sender, error) {
	s := &proxySender{}
	s.url = viper.GetString("collector.url")
	if s.url == "" {
		return nil, errors.New("collector.url is undefined")
	}
	s.client = &http.Client{
		Timeout: defaultTimeout,
	}
	maxMessageCount := viper.GetInt("collector.queue_size")
	if maxMessageCount == 0 {
		maxMessageCount = defaultMaxMessageCount
	}
	s.queue = make(chan struct{}, maxMessageCount)
	return s, nil
}

func (s *proxySender) sendMessage(m *message) {
	go s.trySendMessage(m)
}

func (s *proxySender) trySendMessage(m *message) {
	select {
	case s.queue <- struct{}{}:
		defer func() { <-s.queue }()
	default:
		logrus.Warnf("collector: message queue is full, dropping %s", m.SandeshType)
		return
	}
	if err := s.postMessage(m); err != nil {
		logrus.WithError(err).Warn("send message to collector failed")
	}
}

func (s *proxySender) postMessage(m *message) error {
	data, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "failed to marshal message")
	}
	req, err := http.NewRequest("POST", s.url, bytes.NewReader(data))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	resp, err := s.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close() // nolint: errcheck
	if resp.StatusCode != http.StatusOK {
		return errors.New("http status " + resp.Status)
	}
	return nil
}
