package apiserver

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestServer_URL(t *testing.T) {
	tests := []struct {
		address    string
		wantPrefix string
	}{
		{wantPrefix: "http://[::]:"},
		{address: ":9091", wantPrefix: "http://[::]:9091"},
		{address: "localhost:9092", wantPrefix: "http://127.0.0.1:9092"},
		{address: "127.0.0.1:9093", wantPrefix: "http://127.0.0.1:9093"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("address:%q", tt.address), func(t *testing.T) {
			viper.Set("server.address", tt.address)
			s, err := NewServer(nil, nil)
			assert.NoError(t, err)
			defer s.Close()

			go s.Run()
			<-s.ReadyChan()
			if got := s.URL(); !strings.HasPrefix(got, tt.wantPrefix) {
				t.Errorf("Server.URL() = %v, want %v", got, tt.wantPrefix)
			}
		})
	}
}
