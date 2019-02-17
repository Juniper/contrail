package overcloud

const (
	createAction = "create"
	updateAction = "update"
	deleteAction = "delete"
)

type contrailCloudDeployer struct {
	deployOverCloud
}

func (c *contrailCloudDeployer) isUpdated() (bool, error) {
	return true, nil
}

func (c *contrailCloudDeployer) createOvercloud() error {
	return nil
}

func (c *contrailCloudDeployer) updateOvercloud() error {
	return nil
}

func (c *contrailCloudDeployer) deleteOvercloud() error {
	return nil
}

func (c *contrailCloudDeployer) Deploy() error {
	switch c.action {
	case createAction:
		err := c.createOvercloud()
		if err != nil {
			return err
		}
		return nil
	case updateAction:
		updated, err := c.isUpdated()
		if err != nil {
			return err
		}
		if updated {
			return nil
		}
		err = c.createOvercloud()
		if err != nil {
			return err
		}
		return nil
	case deleteAction:
		err := c.deleteOvercloud()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
