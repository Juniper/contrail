package models

// UserCredentials

// UserCredentials
//proteus:generate
type UserCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// MakeUserCredentials makes UserCredentials
func MakeUserCredentials() *UserCredentials {
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
