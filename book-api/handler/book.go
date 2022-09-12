package handler

import (
	"book-api/core/entity"
	"book-api/core/module"
	"context"
	"strconv"

	api "book-api/pkg/api"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service module.BookService
}

func NewBookHandler(service module.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) Create(c *gin.Context) {
	var book *entity.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		api.ResponseFailed(c, err)
		return
	}

	err := h.service.Create(context.Background(), book)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseOK(c, nil)
}

func (h *BookHandler) GetByID(c *gin.Context) {
	bookID := c.Param("id")
	bookInt, err := strconv.Atoi(bookID)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	bookDetail, err := h.service.GetByID(context.Background(), bookInt)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}
	api.ResponseOK(c, bookDetail)
}

func (h *BookHandler) GetAll(c *gin.Context) {
	bookDetail, err := h.service.GetAll(context.Background())
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}
	api.ResponseOK(c, bookDetail)
}

func (h *BookHandler) Update(c *gin.Context) {
	var book *entity.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		api.ResponseFailed(c, err)
		return
	}

	err := h.service.UpdateByID(context.Background(), book)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseOK(c, nil)
}

func (h *BookHandler) UpdateAuthor(c *gin.Context) {
	type updateAuthorPayload struct {
		ID     int    `json:"id"`
		Author string `json:"author"`
	}
	var payload *updateAuthorPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		api.ResponseFailed(c, err)
		return
	}

	err := h.service.UpdateAuthorBook(context.Background(), payload.ID, payload.Author)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseOK(c, nil)
}

func (h *BookHandler) DeleteByID(c *gin.Context) {
	bookID := c.Param("id")
	bookInt, err := strconv.Atoi(bookID)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	err = h.service.DeleteByID(context.Background(), bookInt)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}
	api.ResponseOK(c, nil)
}
