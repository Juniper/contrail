package httputil

import "net/http"

// CopyHeader copies header values from src to dst.
func CopyHeader(src, dst http.Header) {
	for k := range src {
		dst.Set(k, src.Get(k))
	}
}

// cloneHeader returns a copy of h or nil if h is nil.
// TODO(mblotniak): use http.Header.Clone() when Go version is updated to 1.13.
func cloneHeader(h http.Header) http.Header {
	if h == nil {
		return nil
	}

	// Find total number of values.
	nv := 0
	for _, vv := range h {
		nv += len(vv)
	}
	sv := make([]string, nv) // shared backing array for headers' values
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		n := copy(sv, vv)
		h2[k] = sv[:n:n]
		sv = sv[n:]
	}
	return h2
}
