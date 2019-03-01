package logic

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/twinj/uuid"
)

// VncUUIDToNeutronID translates contrail vnc uuids into neutron id
func VncUUIDToNeutronID(uuid string) string {
	return strings.Replace(uuid, "-", "", -1)
}

func neutronIDToVncUUID(id string) (string, error) {
	if id == "" {
		return "", nil
	}

	u, err := uuid.Parse(id)
	if err != nil {
		return "", errors.Wrapf(err, "failed to translate neutron id: %v to contrail uuid", id)
	}
	return u.String(), nil
}
