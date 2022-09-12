package book

import (
	"book-api/core/entity"
	"book-api/core/repository"
	errlib "book-api/pkg/error"
	"context"
	"fmt"
)

var bookDatabases []*entity.Book

type bookRepo struct {
}

func NewBookRepo() repository.BookRepository {
	return &bookRepo{}
}

func (r *bookRepo) FindByID(ctx context.Context, ID int) (*entity.Book, error) {
	for _, b := range bookDatabases {
		if b.ID == ID {
			return b, nil
		}
	}
	return nil, errlib.ErrNotFound
}

func (r *bookRepo) Insert(ctx context.Context, book *entity.Book) error {
	bookDatabases = append(bookDatabases, book)
	for _, b := range bookDatabases {
		fmt.Printf("list book = %+v\n", b)
	}
	return nil
}

func (r *bookRepo) FindAll(ctx context.Context) ([]*entity.Book, error) {
	return bookDatabases, nil
}

func (r *bookRepo) Update(ctx context.Context, book *entity.Book) error {
	for _, existingBook := range bookDatabases {
		if existingBook.ID == book.ID {
			existingBook.Author = book.Author
			existingBook.Description = book.Description
			existingBook.Title = book.Title
			existingBook.TotalPage = book.TotalPage
			return nil
		}
	}

	return errlib.ErrNoRowsAffected
}

func (r *bookRepo) UpdateAuthorByID(ctx context.Context, ID int, newAuthor string) error {
	for _, existingBook := range bookDatabases {
		if existingBook.ID == ID {
			existingBook.Author = newAuthor
			return nil
		}
	}
	return errlib.ErrNoRowsAffected
}

func (r *bookRepo) DeleteByID(ctx context.Context, ID int) error {
	for i, existingBook := range bookDatabases {
		if existingBook.ID == ID {
			temp := bookDatabases[:i]
			if i != 0 && len(bookDatabases) < i+1 {
				temp = append(temp, bookDatabases[i+1:]...)
			}

			bookDatabases = temp
			return nil
		}
	}

	return errlib.ErrNoRowsAffected
}
