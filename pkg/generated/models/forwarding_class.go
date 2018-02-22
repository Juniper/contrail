package models


// MakeForwardingClass makes ForwardingClass
func MakeForwardingClass() *ForwardingClass{
    return &ForwardingClass{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ForwardingClassDSCP: 0,
        ForwardingClassVlanPriority: 0,
        ForwardingClassMPLSExp: 0,
        ForwardingClassID: 0,
        
    }
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
    return []*ForwardingClass{}
}


