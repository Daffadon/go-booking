package service

import (
	"context"
	"fmt"
	"go-booking/dto"
	"go-booking/entity"
	"go-booking/repository"
	"go-booking/utils"

	"github.com/google/uuid"
)

type (
	BookService interface {
		GetBooksWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BookAllResponse, error)
		GetBookByID(ctx context.Context, id string) (dto.BookResponseWithoutTimestamp, error)
		CreateBook(ctx context.Context, req dto.BookCreateRequest) (entity.Book, error)
		UpdateBook(ctx context.Context, req dto.BookUpdateRequest) (dto.BookResponseWithoutTimestamp, error)
		DeleteBook(ctx context.Context, id string) error
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

func (b *bookService) CreateBook(ctx context.Context, req dto.BookCreateRequest) (entity.Book, error) {
	coverFilename := ""
	if req.Cover != nil {
		coverId := uuid.New().String()
		ext := utils.GetFileExtension(req.Cover.Filename)

		filename := fmt.Sprintf("book/%s.%s", coverId, ext)
		if err := utils.UploadFile(req.Cover, filename); err != nil {
			return entity.Book{}, dto.ErrFailedToUploadCover
		}
	}
	book := entity.Book{
		Title:       req.Title,
		Author:      req.Author,
		Cover:       coverFilename,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
	}

	createdBook, err := b.bookRepository.CreateBook(ctx, book)
	if err != nil {
		return entity.Book{}, dto.ErrFailedToCreateBook
	}
	return createdBook, nil
}
func (b *bookService) UpdateBook(ctx context.Context, req dto.BookUpdateRequest) (dto.BookResponseWithoutTimestamp, error) {
	book, err := b.bookRepository.GetBookByID(ctx, req.ID)
	if err != nil {
		return dto.BookResponseWithoutTimestamp{}, err
	}
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.Cover != nil {
		// if the cover is new, delete pervious cover and upload new
		err := utils.DeleteFile(book.Cover)
		if err != nil {
			return dto.BookResponseWithoutTimestamp{}, err
		}
		coverId := uuid.New().String()
		ext := utils.GetFileExtension(req.Cover.Filename)

		filename := fmt.Sprintf("book/%s.%s", coverId, ext)
		if err := utils.UploadFile(req.Cover, filename); err != nil {
			return dto.BookResponseWithoutTimestamp{}, err
		}
		book.Cover = filename
	}
	if req.Description != nil {
		book.Description = *req.Description
	}
	if req.Price != nil {
		book.Price = *req.Price
	}
	if req.Stock != nil {
		book.Stock = *req.Stock
	}
	bookUpdated, err := b.bookRepository.UpdateBook(ctx, &book)
	if err != nil {
		return dto.BookResponseWithoutTimestamp{}, err
	}
	return bookUpdated, nil
}

func (b *bookService) DeleteBook(ctx context.Context, id string) error {
	err := b.bookRepository.DeleteBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
