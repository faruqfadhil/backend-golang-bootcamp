package entity

type Credential struct {
	Username string
	Password string
}

type UserDetail struct {
	UserName string
	Password string
	Role     string
}

type AccessTokenCredentialClaim struct {
	Username string
	Role     string
}
