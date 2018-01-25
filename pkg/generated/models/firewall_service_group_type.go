package models
// FirewallServiceGroupType



import "encoding/json"

// FirewallServiceGroupType 
//proteus:generate
type FirewallServiceGroupType struct {

    FirewallService []*FirewallServiceType `json:"firewall_service,omitempty"`


}



// String returns json representation of the object
func (model *FirewallServiceGroupType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
func MakeFirewallServiceGroupType() *FirewallServiceGroupType{
    return &FirewallServiceGroupType{
    //TODO(nati): Apply default
    
            
                FirewallService:  MakeFirewallServiceTypeSlice(),
            
        
    }
}



// MakeFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
    return []*FirewallServiceGroupType{}
}
