package main

import (
	"auth-api/core/module"
	"auth-api/handler"
	"auth-api/pkg/api"
	authRepo "auth-api/repository/auth"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	authRepo := authRepo.New("secretfrq", 10*time.Minute)
	authSvc := module.NewAuthService(authRepo)
	authHandler := handler.New(authSvc)

	middlewareSvc := api.NewMiddlewareService(authSvc)

	router := gin.New()
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authHandler.Login)
	}

	contentRouter := router.Group("/content")
	contentRouter.Use(middlewareSvc.AuthenticateRequest())
	contentRouter.Use(middlewareSvc.AuthorizeRoles([]string{"student", "admin"}))
	{
		contentRouter.GET("/welcome", authHandler.Welcome)
	}

	contentStudentRouter := router.Group("/content/student")
	contentStudentRouter.Use(middlewareSvc.AuthenticateRequest())
	contentStudentRouter.Use(middlewareSvc.AuthorizeRoles([]string{"student"}))
	{
		contentStudentRouter.GET("/welcome", authHandler.Welcome)
	}

	contentAdminRouter := router.Group("/content/admin")
	contentAdminRouter.Use(middlewareSvc.AuthenticateRequest())
	contentAdminRouter.Use(middlewareSvc.AuthorizeRoles([]string{"admin"}))
	{
		contentAdminRouter.GET("/welcome", authHandler.Welcome)
	}

	router.Run("localhost:9091")
}
