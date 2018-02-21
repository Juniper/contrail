package models


// MakeInterfaceMirrorType makes InterfaceMirrorType
func MakeInterfaceMirrorType() *InterfaceMirrorType{
    return &InterfaceMirrorType{
    //TODO(nati): Apply default
    TrafficDirection: "",
        MirrorTo: MakeMirrorActionType(),
        
    }
}

// MakeInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
func MakeInterfaceMirrorTypeSlice() []*InterfaceMirrorType {
    return []*InterfaceMirrorType{}
}


