package service

import (
	"context"

	"github.com/XeniaBgd/CleanArch/internal/domain/entity"
)

type AuthorStorage interface {
	GetOne(id string) (entity.Author, error)
	GetAll(limit, offset int) ([]entity.Author, error)
	Create(book entity.Author) (entity.Author, error)
	Delete(book entity.Author) error
}

type Service interface {
	GetByID(ctx context.Context, uuid string) (entity.Author, error)
	GetAll(ctx context.Context, limit, offset int) ([]entity.Author, error)
	Create(ctx context.Context) (entity.Author, error)
}

type authorService struct {
	storage AuthorStorage
}

func NewAuthorService(storage AuthorStorage) authorService {
	return authorService{storage: storage}
}

func (s authorService) Create(ctx context.Context) (entity.Author, error) {
	return entity.Author{}, nil
}

func (s authorService) GetByID(ctx context.Context, id string) (entity.Author, error) {
	return s.storage.GetOne(id)
}

func (s authorService) GetAll(ctx context.Context, limit, offset int) ([]entity.Author, error) {
	return s.storage.GetAll(limit, offset)
}
