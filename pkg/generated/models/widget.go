package models


// MakeWidget makes Widget
func MakeWidget() *Widget{
    return &Widget{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ContainerConfig: "",
        ContentConfig: "",
        LayoutConfig: "",
        
    }
}

// MakeWidgetSlice() makes a slice of Widget
func MakeWidgetSlice() []*Widget {
    return []*Widget{}
}


