package models

import "strconv"

func (st *SubnetType) GetStringRepresentation() string {
	return st.GetIPPrefix() + "/" + strconv.FormatInt(st.GetIPPrefixLen(), 10)
}
