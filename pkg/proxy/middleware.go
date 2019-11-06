package proxy

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// TODO(dfurman): move rest of dynamicProxyMiddleware here

const (
	skipServerCertificateVerification = true // TODO: add "insecure" field to endpoint schema
)

// HandleRequest proxies request from given context to first available target URL and
// returns response to client.
func HandleRequest(ctx echo.Context, rawTargetURLs []string, log *logrus.Entry) error {
	rb := &responseBuffer{}
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return errors.Wrap(err, "reading request body")
	}

	for i, rawTargetURL := range rawTargetURLs {
		targetURL, pErr := url.Parse(rawTargetURL)
		if pErr != nil {
			logrus.WithError(pErr).WithField("target-url", targetURL).Error("Failed to parse target URL - ignoring")
			continue
		}

		rp := httputil.NewSingleHostReverseProxy(targetURL)
		rp.Transport = newTransport()
		rb = newResponseBuffer(ctx.Response())
		rp.ServeHTTP(rb, requestWithBody(ctx.Request(), body))
		if rb.Status() != http.StatusBadGateway && rb.Status() != http.StatusServiceUnavailable {
			break
		}

		e := log.WithFields(logrus.Fields{
			"last-response-status": rb.Status(),
			"last-target-url":      rawTargetURL,
			"target-urls":          rawTargetURLs,
		})
		if i < (len(rawTargetURLs) - 1) {
			e.Debug("Target server unavailable - retrying request to next target")
		} else {
			e.Info("All target servers unavailable")
		}
	}
	if err = rb.FlushToRW(); err != nil {
		return errors.Wrap(err, "flush response buffer to response writer")
	}

	return nil
}

func newTransport() *http.Transport {
	t := cleanhttp.DefaultPooledTransport()
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: skipServerCertificateVerification}
	return t
}

func requestWithBody(r *http.Request, body []byte) *http.Request {
	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	r.ContentLength = int64(len(body))
	return r
}

// responseBuffer wraps response writer and allows postponing writing response to it.
// Wrapped response writer needs to implement http.ResponseWriter and http.Hijacker interfaces.
type responseBuffer struct {
	rw         responseWriterHijacker
	statusCode int
	header     http.Header
	data       *bytes.Buffer
}

func newResponseBuffer(rw responseWriterHijacker) *responseBuffer {
	return &responseBuffer{
		rw:     rw,
		header: rw.Header().Clone(),
		data:   &bytes.Buffer{},
	}
}

type responseWriterHijacker interface {
	http.ResponseWriter
	http.Hijacker
}

// Header returns the header map of wrapped response writer that will be sent by WriteHeader.
func (rb *responseBuffer) Header() http.Header {
	return rb.header
}

// Write writes given data to the buffer.
func (rb *responseBuffer) Write(data []byte) (int, error) {
	return rb.data.Write(data)
}

// WriteHeader sets status code field with given status code.
func (rb *responseBuffer) WriteHeader(statusCode int) {
	rb.statusCode = statusCode
}

// Hijack calls hijacks connection of wrapped response writer.
func (rb *responseBuffer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return rb.rw.Hijack()
}

// Status returns status code written to buffer.
func (rb *responseBuffer) Status() int {
	return rb.statusCode
}

// FlushToRW writes the header and data to wrapped response writer.
// It is intentionally named different from Flush() method (see http.Flusher())
// to prevent premature buffer flushing triggered by other actors.
func (rb *responseBuffer) FlushToRW() error {
	copyHeader(rb.rw.Header(), rb.header)

	if rb.statusCode != 0 {
		rb.rw.WriteHeader(rb.statusCode)
	}

	_, err := rb.rw.Write(rb.data.Bytes())
	if err != nil {
		return errors.Wrap(err, "write the target's response to client")
	}
	return nil
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
