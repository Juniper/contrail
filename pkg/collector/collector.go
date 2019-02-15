package collector

/* to use add to config.yml a section:
collector:
  enabled: true
  url: http://collectorproxyhost:port/url
*/

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	defaultTimeout = 30 * time.Second
)

type message struct {
	SandeshType string      `json:"sandesh_type"`
	Payload     interface{} `json:"payload"`
}

type messager interface {
	Get() *message
}

// Config represents parameters of Сollector in the config file
type Config struct {
	Enabled bool
	URL     string
}

// Collector represent interface to Json2Sandesh proxy
type Collector interface {
	Send(m messager)
}

type noSender struct{}

func (*noSender) Send(i messager) {}

type proxySender struct {
	url    string
	client *http.Client
}

// NewCollector makes a collector
func NewCollector(cfg *Config) (Collector, error) {
	if !cfg.Enabled {
		ignoreAPIMessage().Warn("collector is disabled")
		return &noSender{}, nil
	}
	return newProxySender(cfg)
}

func newProxySender(cfg *Config) (Collector, error) {
	if cfg.URL == "" {
		return nil, errors.New("collector.url is undefined")
	}
	s := &proxySender{
		url: cfg.URL,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
	return s, nil
}

// Send sends a message to Json2Sandesh
func (s *proxySender) Send(i messager) {
	m := i.Get()
	if m == nil {
		return
	}
	go func() {
		if err := s.postMessage(m); err != nil {
			ignoreAPIMessage().WithError(err).Warn("send message to collector failed")
		}
	}()
	// sendMessage does not return an error, since asynchronous postMessage usage is suggested
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
