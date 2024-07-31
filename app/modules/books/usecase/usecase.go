package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	repository "github.com/1rhino/clean_architecture/app/modules/books/repositories"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	CreateBook(ctx *gin.Context, bookInput *models.BookInput, userID uint) (*models.BookResponse, error)
	GetAllBooks() ([]*models.BookResponse, error)
	GetBooks(userID uint) ([]*models.BookResponse, error)
	GetBook(bookID uint) (*models.BookResponse, error)
	UpdateBook(ctx *gin.Context, bookInput *models.UpdateBook) (*models.BookResponse, error)
	DeleteBook(bookID uint) error
}

type BookUseCase struct {
	bookRepo repository.BookRepository
}

func NewBookUseCase(bookRepo repository.BookRepository) UseCase {
	return &BookUseCase{bookRepo: bookRepo}
}

func (u *BookUseCase) CreateBook(ctx *gin.Context, bookInput *models.BookInput, userID uint) (*models.BookResponse, error) {
	book := &models.Book{
		Name:        bookInput.Name,
		Image:       bookInput.Image,
		Author:      bookInput.Author,
		CategoryID:  bookInput.CategoryID,
		UserID:      userID,
		PublicDate:  bookInput.PublicDate,
		Description: bookInput.Description,
	}

	createBook, err := u.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}
	return models.FilterBookRecord(createBook), nil
}

func (u *BookUseCase) GetAllBooks() ([]*models.BookResponse, error) {
	bookCategories, err := u.bookRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var bookResponses []*models.BookResponse
	for _, book := range bookCategories {
		bookResponses = append(bookResponses, models.FilterBookRecord(book))
	}
	return bookResponses, nil
}

func (u *BookUseCase) GetBooks(userID uint) ([]*models.BookResponse, error) {
	books, err := u.bookRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var bookResponses []*models.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, models.FilterBookRecord(book))
	}
	return bookResponses, nil
}

func (u *BookUseCase) GetBook(bookID uint) (*models.BookResponse, error) {
	book, err := u.bookRepo.FindByID(bookID)
	if err != nil {
		return nil, err
	}
	return models.FilterBookRecord(book), nil
}

func (u *BookUseCase) UpdateBook(ctx *gin.Context, bookInput *models.UpdateBook) (*models.BookResponse, error) {
	book, err := u.bookRepo.FindByID(bookInput.ID)
	if err != nil {
		return nil, err
	}

	book.Name = bookInput.Name
	book.Author = bookInput.Author
	book.PublicDate = bookInput.PublicDate
	book.Description = bookInput.Description
	if bookInput.Image != "" {
		book.Image = bookInput.Image
	}

	updatedBook, err := u.bookRepo.Update(book)
	if err != nil {
		return nil, err
	}

	return models.FilterBookRecord(updatedBook), nil
}

func (u *BookUseCase) DeleteBook(bookID uint) error {
	err := u.bookRepo.Delete(bookID)
	if err != nil {
		return err
	}
	return nil
}
