package intent_test

import "github.com/Juniper/contrail/pkg/models"

type TestIntent struct {
	BaseIntent
	*models.VirtualNetwork
}
