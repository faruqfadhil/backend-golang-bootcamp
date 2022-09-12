package module

import (
	"book-api/core/entity"
	"book-api/core/repository"
	"context"
)

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	Username string
	Role     string
}

type AuthService interface {
	Login(ctx context.Context, cred *entity.Credential) (*LoginResponse, error)
	ValidateAccessToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error)
	ValidateRefreshToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest) (string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(ctx context.Context, cred *entity.Credential) (*LoginResponse, error) {
	// Check username dan password
	// If match, then claim token & refresh token.
	// Kalau nggak, maka gagal login.

	authenticated, err := s.repo.AuthenticateCredential(ctx, cred)
	if err != nil {
		return nil, err
	}

	// Generate token & refresh token.
	generatedToken, err := s.repo.GenerateToken(ctx, &repository.GenerateTokenRequest{
		Role:     authenticated.Role,
		Username: authenticated.UserName,
	})
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  generatedToken.AccessToken,
		RefreshToken: generatedToken.RefreshToken,
	}, nil
}

func (s *authService) ValidateAccessToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error) {
	// langsung validate ke jwt parse.
	return s.repo.ValidateAccessToken(ctx, token)
}

func (s *authService) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (string, error) {
	// simply generate new token.
	return s.repo.GenerateAccessToken(ctx, &repository.GenerateTokenRequest{
		Username: req.Username,
		Role:     req.Role,
	})
}

func (s *authService) ValidateRefreshToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error) {
	return s.repo.ValidateRefreshToken(ctx, token)
}
