package models

func (fip *FloatingIP) IsParentTypeInstanceIp() bool {
	return fip.GetParentType() == "instance-ip"
}
