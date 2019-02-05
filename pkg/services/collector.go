package services

/* to use add to config.yml a section:
collector:
  url: http://collectorproxyhost:port/url
  queuesize: 32
*/

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const (
	defaultCollectorTimeout   = 30 * time.Second
	defaultCollectorQueueSize = 16
)

const (
	typeRestAPITrace = "RestApiTrace"
)

type collectorMessage struct {
	SandeshType string      `json:"sandesh_type"`
	Payload     interface{} `json:"payload"`
}

type collectorReastAPITrace struct {
	RequestID    string `json:"request_id"`
	URL          string `json:"url"`
	Method       string `json:"method"`
	RequestData  string `json:"request_data"`
	Status       string `json:"status"`
	ResponseBody string `json:"response_body"`
	RequestError string `json:"request_error"`
}

// Collector represent interface to Json2Sandesh proxy
type Collector struct {
	url        string
	timeout    time.Duration
	client     *http.Client
	queueSize  int
	queueCount int
	queueMutex sync.Mutex
}

// NewCollector makes a collector
func NewCollector(url string, queueSize int) *Collector {
	if url == "" {
		logrus.Info("collector.url not specified in configuration file")
	}
	if queueSize == 0 {
		queueSize = defaultCollectorQueueSize
	}
	return &Collector{
		url:       url,
		timeout:   defaultCollectorTimeout,
		queueSize: queueSize,
	}
}

// RestAPITrace sends message with type RestApiTrace
func (collector *Collector) RestAPITrace(c echo.Context, reqBody, resBody []byte) {
	if collector.url == "" {
		return
	}
	if c.Request().Method == "GET" {
		return
	}

	m := &collectorMessage{
		SandeshType: typeRestAPITrace,
		Payload: &collectorReastAPITrace{
			RequestID:    "req-" + uuid.NewV4().String(),
			URL:          c.Request().URL.String(),
			Method:       c.Request().Method,
			RequestData:  string(reqBody),
			Status:       strconv.Itoa(c.Response().Status) + " " + http.StatusText(c.Response().Status),
			ResponseBody: string(resBody),
			RequestError: "",
		},
	}

	data, err := json.Marshal(m)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Info("send message to collector failed")
		return
	}
	go collector.sendMessage(data)
}

func (collector *Collector) sendMessage(data []byte) {
	collector.queueMutex.Lock()
	if collector.queueCount > collector.queueSize {
		collector.queueMutex.Unlock()
		logrus.Info("collector message queue is full")
		return
	}
	collector.queueCount++
	collector.queueMutex.Unlock()
	defer func() {
		collector.queueMutex.Lock()
		collector.queueCount--
		collector.queueMutex.Unlock()
	}()
	if collector.client == nil {
		collector.client = &http.Client{
			Timeout: collector.timeout,
		}
	}
	req, err := http.NewRequest("POST", collector.url, bytes.NewReader(data))
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Info("send message to collector failed")
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	resp, err := collector.client.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Info("send message to collector failed")
		return
	}
	defer resp.Body.Close() // nolint: errcheck
	if resp.StatusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"status_code": resp.StatusCode}).Info("send message to collector failed")
	}
}
