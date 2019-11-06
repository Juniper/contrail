package proxy

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// TODO(dfurman): move rest of dynamicProxyMiddleware here

// HandleRequest proxies request from given context to first available target URL and
// returns response to client.
func HandleRequest(ctx echo.Context, targetURLs []string, log *logrus.Entry) error {
	rb := &responseBuffer{}
	b, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return errors.Wrap(err, "reading request body")
	}

	for i, targetURL := range targetURLs {
		rp, err := NewReverseProxy([]string{targetURL})
		if err != nil {
			return errors.Wrap(err, "new reverse proxy")
		}

		fmt.Printf("hoge before rp.ServeHTTP()\n"+
			"ctx: %#+v\n request: %#+v\n request URL: %#+v\n, response: %#+v\n, rb: %#+v\n",
			ctx, ctx.Request(), ctx.Request().URL, ctx.Response(), rb,
		)

		rb = newResponseBuffer(ctx.Response())
		req := ctx.Request().Clone(ctx.Request().Context())
		req.Body = ioutil.NopCloser(bytes.NewReader(b))
		req.ContentLength = int64(len(b))
		rp.ServeHTTP(rb, req)
		fmt.Printf("hoge after  rp.ServeHTTP()\n"+
			"ctx: %#+v\n request: %#+v\n request URL: %#+v\n, response: %#+v\n, rb: %#+v\n",
			ctx, ctx.Request(), ctx.Request().URL, ctx.Response(), rb,
		)

		if rb.Status() != http.StatusBadGateway && rb.Status() != http.StatusServiceUnavailable {
			fmt.Printf("hoge break with status %+v\n", rb.Status())
			break
		}

		if i < (len(targetURLs) - 1) {
			log.WithFields(logrus.Fields{
				"last-response-status": rb.Status(),
				"last-target-url":      targetURL,
			}).Debug("Target server unavailable - retrying request to next target")
		} else {
			log.WithFields(logrus.Fields{
				"last-response-status": rb.Status(),
				"last-target-url":      targetURL,
				"target-urls":          targetURLs,
			}).Info("All target servers unavailable")
		}
	}

	fmt.Printf("hoge before  rb.FlushToRW()\n"+
		"ctx: %#+v\n request: %#+v\n request URL: %#+v\n, response: %#+v\n, rb: %#+v\n",
		ctx, ctx.Request(), ctx.Request().URL, ctx.Response(), rb,
	)
	if err = rb.FlushToRW(); err != nil {
		return errors.Wrap(err, "flush response buffer to response writer")
	}
	fmt.Printf("hoge after  rb.FlushToRW()\n"+
		"ctx: %#+v\n request: %#+v\n request URL: %#+v\n, response: %#+v\n, rb: %#+v\n",
		ctx, ctx.Request(), ctx.Request().URL, ctx.Response(), rb,
	)

	return nil
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
	fmt.Printf("hoge rb.Header() rb.header: %#+v\n", rb.header)
	return rb.header
}

// Write writes given data to the buffer.
func (rb *responseBuffer) Write(data []byte) (int, error) {
	fmt.Printf("hoge rb.Write() string(data): %#+v\n", string(data))
	return rb.data.Write(data)
}

// WriteHeader sets status code field with given status code.
func (rb *responseBuffer) WriteHeader(statusCode int) {
	fmt.Printf("hoge rb.WriteHeader() statusCode: %#+v\n", statusCode)
	rb.statusCode = statusCode
}

// Hijack calls hijacks connection of wrapped response writer.
func (rb *responseBuffer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	fmt.Printf("hoge rb.Hijack()\n")
	return rb.rw.Hijack()
}

// Status returns status code written to buffer
func (rb *responseBuffer) Status() int {
	return rb.statusCode
}

// FlushToRW writes the header and data to wrapped response writer.
// It is intentionally named different from Flush() method (see http.Flusher())
// to prevent premature buffer flushing triggered by other actors.
func (rb *responseBuffer) FlushToRW() error {
	fmt.Printf("hoge rb.FlushToRW() before copy rb.rw.Header(): %#+v\n rb.header: %#+v \n", rb.rw.Header(), rb.header)
	copyHeader(rb.rw.Header(), rb.header)
	fmt.Printf("hoge rb.FlushToRW() after copy rb.rw.Header(): %#+v\n rb.header: %#+v \n", rb.rw.Header(), rb.header)
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
