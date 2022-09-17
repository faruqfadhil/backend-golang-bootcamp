package module

import (
	"book-api/core/entity"
	"book-api/core/repository"
	errlib "book-api/pkg/error"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := &repository.BookRepoMock{Mock: mock.Mock{}}
	repo.Mock.On("FindByID", context.Background(), 10).Return(&entity.Book{
		ID:          10,
		Title:       "a",
		Description: "a",
		Author:      "a",
		TotalPage:   10,
	}, nil)
	repo.Mock.On("FindByID", context.Background(), 1).Return(nil, nil)
	repo.Mock.On("Insert", context.Background(), &entity.Book{
		ID:          1,
		Title:       "a",
		Description: "a",
		Author:      "a",
		TotalPage:   10,
	}).Return(nil, nil)

	init := NewBookService(repo)
	tests := map[string]struct {
		request *entity.Book
		err     error
	}{
		"success create": {
			request: &entity.Book{
				ID:          1,
				Title:       "a",
				Description: "a",
				Author:      "a",
				TotalPage:   10,
			},
			err: nil,
		},
		"error already exist": {
			request: &entity.Book{
				ID:          10,
				Title:       "a",
				Description: "a",
				Author:      "a",
				TotalPage:   10,
			},
			err: errlib.ErrBookAlreadyExist,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := init.Create(context.Background(), test.request)
			assert.Equal(t, test.err, err)
		})
	}
}
