package analytics

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

	"github.com/Juniper/contrail/pkg/collector"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	defaultTimeout = 30 * time.Second
)

// Config represents parameters of Сollector in the config file
type Config struct {
	Enabled bool
	URL     string
}

// noSender implements collector.Collector
type noSender struct{}

func (*noSender) Send(_ collector.MessageBuilder) {}

// proxySender implements collector.Collector
type proxySender struct {
	url    string
	client *http.Client
}

// NewCollectorFromGlobalConfig makes a collector using global Viper configuration.
func NewCollectorFromGlobalConfig() (collector.Collector, error) {
	var cfg Config
	if err := viper.UnmarshalKey("collector", &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal collector config")
	}
	return NewCollector(&cfg)
}

// NewCollector makes a collector
func NewCollector(cfg *Config) (collector.Collector, error) {
	if !cfg.Enabled {
		ignoreAPIMessage().Info("Collector is disabled")
		return &noSender{}, nil
	}
	return newProxySender(cfg)
}

func newProxySender(cfg *Config) (*proxySender, error) {
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

func (s *proxySender) Send(b collector.MessageBuilder) {
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

func (s *proxySender) postMessage(m *collector.Message) error {
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
