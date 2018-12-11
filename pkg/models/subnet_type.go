package models

import "strconv"

// GetStringRepresentation converts SubnetType struct to a string of a full network address in format: "<ip>/<cidr>".
func (st *SubnetType) GetStringRepresentation() string {
	return st.GetIPPrefix() + "/" + strconv.FormatInt(st.GetIPPrefixLen(), 10)
}
