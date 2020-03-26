package testutil_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/httputil"
	. "github.com/Juniper/asf/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestRegexMatcher_Match(t *testing.T) {
	tests := []struct {
		name string
		m    RegexMatcher
		url  string
		want bool
	}{{
		name: "empty", want: true,
	}, {
		name: "empty pattern matches any", url: "fo/ob/ar", want: true,
	}, {
		name: "empty url", m: "/create-config-object", want: false,
	}, {
		name: "optional plural", url: "/create-config-object", m: "/create-config-objects?", want: true,
	}, {
		name: "optional plural", url: "/create-config-objects", m: "/create-config-objects?", want: true,
	}, {
		name: "substring", url: "/create-config-objects", m: "config-obje", want: true,
	}, {
		name: "different values", url: "foo", m: "bar", want: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Match(tt.url); got != tt.want {
				t.Errorf("RegexMatcher(%q).Match(%q) = %v, want %v", tt.m, tt.url, got, tt.want)
			}
		})
	}
}

func TestPayloadMatcher_Match(t *testing.T) {
	tests := []struct {
		name string
		m    PayloadMatcher
		data []byte
		want bool
	}{{
		name: "nil", want: true,
	}, {
		name: "nil matcher", data: []byte(`{"foo": ["bar"]}`), want: true,
	}, {
		name: "nil value", m: PayloadMatcher{"foo": 1}, want: false,
	}, {
		name: "empty matcher", m: PayloadMatcher{}, data: []byte(`{"foo": ["bar"]}`), want: true,
	}, {
		name: "match value", m: PayloadMatcher{"foo": Arr{"bar"}}, data: []byte(`{"foo": ["bar"]}`), want: true,
	}, {
		name: "match nested key",
		m:    PayloadMatcher{"foo": Obj{"bar": Obj{"x": 1}}},
		data: []byte(`{"foo": {"bar": {"x": 1}}}`),
		want: true,
	}, {
		name: "match one of keys", m: PayloadMatcher{"foo": 1}, data: []byte(`{"foo": 1, "bar": 2}`), want: true,
	}, {
		name: "bad value", m: PayloadMatcher{"foo": 1}, data: []byte(`{"foo": 2}`), want: false,
	}, {
		name: "notnull value", m: PayloadMatcher{"foo": "$notnull"}, data: []byte(`{"foo": 2}`), want: true,
	}, {
		name: "notnull value", m: PayloadMatcher{"foo": "$notnull"}, data: []byte(`{}`), want: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Match(tt.data); got != tt.want {
				t.Errorf("PayloadMatcher(%v).Match(%s) = %v, want %v", tt.m, tt.data, got, tt.want)
			}
		})
	}
}

func newTestRequestFactory(t *testing.T) func(string, string, string, ...http.Header) *http.Request {
	return func(method, url, payload string, headers ...http.Header) *http.Request {
		r := httptest.NewRequest(method, url, strings.NewReader(payload))
		for _, h := range headers {
			httputil.CopyHeader(h, r.Header)
		}
		return r
	}
}

func TestRequestMatcher_Match(t *testing.T) {
	newRequest := newTestRequestFactory(t)
	tests := []struct {
		name    string
		matcher *RequestMatcher
		req     *http.Request
		want    bool
	}{{
		name: "nil", want: true,
	}, {
		name: "empty matcher", matcher: &RequestMatcher{}, req: newRequest(http.MethodGet, "/foo", "{}"), want: true,
	}, {
		name:    "complete matcher",
		matcher: &RequestMatcher{Method: http.MethodGet, Path: "foo", Payload: PayloadMatcher{"bar": 1}},
		req:     newRequest(http.MethodGet, "/foo", `{"bar": 1}`),
		want:    true,
	}, {
		name:    "nil request",
		matcher: &RequestMatcher{Method: http.MethodGet, Path: "foo", Payload: PayloadMatcher{"bar": 1}},
		want:    false,
	}, {
		name:    "bad method",
		matcher: &RequestMatcher{Method: http.MethodGet, Path: "foo", Payload: PayloadMatcher{"bar": 1}},
		req:     newRequest(http.MethodPost, "/foo", `{"bar": 1}`),
		want:    false,
	}, {
		name:    "bad path",
		matcher: &RequestMatcher{Method: http.MethodGet, Path: "foo", Payload: PayloadMatcher{"bar": 1}},
		req:     newRequest(http.MethodGet, "/bar", `{"bar": 1}`),
		want:    false,
	}, {
		name:    "bad payload",
		matcher: &RequestMatcher{Method: http.MethodGet, Path: "foo", Payload: PayloadMatcher{"bar": 1}},
		req:     newRequest(http.MethodGet, "/foo", `{"bar": 2}`),
		want:    false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matcher.Match(tt.req); got != tt.want {
				t.Errorf("RequestMatcher.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestMatcher_String(t *testing.T) {
	tests := []struct {
		matcher RequestMatcher
		want    string
	}{{
		matcher: RequestMatcher{}, want: "()",
	}, {
		matcher: RequestMatcher{Method: http.MethodGet, Path: "/foo", Payload: PayloadMatcher{"foo": int(1)}},
		want:    `(path ^= "/foo" && method == "GET" && payload == {"foo":1})`,
	}}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.matcher.String(); got != tt.want {
				t.Errorf("RequestMatcher.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func assertReaderEqual(t *testing.T, got, want io.Reader) {
	var dataA, dataB []byte
	var err error
	if got != nil {
		dataA, err = ioutil.ReadAll(got)
		assert.NoError(t, err)
	}
	if want != nil {
		dataB, err = ioutil.ReadAll(want)
		assert.NoError(t, err)
	}
	assert.Equal(t, dataA, dataB)
}

func TestResponse_ServeHTTP(t *testing.T) {
	nop := ioutil.NopCloser
	tests := []struct {
		name     string
		obj      Response
		expected http.Response
	}{{
		name:     "empty",
		expected: http.Response{StatusCode: http.StatusOK, Header: http.Header{}, Body: nop(&bytes.Reader{})},
	}, {
		name: "complete response",
		obj: Response{
			StatusCode: http.StatusCreated,
			Header:     http.Header{"Foo": []string{"bar"}},
			Payload:    `{"foo":"bar"}`,
		},
		expected: http.Response{
			StatusCode: http.StatusCreated,
			Header:     http.Header{"Foo": []string{"bar"}},
			Body:       nop(strings.NewReader(`{"foo":"bar"}`)),
		},
	}}

	for _, tt := range tests {
		r := httptest.NewRecorder()
		tt.obj.ServeHTTP(r, nil)
		result := r.Result()
		assert.Equal(t, result.StatusCode, tt.expected.StatusCode)
		assert.Equal(t, result.Header, tt.expected.Header)
		assertReaderEqual(t, result.Body, tt.expected.Body)
	}
}

func TestFlows_VerifyExpectations(t *testing.T) {
	tests := []struct {
		name    string
		f       Flows
		wantErr bool
	}{{
		name: "nil",
	}, {
		name: "some entries", f: Flows{{}, {}}, wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.VerifyExpectations(); (err != nil) != tt.wantErr {
				t.Errorf("Flows.VerifyExpectations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlows_PopResponse(t *testing.T) {
	newRequest := newTestRequestFactory(t)
	tests := []struct {
		name          string
		f             Flows
		r             *http.Request
		want          Response
		wantOk        bool
		expectedFlows Flows
	}{{
		name: "nil",
	}, {
		name:          "nil request",
		f:             Flows{{}, {}},
		expectedFlows: Flows{{}, {}},
	}, {
		name:          "no expectation flows",
		f:             Flows{{Response: Response{StatusCode: http.StatusOK}}, {Response: Response{}}},
		r:             newRequest(http.MethodGet, "http://localhost/foo", ""),
		want:          Response{StatusCode: http.StatusOK},
		wantOk:        true,
		expectedFlows: Flows{{}},
	}, {
		name: "matches method",
		f: Flows{
			{RequestMatcher: RequestMatcher{Method: http.MethodPost}, Response: Response{StatusCode: http.StatusCreated}},
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusOK}},
		},
		r:      newRequest(http.MethodGet, "http://localhost/foo", ""),
		want:   Response{StatusCode: http.StatusOK},
		wantOk: true,
		expectedFlows: Flows{
			{RequestMatcher: RequestMatcher{Method: http.MethodPost}, Response: Response{StatusCode: http.StatusCreated}},
		},
	}, {
		name: "get first matching",
		f: Flows{
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusCreated}},
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusOK}},
		},
		r:      newRequest(http.MethodGet, "http://localhost/foo", ""),
		want:   Response{StatusCode: http.StatusCreated},
		wantOk: true,
		expectedFlows: Flows{
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusOK}},
		},
	}, {
		name: "no match",
		f: Flows{
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusCreated}},
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusOK}},
		},
		r: newRequest(http.MethodDelete, "http://localhost/foo", ""),
		expectedFlows: Flows{
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusCreated}},
			{RequestMatcher: RequestMatcher{Method: http.MethodGet}, Response: Response{StatusCode: http.StatusOK}},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := tt.f.PopResponse(tt.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flows.PopResponse() got = %v, want %v", got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Flows.PopResponse() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if !reflect.DeepEqual(tt.f, tt.expectedFlows) {
				t.Errorf("bad flows content after pop got %v, want %v", tt.f, tt.expectedFlows)
			}
		})
	}
}

func TestPayloadMatcher_WithField(t *testing.T) {
	tests := []struct {
		name string
		got  PayloadMatcher
		want PayloadMatcher
	}{{
		name: "nil",
	}, {
		name: "some fields",
		got:  PayloadMatcher{}.WithField("foo", 1).WithField("bar.foo", 2),
		want: PayloadMatcher{"foo": 1, "bar": Obj{"foo": 2}},
	}, {
		name: "overwrite field",
		got:  PayloadMatcher{}.WithField("foo", 1).WithField("foo", 2),
		want: PayloadMatcher{"foo": 2},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}

const (
	applicationJSONContentType = "application/json"
)

func TestFlowsMockServer(t *testing.T) {
	flows := Flows{{
		RequestMatcher: RequestMatcher{Method: http.MethodPost, Payload: FieldMatcher("foo", 1)},
		Response:       Response{StatusCode: http.StatusCreated, Payload: `{"status": "created"}`},
	}, {
		RequestMatcher: RequestMatcher{Path: "/get-config-object", Method: http.MethodGet},
		Response:       Response{Payload: `{"status":"ok"}`},
	}, {
		Response: Response{StatusCode: http.StatusTeapot},
	}}

	server := httptest.NewServer(&flows)

	t.Run("fresh server does not match expectations", func(t *testing.T) {
		assert.Error(t, flows.VerifyExpectations())
	})

	t.Run("simple get hits the last case", func(t *testing.T) {
		resp, err := http.Get(server.URL)
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusTeapot)
		assert.Error(t, flows.VerifyExpectations())
	})

	t.Run("non existing flow", func(t *testing.T) {
		resp, err := http.Get(server.URL)
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusNotImplemented)
		assert.Error(t, flows.VerifyExpectations())
	})

	t.Run("match path", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/get-config-object")
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assertReaderEqual(t, resp.Body, strings.NewReader(`{"status":"ok"}`))
		assert.Error(t, flows.VerifyExpectations())
	})

	t.Run("match method and payload", func(t *testing.T) {
		resp, err := http.Post(
			server.URL,
			applicationJSONContentType,
			strings.NewReader(`{"foo":1, "bar":2}`),
		)
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusCreated)
		assert.NoError(t, flows.VerifyExpectations())
	})
}
