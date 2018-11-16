package undercloud

import (
	"context"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// Data is the representation of cloudManager details.
type Data struct {
	cloudManagerInfo *models.RhospdCloudManager
	client           *client.HTTP
}

// NewData creates a undercloud data
func NewData(apiClient *client.HTTP) *Data {
	return &Data{
		client: apiClient,
	}
}

func (d *Data) updateUndercloudDetails(undercloudID string) error {
	request := new(services.GetRhospdCloudManagerRequest)
	request.ID = undercloudID

	resp, err := d.client.GetRhospdCloudManager(context.Background(), request)
	if err != nil {
		return err
	}
	d.cloudManagerInfo = resp.GetRhospdCloudManager()
	return nil
}
