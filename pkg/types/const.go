package types

const (
	//PermsNone for no permission
	PermsNone = 0
	//PermsX for exec permission
	PermsX = 1
	//PermsW for write permission
	PermsW = 2
	//PermsR if read permission
	PermsR = 4
	//PermsWX for exec and write permission
	PermsWX = 3
	//PermsRX for exec and read permission
	PermsRX = 5
	//PermsRW for read and write permission
	PermsRW = 6
	//PermsRWX for all permission
	PermsRWX = 7
)
