package models

const (
	noneString = "None"
)

// GetVXLanIDInLogicaRouter returns vxlan network identifier property
func (lr *LogicalRouter) GetVXLanIDInLogicaRouter() string {
	id := lr.GetVxlanNetworkIdentifier()
	if id == noneString {
		return ""
	}

	return id
}
