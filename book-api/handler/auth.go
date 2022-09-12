package handler

import (
	"book-api/core/entity"
	"book-api/core/module"
	api "book-api/pkg/api"
	"context"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService module.AuthService
}

func NewAuthHandler(authService module.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var cred *entity.Credential
	if err := c.ShouldBindJSON(&cred); err != nil {
		api.ResponseFailed(c, err)
		return
	}

	response, err := h.authService.Login(context.Background(), cred)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseOK(c, response)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// var cred *entity.Credential
	// if err := c.ShouldBindJSON(&cred); err != nil {
	// 	api.ResponseFailed(c, err)
	// 	return
	// }

	role, _ := c.Get("role")
	username, _ := c.Get("username")
	response, err := h.authService.RefreshToken(context.Background(), &module.RefreshTokenRequest{
		Username: username.(string),
		Role:     role.(string),
	})
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseOK(c, response)
}
