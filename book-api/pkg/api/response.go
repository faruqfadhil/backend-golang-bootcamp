package api

import (
	"errors"
	"net/http"

	errlib "book-api/pkg/error"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ResponseFailed(c *gin.Context, err error) {
	statusCode := http.StatusInternalServerError
	if errors.Is(err, errlib.ErrNotFound) || errors.Is(err, errlib.ErrNoRowsAffected) {
		statusCode = http.StatusNotFound
	}
	if errors.Is(err, errlib.ErrBookAlreadyExist) {
		statusCode = http.StatusBadRequest
	}

	if errors.Is(err, errlib.ErrUnauthenticated) || errors.Is(err, errlib.ErrUnauthorized) {
		statusCode = http.StatusUnauthorized
	}

	c.JSON(statusCode, &ResponseError{
		Status:  "error",
		Message: err.Error(),
	})
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseSuccess{
		Data:    data,
		Message: "success",
	})
}
