package models

// FirewallServiceGroupType

// FirewallServiceGroupType
//proteus:generate
type FirewallServiceGroupType struct {
	FirewallService []*FirewallServiceType `json:"firewall_service,omitempty"`
}

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
func MakeFirewallServiceGroupType() *FirewallServiceGroupType {
	return &FirewallServiceGroupType{
		//TODO(nati): Apply default

		FirewallService: MakeFirewallServiceTypeSlice(),
	}
}

// MakeFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
	return []*FirewallServiceGroupType{}
}
