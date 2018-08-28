package models

// IsVXLanIDInLogicaRouter returns vxlan network identifier property
func (lr *LogicalRouter) GetVXLanIDInLogicaRouter() string {
	id := lr.GetVxlanNetworkIdentifier()
	if id == "None" {
		lr.VxlanNetworkIdentifier = ""
		return ""
	}

	return id
}
