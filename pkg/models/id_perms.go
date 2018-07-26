package models

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// InitIDPerms initializes resource data when not provided.
func InitIDPerms(idPerms *IdPermsType, uuid string) *IdPermsType {
	if idPerms == nil {
		idPerms = &IdPermsType{
			Enable: true,
		}
	}

	if idPerms.UUID == nil {
		idPerms.UUID = initIDPermsUUID(uuid)
	}

	return idPerms
}

func initIDPermsUUID(uuid string) *UuidType {
	uuid = strings.Replace(uuid, "-", "", 4)
	uuidHigh, err1 := strconv.ParseInt(uuid[:len(uuid)/2], 16, 64)
	uuidLow, err2 := strconv.ParseInt(uuid[len(uuid)/2:], 16, 64)

	if len(uuid) != 32 || err1 != nil || err2 != nil {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		uuidHigh = random.Int63()
		uuidLow = random.Int63()
	}

	return &UuidType{
		UUIDMslong: uuidHigh,
		UUIDLslong: uuidLow,
	}
}
