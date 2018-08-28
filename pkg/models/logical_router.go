package models

const (
	NoneString = "None"
)

// GetVXLanIDInLogicaRouter returns vxlan network identifier property
func (lr *LogicalRouter) GetVXLanIDInLogicaRouter() string {
	id := lr.GetVxlanNetworkIdentifier()
	if id == NoneString {
		return ""
	}

	return id
}
