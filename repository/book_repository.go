package repository

import (
	"context"
	"go-booking/dto"
	"go-booking/entity"
	"math"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	BookRepository interface {
		GetBooksWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BookGetAllWithPaginationResponse, error)
		GetBookByID(ctx context.Context, id string) (dto.BookResponseWithoutTimestamp, error)
		CreateBook(ctx context.Context, req entity.Book) (entity.Book, error)
		UpdateBook(ctx context.Context, req *dto.BookResponseWithoutTimestamp) (dto.BookResponseWithoutTimestamp, error)
		DeleteBook(ctx context.Context, id string) error
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
	if err := b.db.WithContext(ctx).Model(&entity.Book{}).Count(&totalData).Error; err != nil {
		return dto.BookGetAllWithPaginationResponse{}, dto.ErrFailedToGetBook
	}
	if totalData == 0 {
		return dto.BookGetAllWithPaginationResponse{}, dto.ErrFailedBooksNotFound
	}

	totalPage := uint16(math.Ceil(float64(totalData) / float64(perPage)))
	if req.Page > totalPage {
		req.Page = totalPage
	}
	offset := perPage * (int(req.Page) - 1)

	if err := b.db.WithContext(ctx).Scopes(Paginate(perPage, offset)).Find(&books).Error; err != nil {
		return dto.BookGetAllWithPaginationResponse{}, dto.ErrFailedToGetBook
	}

	var nextPage, prevPage uint16

	if req.Page == totalPage {
		nextPage = req.Page
	} else {
		nextPage = req.Page + 1
	}

	if req.Page == 1 {
		prevPage = req.Page
	} else {
		prevPage = req.Page - 1
	}

	return dto.BookGetAllWithPaginationResponse{
		Books: books,
		PaginationResponse: dto.PaginationResponse{
			Page:      req.Page,
			PrevPage:  prevPage,
			NextPage:  nextPage,
			TotalPage: totalPage,
		},
	}, nil
}

func (b *bookRepository) GetBookByID(ctx context.Context, id string) (dto.BookResponseWithoutTimestamp, error) {
	var book entity.Book
	if err := b.db.WithContext(ctx).Where("id = ?", id).First(&book).Error; err != nil {
		if err.Error() == "record not found" {
			return dto.BookResponseWithoutTimestamp{}, dto.ErrFailedBooksNotFound
		} else {
			return dto.BookResponseWithoutTimestamp{}, dto.ErrFailedToGetBook
		}
	}
	return dto.BookResponseWithoutTimestamp{
		ID:          book.ID.String(),
		Title:       book.Title,
		Author:      book.Author,
		Cover:       book.Cover,
		Description: book.Description,
		Stock:       book.Stock,
		Price:       book.Price,
	}, nil
}

func (b *bookRepository) CreateBook(ctx context.Context, req entity.Book) (entity.Book, error) {
	res := b.db.WithContext(ctx).Create(&req)
	if res.Error != nil {
		return entity.Book{}, dto.ErrFailedToCreateBook
	}
	return req, nil
}
func (b *bookRepository) UpdateBook(ctx context.Context, req *dto.BookResponseWithoutTimestamp) (dto.BookResponseWithoutTimestamp, error) {
	book := entity.Book{
		ID:          uuid.MustParse(req.ID),
		Title:       req.Title,
		Author:      req.Author,
		Cover:       req.Cover,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
	}
	if res := b.db.WithContext(ctx).Updates(&book); res.Error != nil {
		return dto.BookResponseWithoutTimestamp{}, dto.ErrFailedUpdateBook
	}
	return *req, nil
}

func (b *bookRepository) DeleteBook(ctx context.Context, id string) error {
	res := b.db.WithContext(ctx).Model(entity.Book{}).Where("id = ?", id).Delete(&entity.Book{})
	if res.RowsAffected == 0 {
		return dto.ErrFailedBooksNotFound
	}
	if res.Error != nil {
		return dto.ErrFailedDeleteBook
	}
	return nil
}
