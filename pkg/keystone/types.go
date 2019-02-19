package keystone

import "time"

//AuthRequest is used to request an authentication.
type AuthRequest struct {
	Auth *Auth `json:"auth"`
}

//Auth is used to request an authentication.
type Auth struct {
	Identity *Identity `json:"identity"`
	Scope    *Scope    `json:"scope"`
}

//Scope is used to limit scope of auth request.
type Scope struct {
	Domain  *Domain  `json:"domain,omitempty"`
	Project *Project `json:"project,omitempty"`
}

//Domain represents domain object.
type Domain struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//Project represents project object.
type Project struct {
	Domain *Domain `json:"domain,omitempty"`
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
}

//Identity represents a auth methods.
type Identity struct {
	Methods  []string   `json:"methods"`
	Password *Password  `json:"password,omitempty"`
	Token    *UserToken `json:"token,omitempty"`
	Cluster  *Cluster   `json:"cluster,omitempty"`
}

//Password represents a password.
type Password struct {
	User *User `json:"user,omitempty"`
}

//AuthResponse represents a authentication response.
type AuthResponse struct {
	Token *Token `json:"token"`
}

//User reprenetns a user.
type User struct {
	Domain   *Domain `json:"domain,omitempty"`
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Roles    []*Role `json:"roles"`
}

//Role represents a user role.
type Role struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Project *Project `json:"project"`
}

//Cluster represent a cluster object sent by user
//to get new token using cluster token
type UserToken struct {
	ID    string     `json:"id"`
	Token *UserToken `json:"token,omitempty"`
}

//UserToken represent a token object sent by user to get new token
type UserToken struct {
	ID string `json:"id"`
	Token
}

//Token represents a token object.
type Token struct {
	AuditIds  []string   `json:"audit_ids"`
	Catalog   []*Catalog `json:"catalog"`
	Domain    *Domain    `json:"domain"`
	Project   *Project   `json:"project"`
	User      *User      `json:"user"`
	ExpiresAt time.Time  `json:"expires_at"`
	IssuedAt  time.Time  `json:"issued_at"`
	Methods   []string   `json:"methods"`
	Roles     []*Role    `json:"roles"`
}

//Catalog represents API catalog.
type Catalog struct {
	Endpoints []*Endpoint `json:"endpoints"`
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
}

//Endpoint represents API endpoint.
type Endpoint struct {
	ID        string `json:"id"`
	Interface string `json:"interface"`
	Region    string `json:"region"`
	URL       string `json:"url"`
}

//ValidateTokenResponse represents a response object for validate token request.
type ValidateTokenResponse struct {
	Token *Token `json:"token"`
}
