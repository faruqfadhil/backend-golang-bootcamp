package entity

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialClaim struct {
	Username string
	Role     string
}

type UserDetail struct {
	Username string
	Password string
	Role     string
}
