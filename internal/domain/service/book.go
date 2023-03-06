package service

import (
	"context"

	"github.com/XeniaBgd/CleanArch/internal/domain/entity"
)

type BookStorage interface {
	GetOne(uuid string) (entity.Book, error)
	GetAll(limit, offset int) ([]entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Delete(book entity.Book) error
}

type BookService interface {
	GetByID(ctx context.Context, uuid string) (entity.Book, error)
	GetAll(ctx context.Context, limit, offset int) ([]entity.Book, error)
	Create(ctx context.Context) (entity.Book, error)
}

type bookService struct {
	storage BookStorage
}

func NewBookService(storage BookStorage) bookService {
	return bookService{storage: storage}
}

func (s bookService) Create(ctx context.Context) (entity.Book, error) {
	return entity.Book{}, nil
}

func (s bookService) GetByID(ctx context.Context, uuid string) (entity.Book, error) {
	return s.storage.GetOne(uuid)
}

func (s bookService) GetAll(ctx context.Context, limit, offset int) ([]entity.Book, error) {
	return s.storage.GetAll(limit, offset)
}

func (s bookService) GetAllForList(ctx context.Context) ([]entity.BookView, error) {

	return nil, nil
}
