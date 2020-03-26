package httputil

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// BufferBody reads the request body and buffers it as bytes.Reader.
func BufferBody(req *http.Request) ([]byte, error) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(data))
	req.ContentLength = int64(len(data))
	return data, nil
}
