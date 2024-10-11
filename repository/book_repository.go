package repository

import (
	"context"
	"go-booking/dto"
	"go-booking/entity"
	"math"

	"gorm.io/gorm"
)

type (
	BookRepository interface {
		GetBooksWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BookGetAllWithPaginationResponse, error)
	}
	bookRepository struct {
		db *gorm.DB
	}
)

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) GetBooksWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BookGetAllWithPaginationResponse, error) {
	perPage := 10
	if req.Page == 0 {
		req.Page = 1
	}
	var totalData int64
	var books []entity.Book
	offset := perPage * (int(req.Page) - 1)
	if err := b.db.WithContext(ctx).Model(&entity.Book{}).Count(&totalData).Error; err != nil {
		return dto.BookGetAllWithPaginationResponse{}, err
	}
	if err := b.db.WithContext(ctx).Scopes(Paginate(perPage, offset)).Find(&books).Error; err != nil {
		return dto.BookGetAllWithPaginationResponse{}, err
	}
	totalPage := int64(math.Ceil(float64(totalData) / float64(perPage)))
	return dto.BookGetAllWithPaginationResponse{
		Books: books,
		PaginationResponse: dto.PaginationResponse{
			Page:      req.Page,
			NextPage:  req.Page + 1,
			TotalPage: uint16(totalPage),
		},
	}, nil
}
