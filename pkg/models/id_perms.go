package models

// InitIDPerms initializes resource data when not provided.
func InitIDPerms(idPerms *IdPermsType, uuid string) *IdPermsType {
	if idPerms == nil {
		idPerms = &IdPermsType{
			Enable: true,
		}
	}

	if idPerms.UUID == nil {
		idPerms.UUID = NewUUIDType(uuid)
	}

	return idPerms
}
