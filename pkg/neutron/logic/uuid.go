package logic

import "strings"

func contrailUUIDToNeutronID(uuid string) string {
	return strings.Replace(uuid, "-", "", -1)
}
