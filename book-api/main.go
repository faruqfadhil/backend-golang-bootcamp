package main

import (
	"book-api/core/module"
	"book-api/handler"
	"book-api/pkg/api"
	"book-api/repository/auth"
	"book-api/repository/book"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	accessTokenSecret := os.Getenv("ACCESS_TOKEN")
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN")
	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbPort := os.Getenv("DB_PORT")
	// dbHost := os.Getenv("DB_HOST")

	defaultParams := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", "learn", "ruangguru123", "localhost", "3306", "book_api", defaultParams)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUsername, dbPassword, dbHost, dbPort, dbName, defaultParams)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	}
	redisCache := redisInt.NewRedisCacheEngine(initRedis())
	bookRepo := book.NewBookMySqlRepository(db, redisCache)
	bookELK := book.NewELK()
	bookService := module.NewBookService(bookELK, bookRepo)
	bookHandler := handler.NewBookHandler(bookService)
	authRepo := auth.NewAuthRepository(accessTokenSecret, refreshTokenSecret, 60*time.Minute, 24*time.Hour)
	authService := module.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	middlewareSvc := api.NewMiddlewareService(authService)

	router := gin.Default()
	router.POST("/login", authHandler.Login)

	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewareSvc.AuthticateRequest())
	adminGroup.Use(middlewareSvc.AuthorizedRoles([]string{"admin", "super_user", "super_admin"}))
	{
		adminGroup.POST("/create", bookHandler.Create)
		adminGroup.GET("/books", bookHandler.GetAll)
		adminGroup.PUT("/update", bookHandler.Update)
		adminGroup.PATCH("/author", bookHandler.UpdateAuthor)
		adminGroup.DELETE("/:id", bookHandler.DeleteByID)
	}

	studentGroup := router.Group("/student")
	studentGroup.Use(middlewareSvc.AuthticateRequest())
	studentGroup.Use(middlewareSvc.AuthorizedRoles([]string{"student", "admin"}))
	{
		studentGroup.GET("/:id", bookHandler.GetByID)
	}

	routerAuthGroup := router.Group("/auth")
	routerAuthGroup.Use(middlewareSvc.AuthticateRefreshTokenRequest())
	{
		routerAuthGroup.GET("/refresh", authHandler.RefreshToken)
	}

	router.Run("localhost:9000")
}

func initRedis() *redis.Pool {
	redisHost := fmt.Sprintf("%s:%d", "localhost", 6379)
	Redis := &redis.Pool{
		MaxIdle:         80,
		MaxActive:       100,
		IdleTimeout:     time.Duration(10) * time.Minute,
		MaxConnLifetime: time.Duration(10) * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	conn := Redis.Get()
	_, err := conn.Do("PING")
	if err != nil {
		panic(fmt.Sprintf("[ERR] Redis connection failed, %s", err))
	}
	defer conn.Close()
	return Redis
}
