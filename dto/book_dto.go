package dto

import (
	"go-booking/entity"

	"github.com/shopspring/decimal"
)

var (
	MESSAGE_FAILED_BOOKS_NOT_FOUND = "Book Not Found"
	MESSAGE_FAILED_PAGE_IS_WRONG   = "Page is in wrong format"
	MESSAGE_SUCCESS_GET_BOOKS      = "Succesfully get books"
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
)
