package models


// MakeUserCredentials makes UserCredentials
func MakeUserCredentials() *UserCredentials{
    return &UserCredentials{
    //TODO(nati): Apply default
    Username: "",
        Password: "",
        
    }
}

// MakeUserCredentialsSlice() makes a slice of UserCredentials
func MakeUserCredentialsSlice() []*UserCredentials {
    return []*UserCredentials{}
}


