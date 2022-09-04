package repository

import (
	"auth-api/core/entity"
	"context"
)

type AuthRepository interface {
	AuthenticateCredential(ctx context.Context, cred *entity.Credential) (*entity.UserDetail, error)
	GenerateToken(ctx context.Context, username, role string) (string, error)
	ValidateToken(ctx context.Context, token string) (*entity.CredentialClaim, error)
}
