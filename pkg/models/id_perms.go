package models

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// InitIDPerms initializes resource data when not provided
func InitIDPerms(idPerms *IdPermsType, uuid string) IdPermsType {

	var result IdPermsType
	if idPerms == nil {
		result = IdPermsType{
			Enable: true,
		}
	} else {
		result = *idPerms
	}

	if result.UUID == nil {
		uuid = strings.Replace(uuid, "-", "", 4)
		uuidHigh, err1 := strconv.ParseInt(uuid[:len(uuid)/2], 16, 64)
		uuidLow, err2 := strconv.ParseInt(uuid[len(uuid)/2:], 16, 64)
		if len(uuid) != 32 || err1 != nil || err2 != nil {
			random := rand.New(rand.NewSource(time.Now().UnixNano()))
			uuidHigh = random.Int63()
			uuidLow = random.Int63()
		}

		result.UUID = &UuidType{
			UUIDMslong: uuidHigh,
			UUIDLslong: uuidLow,
		}
	}

	return result
}
