package vncapi

import (
	"context"
	"net/url"

	"github.com/Juniper/contrail/pkg/common"
)

//VncProjectListResponse represents a project list response.
type VncProjectListResponse struct {
	Projects []*VncProject `json:"projects"`
}

//VncProject represents a vnc config project object.
type VncProject struct {
	Project *ConfigProject `json:"project"`
}

//Project represents project object.
type ConfigProject struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// VncApiClient represents a client.
type VncApiClient struct {
	*common.BaseHTTP
}

// NewVncApiClient makes vncapi client.
func NewVncApiClient() *VncApiClient {
	c := &VncApiClient{
		common.NewBaseHTTP("", true),
	}
	return c
}

// Initialize is used to initialize a vncapi client.
func (v *VncApiClient) Initialize(configURL string) {
	v.Endpoint = configURL
	v.Init()
}

// ListProjects reads all the projects from vnc config
func (v *VncApiClient) ListProjects() (*VncProjectListResponse, error) {
	projectURI := "/projects"
	query := url.Values{"detail": []string{"True"}}
	vncProjectsResponse := &VncProjectListResponse{}
	_, err := v.ReadWithQuery(context.Background(), projectURI, query, vncProjectsResponse)
	return vncProjectsResponse, err
}
