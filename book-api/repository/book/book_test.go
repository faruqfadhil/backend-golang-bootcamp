package book

import (
	"book-api/core/entity"
	errlib "book-api/pkg/error"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectationResult = &entity.Book{
	ID:          1,
	Title:       "test",
	Description: "asd",
	Author:      "sada",
	TotalPage:   10,
}

func TestMain(m *testing.M) {
	fmt.Println("sebelum unit test")
	bookDatabases = append(bookDatabases, expectationResult)
	m.Run()
	fmt.Println("sesudah unit test")
	bookDatabases = []*entity.Book{}
}

func TestFindByID(t *testing.T) {
	initRepo := NewBookRepo()
	tests := map[string]struct {
		request  int
		response *entity.Book
		err      error
	}{
		"success scenario": {
			request:  1,
			response: expectationResult,
			err:      nil,
		},
		"error scenario": {
			request:  999,
			response: nil,
			err:      errlib.ErrNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := initRepo.FindByID(context.Background(), test.request)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.response, result)
		})
	}
}

func TestFindAll(t *testing.T) {
	initRepo := NewBookRepo()
	t.Run("success scenario", func(t *testing.T) {
		result, err := initRepo.FindAll(context.Background())
		if err != nil {
			t.Fatalf(err.Error())
		}
		assert.Equal(t, bookDatabases, result)
	})

}

func BenchmarkFindAll(b *testing.B) {
	initRepo := NewBookRepo()
	for i := 0; i < b.N; i++ {
		initRepo.FindAll(context.Background())
	}
}
