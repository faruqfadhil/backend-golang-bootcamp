package repository

import (
	"book-api/core/entity"
	"context"
)

type GenerateTokenRequest struct {
	Role     string
	Username string
}

type GenerateTokenResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthRepository interface {
	AuthenticateCredential(ctx context.Context, cred *entity.Credential) (*entity.UserDetail, error)
	GenerateToken(ctx context.Context, req *GenerateTokenRequest) (*GenerateTokenResponse, error)
	ValidateAccessToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error)
	ValidateRefreshToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error)
	GenerateAccessToken(ctx context.Context, req *GenerateTokenRequest) (string, error)
}
