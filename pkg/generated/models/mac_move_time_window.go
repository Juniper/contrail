package models

// MACMoveTimeWindow

type MACMoveTimeWindow int

// MakeMACMoveTimeWindow makes MACMoveTimeWindow
func MakeMACMoveTimeWindow() MACMoveTimeWindow {
	var data MACMoveTimeWindow
	return data
}

// InterfaceToMACMoveTimeWindow makes MACMoveTimeWindow from interface
func InterfaceToMACMoveTimeWindow(data interface{}) MACMoveTimeWindow {
	return data.(MACMoveTimeWindow)
}

// InterfaceToMACMoveTimeWindowSlice makes a slice of MACMoveTimeWindow from interface
func InterfaceToMACMoveTimeWindowSlice(data interface{}) []MACMoveTimeWindow {
	list := data.([]interface{})
	result := MakeMACMoveTimeWindowSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACMoveTimeWindow(item))
	}
	return result
}

// MakeMACMoveTimeWindowSlice() makes a slice of MACMoveTimeWindow
func MakeMACMoveTimeWindowSlice() []MACMoveTimeWindow {
	return []MACMoveTimeWindow{}
}
