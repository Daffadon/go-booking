package dto

import (
	"errors"
	"go-booking/entity"
	"mime/multipart"

	"github.com/shopspring/decimal"
)

var (
	ErrFailedToCreateBook  = errors.New("failed to create book")
	ErrFailedToUploadCover = errors.New("failed to upload cover")
	ErrFailedBooksNotFound = errors.New("books not found")
	ErrFailedDeleteBook    = errors.New("failed to delete book")
	ErrFailedToGetBook     = errors.New("failed to get books")
	ErrFailedUpdateBook    = errors.New("failed to update book")
)
var (
	MESSAGE_SUCCESS_GET_BOOKS   = "Succesfully get books"
	MESSAGE_SUCCESS_CREATE_BOOK = "Succesfully add book"
	MESSAGE_SUCCESS_DELETE_BOOK = "Succesfully Delete book"
	MESSAGE_SUCCESS_UPDATE_BOOK = "Succesfully update book"

	MESSAGE_FAILED_WRONG_FILE_EXT  = "Wrong file extension. allowed: jpg, jpeg, png"
	MESSAGE_FAILED_BOOKS_NOT_FOUND = "Book Not Found"
	MESSAGE_FAILED_CREATE_BOOK     = "Failed to create book"
	MESSAGE_FAILED_PAGE_IS_WRONG   = "Page is in wrong format"
	MESSAGE_FAILED_GET_DATA        = "data is missing"
	MESSAGE_FAILED_DELETE_BOOK     = "Failed to delete book"
	MESSAGE_FAILED_UPDATE_BOOK     = "Failed to Update Book"
)

type (
	BookGetAllWithPaginationResponse struct {
		Books []entity.Book `json:"books"`
		PaginationResponse
	}
	BookResponseWithoutTimestamp struct {
		ID          string          `json:"id"`
		Title       string          `json:"title"`
		Author      string          `json:"author"`
		Cover       string          `json:"cover"`
		Description string          `json:"description"`
		Stock       int             `json:"stock"`
		Price       decimal.Decimal `json:"price" `
	}

	BookAllResponse struct {
		Data []BookResponseWithoutTimestamp `json:"data"`
		PaginationResponse
	}

	BookGetByIDRequest struct {
		ID string `uri:"id" binding:"required"`
	}
	BookCreateRequest struct {
		Title       string                `json:"title" form:"title" binding:"required"`
		Author      string                `json:"author" form:"author" binding:"required"`
		Cover       *multipart.FileHeader `json:"cover" form:"cover" binding:"required"`
		Description string                `json:"description" form:"description" binding:"required"`
		Stock       int                   `json:"stock" form:"stock" binding:"required"`
		Price       decimal.Decimal       `json:"price" form:"price" binding:"required"`
	}
	BookUpdateParam struct {
		ID string `uri:"id" binding:"required"`
	}
	BookUpdateRequest struct {
		ID          string                `json:"id" form:"id"`
		Title       *string               `json:"title,omitempty" form:"title"`
		Author      *string               `json:"author,omitempty" form:"author"`
		Cover       *multipart.FileHeader `json:"cover,omitempty" form:"cover"`
		Description *string               `json:"description,omitempty" form:"description"`
		Stock       *int                  `json:"stock,omitempty" form:"stock"`
		Price       *decimal.Decimal      `json:"price,omitempty" form:"price"`
	}

	BookDeleteRequest struct {
		ID string `uri:"id" binding:"required"`
	}
)
