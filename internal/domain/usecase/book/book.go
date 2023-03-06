package book_usecase

import (
	"context"

	"github.com/XeniaBgd/CleanArch/internal/domain/entity"
)

type BookService interface {
	GetAllForList(ctx context.Context) ([]entity.BookView, error)
	GetByID(ctx context.Context, uuid string) (entity.Book, error)
}

type AuthorService interface {
	GetByID(ctx context.Context, uuid string) (entity.Author, error)
}

type bookUsecase struct {
	bookService   BookService
	authorService AuthorService
}

func NewBookUsecase(bookService BookService, authorService AuthorService) bookUsecase {
	return bookUsecase{bookService: bookService, authorService: authorService}
}

func (u bookUsecase) CreateBook(ctx context.Context, dto CreateBookDTO) (string, error) {
	return "", nil
}

func (u bookUsecase) ListAllBooks(ctx context.Context) []entity.BookView {
	res, _ := u.bookService.GetAllForList(ctx)
	return res
}

func (u bookUsecase) GetFullBook(ctx context.Context, id string) entity.FullBook {
	book, _ := u.bookService.GetByID(ctx, id)
	author, _ := u.authorService.GetByID(ctx, book.AuthorID)

	return entity.FullBook{
		Book:   book,
		Author: author,
	}
}
