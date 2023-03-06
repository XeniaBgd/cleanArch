package http_v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/XeniaBgd/CleanArch/internal/domain/entity"
	book_usecase "github.com/XeniaBgd/CleanArch/internal/domain/usecase/book"
	"github.com/XeniaBgd/CleanArch/internal/transport/http/dto"
	"github.com/gin-gonic/gin"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type BookUsecase interface {
	CreateBook(ctx context.Context, dto book_usecase.CreateBookDTO) (string, error)
	ListAllBooks(ctx context.Context) []entity.BookView
	GetFullBook(ctx context.Context, id string) entity.FullBook
}

type bookHandler struct {
	bookUsecase BookUsecase
}

func NewBookHandler(usecase BookUsecase) *bookHandler {
	return &bookHandler{bookUsecase: usecase}
}

func (h *bookHandler) Register(router *gin.Engine) {
	router.GET(bookURL, h.GetAllBooks)
}

func (h *bookHandler) GetAllBooks(ctx *gin.Context) {
	// books, _ := h.bookService.GetAllBooks(context.Background(), 0, 0)
	ctx.Writer.Write([]byte("books"))
	ctx.Writer.WriteHeader(http.StatusOK)
}

func (h *bookHandler) CreateBook(ctx *gin.Context) {
	var data dto.CreateBookDTO
	defer ctx.Request.Body.Close()
	if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
		return //error
	}

	// validate data

	// MAPPING dto.CreateBookDTO -> book_usecase.CreateBookDTO

	usecaseDTO := book_usecase.CreateBookDTO{
		Name:     "",
		Year:     0,
		AuthorID: "",
	}
	book, err := h.bookUsecase.CreateBook(ctx.Request.Context(), usecaseDTO)
	if err != nil {
		// 200, error: {msg, ..., dev_msg}
		return
	}
	ctx.Writer.Write([]byte(book))
	ctx.Writer.WriteHeader(http.StatusOK)
}
