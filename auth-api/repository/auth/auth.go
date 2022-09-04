package auth

import (
	"auth-api/core/entity"
	repoInterface "auth-api/core/repository"
	"context"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var users = map[string]*entity.UserDetail{
	"user1": {
		Username: "user1",
		Password: "password1",
		Role:     "student",
	},
	"user2": {
		Username: "user2",
		Password: "password2",
		Role:     "admin",
	},
}

type repository struct {
	tokenSecretKey              string
	tokenExpirationTimeInMinute time.Duration
}

type jwtClaim struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func New(
	tokenSecretKey string,
	tokenExpirationTimeInMinute time.Duration,
) repoInterface.AuthRepository {
	return &repository{
		tokenSecretKey:              tokenSecretKey,
		tokenExpirationTimeInMinute: tokenExpirationTimeInMinute,
	}
}

func (r *repository) AuthenticateCredential(ctx context.Context, cred *entity.Credential) (*entity.UserDetail, error) {
	if _, ok := users[cred.Username]; !ok {
		return nil, fmt.Errorf("wrong username")
	}

	userDetail := users[cred.Username]
	if userDetail.Password != cred.Password {
		return nil, fmt.Errorf("wrong password")
	}

	return userDetail, nil
}

func (r *repository) GenerateToken(ctx context.Context, username, role string) (string, error) {
	expirationTime := time.Now().Add(r.tokenExpirationTimeInMinute)

	claim := &jwtClaim{
		username,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		}}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	newTokenString, err := newToken.SignedString([]byte(r.tokenSecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate new token")
	}

	return newTokenString, nil
}

func (r *repository) ValidateToken(ctx context.Context, token string) (*entity.CredentialClaim, error) {
	claim := &jwtClaim{}
	jwtToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(r.tokenSecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized, err: %v", err.Error())
	}
	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return &entity.CredentialClaim{
		Username: claim.UserName,
		Role:     claim.Role,
	}, nil
}
