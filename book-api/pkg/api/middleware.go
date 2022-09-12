package api

import (
	"book-api/core/module"
	"context"
	"fmt"

	errlib "book-api/pkg/error"

	"github.com/gin-gonic/gin"
)

type MiddlewareService struct {
	authSvc module.AuthService
}

func NewMiddlewareService(authSvc module.AuthService) *MiddlewareService {
	return &MiddlewareService{
		authSvc: authSvc,
	}
}

func (s *MiddlewareService) AuthticateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		validate, err := s.authSvc.ValidateAccessToken(context.Background(), token)
		if err != nil {
			ResponseFailed(ctx, errlib.ErrUnauthenticated)
			ctx.Abort()
			return
		}
		if validate == nil {
			ResponseFailed(ctx, errlib.ErrUnauthenticated)
			ctx.Abort()
			return
		}
		if validate != nil {
			fmt.Printf("berhasil authenticate dengan claim: %+v\n", validate)
			ctx.Set("role", validate.Role)
			ctx.Set("username", validate.Username)
			ctx.Next()
			return
		}
	}
}

func (s *MiddlewareService) AuthorizedRoles(authorizedRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Get("role"); !ok {
			ResponseFailed(ctx, errlib.ErrUnauthorized)
			ctx.Abort()
		}

		role, _ := ctx.Get("role")

		for _, authorizedRole := range authorizedRoles {
			if authorizedRole == role {
				ctx.Next()
				return
			}
		}

		ResponseFailed(ctx, errlib.ErrUnauthorized)
		ctx.Abort()
	}
}

func (s *MiddlewareService) AuthticateRefreshTokenRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		validate, err := s.authSvc.ValidateRefreshToken(context.Background(), token)
		if err != nil {
			ResponseFailed(ctx, errlib.ErrUnauthenticated)
			ctx.Abort()
			return
		}
		if validate == nil {
			ResponseFailed(ctx, errlib.ErrUnauthenticated)
			ctx.Abort()
			return
		}
		if validate != nil {
			fmt.Printf("berhasil authenticate dengan claim: %+v\n", validate)
			ctx.Set("role", validate.Role)
			ctx.Set("username", validate.Username)
			ctx.Next()
			return
		}
	}
}
