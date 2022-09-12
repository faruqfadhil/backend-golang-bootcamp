package repository

import (
	"book-api/core/entity"
	"context"
)

type BookRepository interface {
	FindByID(ctx context.Context, ID int) (*entity.Book, error)
	Insert(ctx context.Context, book *entity.Book) error
	FindAll(ctx context.Context) ([]*entity.Book, error)
	Update(ctx context.Context, book *entity.Book) error
	UpdateAuthorByID(ctx context.Context, ID int, newAuthor string) error
	DeleteByID(ctx context.Context, ID int) error
}
