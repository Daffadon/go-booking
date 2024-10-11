package controller

import (
	"go-booking/dto"
	"go-booking/service"
	"go-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	BookController interface {
		GetBooks(ctx *gin.Context)
	}
	bookController struct {
		bookService service.BookService
	}
)

func NewBookController(bookService service.BookService) BookController {
	return &bookController{bookService}
}

func (b *bookController) GetBooks(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		res := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_PAGE_IS_WRONG)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	books, err := b.bookService.GetBooksWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.ReturnResponseError(404, dto.MESSAGE_FAILED_BOOKS_NOT_FOUND)
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := utils.ReturnResponseSuccess(200, dto.MESSAGE_SUCCESS_GET_BOOKS, books.Data, books.PaginationResponse)
	ctx.JSON(http.StatusOK, res)
}
