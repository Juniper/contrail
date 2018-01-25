package models


import ("fmt"
        "testing")
func TestNetworkDeviceConfig(t *testing.T) {
    model := MakeNetworkDeviceConfig()
    fmt.Println(model)
}
