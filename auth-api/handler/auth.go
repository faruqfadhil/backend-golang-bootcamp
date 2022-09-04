package handler

import (
	"auth-api/core/entity"
	"auth-api/core/module"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc module.AuthService
}

type ResponseError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(svc module.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload *entity.Credential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, &ResponseError{
			Status:  "error",
			Code:    1,
			Message: "err",
		})
		return
	}

	resp, err := h.svc.Login(context.Background(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, &ResponseError{
			Status:  "error",
			Code:    2,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) Welcome(c *gin.Context) {
	role, _ := c.Get("role")
	status := fmt.Sprintf("you are logged in with role %s", role)
	c.JSON(http.StatusOK, gin.H{"status": status})
}
