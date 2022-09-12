package book

import (
	"book-api/core/entity"
	"book-api/core/repository"
	errlib "book-api/pkg/error"
	"context"
	"errors"

	"gorm.io/gorm"
)

const (
	bookTable = "book"
)

type bookMySqlRepository struct {
	db *gorm.DB
}

func NewBookMySqlRepository(db *gorm.DB) repository.BookRepository {
	return &bookMySqlRepository{db: db}
}

func (r *bookMySqlRepository) FindByID(ctx context.Context, ID int) (*entity.Book, error) {
	var book *entity.Book
	err := r.db.Debug().Table(bookTable).
		Where("id = ?", ID).
		Where("is_deleted = ?", false).
		Take(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errlib.ErrNotFound
		}
		return nil, errlib.ErrDatabaseError
	}
	return book, nil
}
func (r *bookMySqlRepository) Insert(ctx context.Context, book *entity.Book) error {
	err := r.db.Debug().Table(bookTable).Create(&book).Error
	if err != nil {
		return errlib.ErrDatabaseError
	}
	// err := r.db.Debug().Raw("INSERT INTO book(title, description, author, total_page)VALUES (?,?,?,?)", book.Title, book.Description, book.Author, book.TotalPage).Error
	// if err != nil {
	// 	return errlib.ErrDatabaseError
	// }
	return nil
}

func (r *bookMySqlRepository) FindAll(ctx context.Context) ([]*entity.Book, error) {
	var books []*entity.Book
	err := r.db.Debug().Table(bookTable).Where("is_deleted = ?", false).
		Find(&books).Error
	if err != nil {
		return nil, errlib.ErrDatabaseError
	}
	return books, nil
}

func (r *bookMySqlRepository) Update(ctx context.Context, book *entity.Book) error {
	err := r.db.Debug().Table(bookTable).Where("id = ?", book.ID).
		Where("is_deleted = ?", false).Updates(&book).Error
	if err != nil {
		return errlib.ErrDatabaseError
	}
	return nil
}

func (r *bookMySqlRepository) UpdateAuthorByID(ctx context.Context, ID int, newAuthor string) error {
	err := r.db.Debug().Table(bookTable).Where("id = ?", ID).Where("is_deleted = ?", false).Updates(map[string]interface{}{"author": newAuthor}).Error
	if err != nil {
		return errlib.ErrDatabaseError
	}
	return nil
}
func (r *bookMySqlRepository) DeleteByID(ctx context.Context, ID int) error {
	err := r.db.Debug().Table(bookTable).Where("id = ?", ID).Updates(map[string]interface{}{"is_deleted": true}).Error
	if err != nil {
		return errlib.ErrDatabaseError
	}
	return nil
}
