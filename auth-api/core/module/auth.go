package module

import (
	"auth-api/core/entity"
	"auth-api/core/repository"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, cred *entity.Credential) (string, error)
	ValidateToken(ctx context.Context, token string) (*entity.CredentialClaim, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(ctx context.Context, cred *entity.Credential) (string, error) {
	authenticated, err := s.repo.AuthenticateCredential(ctx, cred)
	if err != nil {
		return "", nil
	}

	// generate token
	generatedToken, err := s.repo.GenerateToken(ctx, authenticated.Username, authenticated.Role)
	if err != nil {
		return "", nil
	}

	return generatedToken, nil
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*entity.CredentialClaim, error) {
	return s.repo.ValidateToken(ctx, token)
}
