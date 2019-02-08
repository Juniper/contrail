package models

// SetPrefixes set prefixes.
func (i *InterfaceRouteTable) SetPrefixes(prefixes []string) {
	i.InterfaceRouteTableRoutes.Route = []*RouteType{}
	for _, prefix := range prefixes {
		i.InterfaceRouteTableRoutes.Route = append(i.InterfaceRouteTableRoutes.Route, &RouteType{
			Prefix: prefix,
		})
	}
}
