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

// Message represents a message to Json2Sandesh proxy
type Message struct {
	SandeshType string      `json:"sandesh_type"`
	Payload     interface{} `json:"payload"`
}

// MessageBuilder is interface to build the message
type MessageBuilder interface {
	Build() *Message
}

type emptyMessageBuilder struct{}

// NewEmptyMessageBuilder returns nil message builder
func NewEmptyMessageBuilder() MessageBuilder { return &emptyMessageBuilder{} }

func (*emptyMessageBuilder) Build() *Message { return nil }

// Config represents parameters of Сollector in the config file
type Config struct {
	Enabled bool
	URL     string
}

// Collector represent interface to Json2Sandesh proxy
type Collector interface {
	Send(m MessageBuilder)
}

type noSender struct{}

func (*noSender) Send(_ MessageBuilder) {}

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

func (s *proxySender) Send(b MessageBuilder) {
	m := b.Build()
	if m == nil {
		return
	}
	go func() {
		if err := s.postMessage(m); err != nil {
			ignoreAPIMessage().WithError(err).Warn("send Message to collector failed")
		}
	}()
	// sendMessage does not return an error, since asynchronous postMessage usage is suggested
}

func (s *proxySender) postMessage(m *Message) error {
	data, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "failed to marshal Message")
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
