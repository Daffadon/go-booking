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
		GetBookByID(ctx *gin.Context)
		CreateBook(ctx *gin.Context)
		UpdateBook(ctx *gin.Context)
		DeleteBook(ctx *gin.Context)
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

func (b *bookController) GetBookByID(ctx *gin.Context) {
	var id dto.BookGetByIDRequest
	if err := ctx.ShouldBindUri(&id); err != nil {
		res := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_PAGE_IS_WRONG)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	book, err := b.bookService.GetBookByID(ctx.Request.Context(), id.ID)
	if err != nil {
		res := utils.ReturnResponseError(404, dto.MESSAGE_FAILED_BOOKS_NOT_FOUND)
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := utils.ReturnResponseSuccess(200, dto.MESSAGE_SUCCESS_GET_BOOKS, book, nil)
	ctx.JSON(http.StatusOK, res)
}

var req dto.BookCreateRequest

func (b *bookController) CreateBook(ctx *gin.Context) {
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_GET_DATA)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	createdBook, err := b.bookService.CreateBook(ctx.Request.Context(), req)
	if err != nil {
		res := utils.ReturnResponseError(500, dto.MESSAGE_FAILED_CREATE_BOOK)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.ReturnResponseSuccess(201, dto.MESSAGE_SUCCESS_CREATE_BOOK, createdBook, nil)
	ctx.JSON(http.StatusCreated, res)
}
func (b *bookController) UpdateBook(ctx *gin.Context) {
	var bookId dto.BookUpdateParam
	if err := ctx.ShouldBindUri(&bookId); err != nil {
		res := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_GET_DATA)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var bookRequest dto.BookUpdateRequest
	if err := ctx.ShouldBind(&bookRequest); err != nil {
		res := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_GET_DATA)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	bookRequest.ID = bookId.ID
	updatedBook, err := b.bookService.UpdateBook(ctx.Request.Context(), bookRequest)
	if err != nil {
		res := utils.ReturnResponseError(500, dto.MESSAGE_FAILED_UPDATE_BOOK)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.ReturnResponseSuccess(200, dto.MESSAGE_SUCCESS_UPDATE_BOOK, updatedBook, nil)
	ctx.JSON(http.StatusOK, res)

}

func (b *bookController) DeleteBook(ctx *gin.Context) {
	var id dto.BookDeleteRequest
	if err := ctx.ShouldBindUri(&id); err != nil {
		res := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_GET_DATA)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err := b.bookService.DeleteBook(ctx.Request.Context(), id.ID)
	if err != nil {
		res := utils.ReturnResponseError(500, dto.MESSAGE_FAILED_DELETE_BOOK)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.ReturnResponseSuccess(200, dto.MESSAGE_SUCCESS_DELETE_BOOK, nil, nil)
	ctx.JSON(http.StatusOK, res)
}
