package repository

import (
	"book-api/core/entity"
	"context"

	"github.com/stretchr/testify/mock"
)

type BookRepoMock struct {
	Mock mock.Mock
}

func (r *BookRepoMock) FindByID(ctx context.Context, ID int) (*entity.Book, error) {
	argument := r.Mock.Called(ctx, ID)
	if argument.Get(0) == nil {
		return nil, argument.Error(1)
	}
	return argument.Get(0).(*entity.Book), nil
}

func (r *BookRepoMock) Insert(ctx context.Context, book *entity.Book) error {
	argument := r.Mock.Called(ctx, book)
	if argument.Error(0) != nil {
		return argument.Error(0)
	}
	return nil
}

func (r *BookRepoMock) FindAll(ctx context.Context) ([]*entity.Book, error) {
	return nil, nil
}
func (r *BookRepoMock) Update(ctx context.Context, book *entity.Book) error {
	return nil
}
func (r *BookRepoMock) UpdateAuthorByID(ctx context.Context, ID int, newAuthor string) error {
	return nil
}
func (r *BookRepoMock) DeleteByID(ctx context.Context, ID int) error {
	return nil
}
