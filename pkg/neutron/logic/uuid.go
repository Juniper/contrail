package logic

import "strings"

func contrailUUIDToNeutronID(contrailUUID string) string {
	return strings.Replace(contrailUUID, "-", "", -1)
}
