package models

// MACMoveTimeWindow

type MACMoveTimeWindow int

func MakeMACMoveTimeWindow() MACMoveTimeWindow {
	var data MACMoveTimeWindow
	return data
}

func InterfaceToMACMoveTimeWindow(data interface{}) MACMoveTimeWindow {
	return data.(MACMoveTimeWindow)
}

func InterfaceToMACMoveTimeWindowSlice(data interface{}) []MACMoveTimeWindow {
	list := data.([]interface{})
	result := MakeMACMoveTimeWindowSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACMoveTimeWindow(item))
	}
	return result
}

func MakeMACMoveTimeWindowSlice() []MACMoveTimeWindow {
	return []MACMoveTimeWindow{}
}
