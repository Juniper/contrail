package keystone

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/databus23/keystone"
)

// KeystoneClient represents a client.
// nolint
type KeystoneClient struct {
	AuthURL    string `yaml:"authurl"`
	httpClient *http.Client
	InSecure   bool `yaml:"insecure"`
}

// NewKeystoneClient makes keystone client.
func NewKeystoneClient(authURL string, insecure bool) *KeystoneClient {
	c := &KeystoneClient{
		AuthURL:  authURL,
		InSecure: insecure,
	}
	c.Init()
	return c
}

// Init is used to initialize a keystone client.
func (k *KeystoneClient) Init() {
	tr := &http.Transport{
		Dial:            (&net.Dialer{}).Dial,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: k.InSecure},
	}
	client := &http.Client{
		Transport: tr,
	}
	k.httpClient = client
}

// SetAuthURL uses specified auth url in the keystone auth.
func (k *KeystoneClient) SetAuthURL(authURL string) {
	k.AuthURL = authURL
}

// NewAuth creates new keystone auth
func (k *KeystoneClient) NewAuth() *keystone.Auth {
	auth := keystone.New(k.AuthURL)
	auth.Client = k.httpClient
	return auth
}
