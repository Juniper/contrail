package models

import "strings"

// SetDNSNameservers sets DNS Nameservers.
func (m *IpamSubnetType) SetDNSNameservers(nameservers []string) {
	if len(nameservers) != 0 {
		m.DHCPOptionList = &DhcpOptionsListType{
			DHCPOption: []*DhcpOptionType{
				{
					DHCPOptionName:  "6",
					DHCPOptionValue: strings.Join(nameservers, " "),
				},
			},
		}
	} else {
		m.DHCPOptionList = &DhcpOptionsListType{}
	}
}
