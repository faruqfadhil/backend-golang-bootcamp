package book

import (
	"book-api/core/entity"
	"book-api/core/repository"
	"context"
)

type elastic struct{}

func NewELK() repository.BookRepository {
	return &elastic{}
}

func (e *elastic) FindByID(ctx context.Context, ID int) (*entity.Book, error) {
	return nil, nil
}
func (e *elastic) Insert(ctx context.Context, book *entity.Book) error {
	return nil
}
func (e *elastic) FindAll(ctx context.Context) ([]*entity.Book, error) {
	return nil, nil
}
func (e *elastic) Update(ctx context.Context, book *entity.Book) error {
	return nil
}
func (e *elastic) UpdateAuthorByID(ctx context.Context, ID int, newAuthor string) error {
	return nil
}
func (e *elastic) DeleteByID(ctx context.Context, ID int) error {
	return nil
}
