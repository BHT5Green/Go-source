package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	repository "github.com/1rhino/clean_architecture/app/modules/book_category/repositories"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	CreateBookCategory(ctx *gin.Context, bookCategory *models.BookCategoryInput, userID uint) (*models.BookCategoryResponse, error)
	GetBookCategories(userID uint) ([]*models.BookCategoryResponse, error)
	GetAllBookCategories() ([]*models.BookCategoryResponse, error)
	GetBookCategory(bookCategoryID uint) (*models.BookCategoryResponse, error)
	UpdateBookCategory(ctx *gin.Context, bookCategoryInput *models.UpdateBookCategory) (*models.BookCategoryResponse, error)
	DeleteBookCategory(bookCategoryID uint) error
}

type BookCategoryUseCase struct {
	bookCategoryRepo repository.BookCategoryRepository
}

func NewBookCategoryUseCase(bookCategoryRepo repository.BookCategoryRepository) UseCase {
	return &BookCategoryUseCase{bookCategoryRepo: bookCategoryRepo}
}

func (u *BookCategoryUseCase) CreateBookCategory(ctx *gin.Context, bookCategoryInput *models.BookCategoryInput, userID uint) (*models.BookCategoryResponse, error) {
	bookCategory := &models.BookCategory{
		Name:        bookCategoryInput.Name,
		Image:       bookCategoryInput.Image,
		Description: bookCategoryInput.Description,
		UserID:      userID,
	}
	createBookCategory, err := u.bookCategoryRepo.Create(bookCategory)
	if err != nil {
		return nil, err
	}
	return models.FilterBookCategoryRecord(createBookCategory), nil
}

func (u *BookCategoryUseCase) GetBookCategories(userID uint) ([]*models.BookCategoryResponse, error) {
	bookCategories, err := u.bookCategoryRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var bookCategoryResponses []*models.BookCategoryResponse
	for _, bookCategory := range bookCategories {
		bookCategoryResponses = append(bookCategoryResponses, models.FilterBookCategoryRecord(bookCategory))
	}
	return bookCategoryResponses, nil
}

func (u *BookCategoryUseCase) GetAllBookCategories() ([]*models.BookCategoryResponse, error) {
	bookCategories, err := u.bookCategoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var bookCategoryResponses []*models.BookCategoryResponse
	for _, bookCategory := range bookCategories {
		bookCategoryResponses = append(bookCategoryResponses, models.FilterBookCategoryRecord(bookCategory))
	}
	return bookCategoryResponses, nil
}

func (u *BookCategoryUseCase) UpdateBookCategory(ctx *gin.Context, bookCategoryInput *models.UpdateBookCategory) (*models.BookCategoryResponse, error) {
	bookCategory, err := u.bookCategoryRepo.FindByID(bookCategoryInput.ID)
	if err != nil {
		return nil, err
	}

	bookCategory.Name = bookCategoryInput.Name
	bookCategory.Description = bookCategoryInput.Description
	if bookCategoryInput.Image != "" {
		bookCategory.Image = bookCategoryInput.Image
	}

	updatedBookCategory, err := u.bookCategoryRepo.Update(bookCategory)
	if err != nil {
		return nil, err
	}

	return models.FilterBookCategoryRecord(updatedBookCategory), nil
}

func (u *BookCategoryUseCase) DeleteBookCategory(bookCategoryID uint) error {
	err := u.bookCategoryRepo.Delete(bookCategoryID)
	if err != nil {
		return err
	}
	return nil
}

func (u *BookCategoryUseCase) GetBookCategory(bookCategoryID uint) (*models.BookCategoryResponse, error) {
	bookCategory, err := u.bookCategoryRepo.FindByID(bookCategoryID)
	if err != nil {
		return nil, err
	}
	return models.FilterBookCategoryRecord(bookCategory), nil
}
