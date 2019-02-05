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

// Config represents parameters of Сollector in the config file
type Config struct {
	Enable    bool
	URL       string
	QueueSize int
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

// NewCollector makes a collector
func NewCollector(cfg *Config) (*Collector, error) {
	if !cfg.Enable {
		logrus.Warn("collector is disabled")
		return &Collector{
			sender: &noSender{},
		}, nil
	}
	s, err := newProxySender(cfg)
	if err != nil {
		return nil, err
	}
	return &Collector{
		sender: s,
	}, nil
}

func newProxySender(cfg *Config) (sender, error) {
	if cfg.URL == "" {
		return nil, errors.New("collector.url is undefined")
	}
	s := &proxySender{
		url: cfg.URL,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
	maxMessageCount := cfg.QueueSize
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
	resp, err := s.client.Post(s.url, "application/json; charset=UTF-8", bytes.NewReader(data))
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close() // nolint: errcheck
	if resp.StatusCode != http.StatusOK {
		return errors.New("http status " + resp.Status)
	}
	return nil
}
