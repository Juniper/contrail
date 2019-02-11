package logic

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/twinj/uuid"
)

func vncUUIDToNeutronID(uuid string) string {
	return strings.Replace(uuid, "-", "", -1)
}

func neutronIDToVncUUID(id string) (string, error) {
	if id == "" {
		return "", nil
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return "", errors.Wrap(err, "failed to translate neutron id to contrail uuid")
	}
	return uuid.String(), nil
}
