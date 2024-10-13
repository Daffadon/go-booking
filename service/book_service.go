package service

import (
	"context"
	"go-booking/dto"
	"go-booking/repository"
)

type (
	BookService interface {
		GetBooksWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BookAllResponse, error)
		GetBookByID(ctx context.Context, id string) (dto.BookResponseWithoutTimestamp, error)
	}
	bookService struct {
		bookRepository repository.BookRepository
	}
)

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{bookRepository}
}

func (b *bookService) GetBooksWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BookAllResponse, error) {
	booksWithPaginate, err := b.bookRepository.GetBooksWithPagination(ctx, req)
	if err != nil {
		return dto.BookAllResponse{}, err
	}
	var bookResponse []dto.BookResponseWithoutTimestamp
	for _, book := range booksWithPaginate.Books {
		data := dto.BookResponseWithoutTimestamp{
			ID:          book.ID.String(),
			Title:       book.Title,
			Author:      book.Author,
			Cover:       book.Cover,
			Description: book.Description,
			Stock:       book.Stock,
			Price:       book.Price,
		}
		bookResponse = append(bookResponse, data)
	}
	return dto.BookAllResponse{
		Data:               bookResponse,
		PaginationResponse: booksWithPaginate.PaginationResponse,
	}, nil
}

func (b *bookService) GetBookByID(ctx context.Context, id string) (dto.BookResponseWithoutTimestamp, error) {
	book, err := b.bookRepository.GetBookByID(ctx, id)
	if err != nil {
		return dto.BookResponseWithoutTimestamp{}, err
	}
	return book, nil
}

