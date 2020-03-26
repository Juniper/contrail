package testutil

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/httputil"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewTestHTTPServer starts and returns new test HTTP 2 Server.
func NewTestHTTPServer(h http.Handler) *httptest.Server {
	s := httptest.NewUnstartedServer(h)
	s.TLS = new(tls.Config)
	s.TLS.NextProtos = append(s.TLS.NextProtos, "h2")
	s.StartTLS()
	return s
}

// Flows is a definiton of possible flows in a mock HTTP server.
// It implements http.Handler which allows using it in httptest.NewServer.
type Flows []Flow

// PopResponse looks for a flow that matches the request and returns the associated
// response object. Then the flow is removed from the array.
func (f *Flows) PopResponse(r *http.Request) (Response, bool) {
	if f == nil {
		return Response{}, false
	}
	i := f.findMatchingIndex(r)
	if i == -1 {
		return Response{}, false
	}
	flow := (*f)[i]
	*f = append((*f)[:i], (*f)[i+1:]...)
	return flow.Response, true
}

func (f Flows) findMatchingIndex(r *http.Request) int {
	for i, flow := range f {
		if flow.Match(r) {
			return i
		}
	}
	return -1
}

// ServeHTTP makes Flows implement http.Handler. It looks for a matching flow
// and serves its response.
func (f *Flows) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		return
	}
	defer r.Body.Close()

	if resp, ok := f.PopResponse(r); ok {
		resp.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Do allows executing the handler without HTTP server.
func (f *Flows) Do(r *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	f.ServeHTTP(recorder, r)
	return recorder.Result(), nil
}

// VerifyExpectations checks if all flows were matched and removed from the array.
func (f Flows) VerifyExpectations() error {
	if len(f) > 0 {
		return errors.Errorf("not all flowes were matched, remaining: %v", f)
	}
	return nil
}

// Flow is an object that describes single request-response interaction with mock
// HTTP server.
type Flow struct {
	RequestMatcher
	Response
}

// RequestMatcher is a predicate for a http.Request that checks it's attributes.
type RequestMatcher struct {
	Path    RegexMatcher
	Method  StringMatcher
	Payload PayloadMatcher
}

// PathMatcher returns RequestMatcher that tries to match against given path value.
func PathMatcher(path string) RequestMatcher {
	return RequestMatcher{Path: RegexMatcher(path)}
}

// Strings returns a string representation of a RequestMatcher in a form of a logical expression.
func (m RequestMatcher) String() string {
	var seg []string
	if m.Path != "" {
		seg = append(seg, fmt.Sprintf("path ^= %q", m.Path))
	}
	if m.Method != "" {
		seg = append(seg, fmt.Sprintf("method == %q", m.Method))
	}
	if m.Payload != nil {
		seg = append(seg, fmt.Sprintf("payload == %v", format.MustJSON(m.Payload)))
	}

	return fmt.Sprintf("(%s)", strings.Join(seg, " && "))
}

// Match checks if given request matches requirements specified in a RequestMatcher.
func (m *RequestMatcher) Match(req *http.Request) bool {
	if m == nil {
		return true
	}
	if req == nil {
		return false
	}
	data, err := httputil.BufferBody(req)
	if err != nil {
		logrus.WithField("req", req).WithError(err).Debug("failed to read request body")
		return false
	}
	return m.Path.Match(req.URL.Path) && m.Method.Match(req.Method) && m.Payload.Match(data)
}

// RegexMatcher is a matcher that checks if a given string matches the regex.
type RegexMatcher string

// Match checks if given string matches the regex.
func (m RegexMatcher) Match(s string) bool {
	r := regexp.MustCompile(string(m))
	return r.MatchString(s)
}

// StringMatcher is a matcher that compares a given string to an exact string value.
type StringMatcher string

// Match checks if the given string is equal to the matcher's value.
func (m StringMatcher) Match(s string) bool {
	return m == "" || string(m) == s
}

// Arr is a shorthand for an JSON array in payload.
type Arr = []interface{}

// Obj is a shorthand for an JSON object in payload.
type Obj = map[string]interface{}

// PayloadMatcher is a predicate that checks if JSON data contains fields specified
// as matcher's keys and values.
type PayloadMatcher Obj

// FieldMatcher creates a PayloadMatcher that matches payloads that contain given value
// under given path.
// Path is a string that contains dot separated list of keys.
func FieldMatcher(path string, value interface{}) PayloadMatcher {
	return PayloadMatcher{}.WithField(path, value)
}

// Match compares PayloadMatcher contents with data in form of JSON bytes.
func (m PayloadMatcher) Match(data []byte) bool {
	return IsObjectSubsetOf((map[string]interface{})(m), format.MustReadJSON(data)) == nil
}

// SetField sets value in PayloadMatcher under given path.
func (m PayloadMatcher) SetField(path string, value interface{}) {
	segments := strings.Split(path, ".")
	if len(segments) == 0 {
		return
	}
	data := m
	last := len(segments) - 1
	for _, field := range segments[:last] {
		next, ok := data[field].(Obj)
		if !ok || next == nil {
			next = Obj{}
			data[field] = next
		}
		data = next
	}
	data[segments[last]] = value
}

// WithField returns a PayloadMatcher with an additional field. This method allows creating
// a matcher by chaning WithField calls.
// Note that that this method mutates the original matcher.
func (m PayloadMatcher) WithField(path string, value interface{}) PayloadMatcher {
	m.SetField(path, value)
	return m
}

// Response is a description of a response of a mock HTTP server.
type Response struct {
	StatusCode int
	Header     http.Header
	Payload    string
}

// ResponseOK returns response with http.StatusOK.
func ResponseOK(payload string) Response {
	return Response{StatusCode: http.StatusOK, Payload: payload}
}

// ResponseCreated returns response with http.StatusCreated.
func ResponseCreated(payload string) Response {
	return Response{StatusCode: http.StatusCreated, Payload: payload}
}

// ResponseNotFound returns response with http.StatusNotFound.
func ResponseNotFound(payload string) Response {
	return Response{StatusCode: http.StatusNotFound, Payload: payload}
}

// ServeHTTP serves writes response data to given http.ResponseWriter.
func (r *Response) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if r == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	code := r.StatusCode
	if code == 0 {
		code = http.StatusOK
	}
	httputil.CopyHeader(r.Header, w.Header())

	w.WriteHeader(code)
	if r.Payload != "" {
		io.WriteString(w, r.Payload)
	}
}
