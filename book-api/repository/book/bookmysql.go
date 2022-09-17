package book

import (
	"book-api/core/entity"
	"book-api/core/repository"
	"book-api/pkg/cache"
	errlib "book-api/pkg/error"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	bookTable              = "book"
	findByIDCacheKeyPrefix = "FindByID"
	findAllCacheKeyPrefix  = "FindAll"
)

type bookMySqlRepository struct {
	db          *gorm.DB
	cacheEngine cache.Cache
}

func NewBookMySqlRepository(db *gorm.DB, cacheEngine cache.Cache) repository.BookRepository {
	return &bookMySqlRepository{db: db, cacheEngine: cacheEngine}
}

func (r *bookMySqlRepository) FindByID(ctx context.Context, ID int) (*entity.Book, error) {
	var book *entity.Book
	cacheKey := fmt.Sprintf("%s:%d", findByIDCacheKeyPrefix, ID)
	respInBytes, err := r.cacheEngine.Get(cacheKey)
	if err != nil {
		fmt.Println("error when fetch cache, err: ", err)
	}

	if len(respInBytes) > 0 {
		if err := json.Unmarshal(respInBytes, &book); err != nil {
			return nil, err
		}
		fmt.Println("cache hit")
		return book, nil
	}

	fmt.Println("cache miss.. try to fetch from DB")

	err = r.db.Debug().Table(bookTable).
		Where("id = ?", ID).
		Where("is_deleted = ?", false).
		Take(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errlib.ErrNotFound
		}
		return nil, errlib.ErrDatabaseError
	}

	dataInBytes, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	err = r.cacheEngine.Set(cacheKey, dataInBytes, 3600)
	if err != nil {
		return nil, err
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
	cacheKey := findAllCacheKeyPrefix
	respInBytes, err := r.cacheEngine.Get(cacheKey)
	if err != nil {
		fmt.Println("error when fetch cache, err: ", err)
	}

	if len(respInBytes) > 0 {
		if err := json.Unmarshal(respInBytes, &books); err != nil {
			return nil, err
		}
		fmt.Println("cache hit")
		return books, nil
	}

	fmt.Println("cache miss.. try to fetch from DB")
	err = r.db.Debug().Table(bookTable).Where("is_deleted = ?", false).
		Find(&books).Error
	if err != nil {
		return nil, errlib.ErrDatabaseError
	}

	dataInBytes, err := json.Marshal(books)
	if err != nil {
		return nil, err
	}

	err = r.cacheEngine.Set(cacheKey, dataInBytes, 3600)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookMySqlRepository) Update(ctx context.Context, book *entity.Book) error {
	err := r.db.Debug().Table(bookTable).Where("id = ?", book.ID).
		Where("is_deleted = ?", false).Updates(&book).Error
	if err != nil {
		return errlib.ErrDatabaseError
	}

	err = r.cacheEngine.Del(fmt.Sprintf("%s:%d", findByIDCacheKeyPrefix, book.ID), findAllCacheKeyPrefix)
	if err != nil {
		return err
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
