package module

import (
	"book-api/core/entity"
	"book-api/core/repository"
	errlib "book-api/pkg/error"
	"context"
	"errors"
)

type BookService interface {
	Create(ctx context.Context, book *entity.Book) error
	GetByID(ctx context.Context, ID int) (*entity.Book, error)
	GetAll(ctx context.Context) ([]*entity.Book, error)
	UpdateByID(ctx context.Context, book *entity.Book) error
	UpdateAuthorBook(ctx context.Context, ID int, newAuthor string) error
	DeleteByID(ctx context.Context, ID int) error
}

type bookService struct {
	repo       repository.BookRepository
	repoBackup repository.BookRepository
}

func NewBookService(repo repository.BookRepository, repoBackup repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) Create(ctx context.Context, book *entity.Book) error {
	// check, apakah buku dengan ID yg sama sudah ada,
	// jika iya, maka return error id already exist.
	// jika tidak, maka boleh insert buku.

	existingBook, err := s.repo.FindByID(ctx, book.ID)
	if err != nil && !errors.Is(err, errlib.ErrNotFound) {
		return err
	}

	if existingBook != nil && existingBook.ID == book.ID {
		return errlib.ErrBookAlreadyExist
	}

	err = s.repo.Insert(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (s *bookService) GetByID(ctx context.Context, ID int) (*entity.Book, error) {
	// Get ke database by id, kalau ada return bukunya, kalau gak ada return error not found.
	return s.repo.FindByID(ctx, ID)
}

func (s *bookService) GetAll(ctx context.Context) ([]*entity.Book, error) {
	// Get semua list buku.
	return s.repo.FindAll(ctx)
}

func (s *bookService) UpdateByID(ctx context.Context, book *entity.Book) error {
	// Get book by id, kalau gak nemu berati return not found.
	// Kalau ketemu bukunya, lakukan update data.
	existingBook, err := s.repo.FindByID(ctx, book.ID)
	if err != nil {
		return err
	}

	if existingBook == nil {
		return errlib.ErrNotFound
	}

	return s.repo.Update(ctx, book)
}

func (s *bookService) UpdateAuthorBook(ctx context.Context, ID int, newAuthor string) error {
	// get book by id, kalau gak ada return not found
	// Update author by book ID.
	existingBook, err := s.repo.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	if existingBook == nil {
		return errlib.ErrNotFound
	}

	return s.repo.UpdateAuthorByID(ctx, ID, newAuthor)
}

func (s *bookService) DeleteByID(ctx context.Context, ID int) error {
	// Get by id dulu, kalau gak ada return not found
	// kalau ada, delete book by id
	existingBook, err := s.repo.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	if existingBook == nil {
		return errlib.ErrNotFound
	}

	return s.repo.DeleteByID(ctx, ID)
}
