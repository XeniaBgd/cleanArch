package mongodb

import (
	"github.com/XeniaBgd/CleanArch/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type authorStorage struct {
	db *mongo.Database
}

func NewAuthorStorage(db *mongo.Database) authorStorage {
	return authorStorage{db: db}
}

func (as authorStorage) GetOne(id string) (entity.Author, error) {
	return entity.Author{}, nil
}

func (as authorStorage) GetAll(limit, offset int) ([]entity.Author, error) {
	return nil, nil
}

func (as authorStorage) Create(book entity.Author) (entity.Author, error) {
	return entity.Author{}, nil
}

func (as authorStorage) Delete(book entity.Author) error {
	return nil
}
