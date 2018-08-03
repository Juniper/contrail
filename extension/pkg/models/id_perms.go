package models

// TODO(nati) remove this.

// NewIDPerms creates new UUIdType instance
func NewIDPerms(uuid string) *IdPermsType {
	return &IdPermsType{
		UUID:   NewUUIDType(uuid),
		Enable: true,
	}
}
