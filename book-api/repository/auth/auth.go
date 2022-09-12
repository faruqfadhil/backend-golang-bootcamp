package auth

import (
	"book-api/core/entity"
	"book-api/core/repository"
	errlib "book-api/pkg/error"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var usersData = map[string]*entity.UserDetail{
	"user1": {
		UserName: "user1",
		Password: "password1",
		Role:     "student",
	},
	"user2": {
		UserName: "user2",
		Password: "password2",
		Role:     "admin",
	},
}

type authRepository struct {
	accessTokenSecretKey       string
	refreshTokenSecretKey      string
	accessTokenExpiryInMinutes time.Duration
	refreshTokenExpiryInHour   time.Duration
}

func NewAuthRepository(
	accessTokenSecretKey string,
	refreshTokenSecretKey string,
	accessTokenExpiryInMinutes time.Duration,
	refreshTokenExpiryInHour time.Duration) repository.AuthRepository {
	return &authRepository{
		accessTokenSecretKey:       accessTokenSecretKey,
		refreshTokenSecretKey:      refreshTokenSecretKey,
		accessTokenExpiryInMinutes: accessTokenExpiryInMinutes,
		refreshTokenExpiryInHour:   refreshTokenExpiryInHour,
	}
}

func (r *authRepository) AuthenticateCredential(ctx context.Context, cred *entity.Credential) (*entity.UserDetail, error) {
	if _, ok := usersData[cred.Username]; !ok {
		return nil, errlib.ErrUnauthenticated
	}

	userDetail := usersData[cred.Username]
	if userDetail.Password != cred.Password {
		return nil, errlib.ErrUnauthenticated
	}

	return userDetail, nil
}

type accessTokenJwtClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type refreshTokenJwtClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (r *authRepository) GenerateToken(ctx context.Context, req *repository.GenerateTokenRequest) (*repository.GenerateTokenResponse, error) {
	accessTokenExp := time.Now().Add(r.accessTokenExpiryInMinutes)
	refreshTokenExp := time.Now().Add(r.refreshTokenExpiryInHour)

	accessTokenClaim := &accessTokenJwtClaim{
		req.Username,
		req.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExp),
		},
	}

	refreshTokenClaim := &refreshTokenJwtClaim{
		req.Username,
		req.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExp),
		},
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)
	newAccessTokenStr, err := newAccessToken.SignedString([]byte(r.accessTokenSecretKey))
	if err != nil {
		return nil, err
	}

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)
	newRefreshTokenStr, err := newRefreshToken.SignedString([]byte(r.refreshTokenSecretKey))
	if err != nil {
		return nil, err
	}

	return &repository.GenerateTokenResponse{
		AccessToken:  newAccessTokenStr,
		RefreshToken: newRefreshTokenStr,
	}, nil
}

func (r *authRepository) ValidateAccessToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error) {
	claim := &accessTokenJwtClaim{}
	jwtToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(r.accessTokenSecretKey), nil
	})

	if err != nil {
		return nil, errlib.ErrUnauthenticated
	}

	if !jwtToken.Valid {
		return nil, errlib.ErrUnauthenticated
	}

	return &entity.AccessTokenCredentialClaim{
		Username: claim.Username,
		Role:     claim.Role,
	}, nil
}

func (r *authRepository) GenerateAccessToken(ctx context.Context, req *repository.GenerateTokenRequest) (string, error) {
	accessTokenExp := time.Now().Add(r.accessTokenExpiryInMinutes)
	accessTokenClaim := &accessTokenJwtClaim{
		req.Username,
		req.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExp),
		},
	}
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)
	newAccessTokenStr, err := newAccessToken.SignedString([]byte(r.accessTokenSecretKey))
	if err != nil {
		return "", err
	}
	return newAccessTokenStr, nil
}

func (r *authRepository) ValidateRefreshToken(ctx context.Context, token string) (*entity.AccessTokenCredentialClaim, error) {
	claim := &refreshTokenJwtClaim{}
	jwtToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(r.refreshTokenSecretKey), nil
	})

	if err != nil {
		return nil, errlib.ErrUnauthenticated
	}

	if !jwtToken.Valid {
		return nil, errlib.ErrUnauthenticated
	}

	return &entity.AccessTokenCredentialClaim{
		Username: claim.Username,
		Role:     claim.Role,
	}, nil
}
