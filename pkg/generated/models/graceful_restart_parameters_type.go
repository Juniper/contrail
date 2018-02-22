package models


// MakeGracefulRestartParametersType makes GracefulRestartParametersType
func MakeGracefulRestartParametersType() *GracefulRestartParametersType{
    return &GracefulRestartParametersType{
    //TODO(nati): Apply default
    Enable: false,
        EndOfRibTimeout: 0,
        BGPHelperEnable: false,
        XMPPHelperEnable: false,
        RestartTime: 0,
        LongLivedRestartTime: 0,
        
    }
}

// MakeGracefulRestartParametersTypeSlice() makes a slice of GracefulRestartParametersType
func MakeGracefulRestartParametersTypeSlice() []*GracefulRestartParametersType {
    return []*GracefulRestartParametersType{}
}


