package undercloud

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// Data is the representation of cloudManager details.
type Data struct {
	cloudManagerInfo *models.RhospdCloudManager
}

// nolint: gocyclo
func (d *Data) updateUndercloudDetails(undercloudID string, u *UnderCloud) error {
	request := new(services.GetRhospdCloudManagerRequest)
	request.ID = undercloudID

	resp, err := u.APIServer.GetRhospdCloudManager(context.Background(), request)
	if err != nil {
		return err
	}
	d.cloudManagerInfo = resp.GetRhospdCloudManager()
	return nil
}
