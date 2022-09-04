package api

import (
	"auth-api/core/module"
	"context"
	"fmt"
	"net/http"

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

func (s *MiddlewareService) AuthenticateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		validate, err := s.authSvc.ValidateToken(context.Background(), token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthenticated", "error": err.Error()})
			ctx.Abort()
			return
		}
		if validate == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthenticated"})
			ctx.Abort()
			return
		}

		if validate != nil {
			ctx.Set("role", validate.Role)
			fmt.Printf("valid = %+v\n", validate)
			ctx.Next()
			return
		}
	}
}

func (s *MiddlewareService) AuthorizeRoles(authorizedRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Get("role"); !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "role not found"})
			ctx.Abort()
			return
		}

		role, _ := ctx.Get("role")

		for _, authorizeRole := range authorizedRoles {
			if role == authorizeRole {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized role"})
		ctx.Abort()
	}
}
