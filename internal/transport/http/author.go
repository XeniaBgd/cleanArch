package http_v1

import (
	"github.com/gin-gonic/gin"
)

type AuthorUsecase interface {
}

type authorHandler struct {
	authorUsecase AuthorUsecase
}

func NewAuthorHandler(usecase AuthorUsecase) *authorHandler {
	return &authorHandler{authorUsecase: usecase}
}

func (h *authorHandler) Register(router *gin.Engine) {

}
