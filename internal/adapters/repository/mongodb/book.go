package mongodb

import (
	"github.com/XeniaBgd/CleanArch/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookStorage struct {
	db *mongo.Database
}

func NewBookStorage(db *mongo.Database) bookStorage {
	return bookStorage{db: db}
}

func (bs bookStorage) GetOne(id string) (entity.Book, error) {
	return entity.Book{}, nil
}

func (bs bookStorage) GetAll(limit, offset int) ([]entity.Book, error) {
	return nil, nil
}

func (bs bookStorage) Create(book entity.Book) (entity.Book, error) {
	return entity.Book{}, nil
}

func (bs bookStorage) Delete(book entity.Book) error {
	return nil
}
