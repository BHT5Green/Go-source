package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
	"gorm.io/gorm"
)

type BookCategoryRepository interface {
	Create(bookCategory *models.BookCategory) (*models.BookCategory, error)
	FindByUserID(userID uint) ([]*models.BookCategory, error)
	FindAll() ([]*models.BookCategory, error)
	FindByID(id uint) (*models.BookCategory, error)
	Update(bookCategory *models.BookCategory) (*models.BookCategory, error)
	Delete(bookCategoryID uint) error
}

type BookCategoryRepo struct {
	DB *gorm.DB
}

func NewBookCategoryRepo(db *gorm.DB) BookCategoryRepository {
	return &BookCategoryRepo{DB: db}
}

func (r *BookCategoryRepo) Create(bookCategory *models.BookCategory) (*models.BookCategory, error) {
	if err := r.DB.Create(bookCategory).Error; err != nil {
		return nil, err
	}
	return bookCategory, nil
}

func (r *BookCategoryRepo) FindByUserID(userID uint) ([]*models.BookCategory, error) {
	var bookCategories []*models.BookCategory
	if err := r.DB.Where("user_id = ?", userID).Find(&bookCategories).Error; err != nil {
		return nil, err
	}
	return bookCategories, nil
}

func (r *BookCategoryRepo) FindAll() ([]*models.BookCategory, error) {
	var bookCategories []*models.BookCategory
	if err := r.DB.Find(&bookCategories).Error; err != nil {
		return nil, err
	}
	return bookCategories, nil
}

func (r *BookCategoryRepo) FindByID(id uint) (*models.BookCategory, error) {
	var bookCategory models.BookCategory
	if err := r.DB.First(&bookCategory, id).Error; err != nil {
		return nil, err
	}
	return &bookCategory, nil
}

func (r *BookCategoryRepo) Update(bookCategory *models.BookCategory) (*models.BookCategory, error) {
	if err := r.DB.Save(bookCategory).Error; err != nil {
		return nil, err
	}
	return bookCategory, nil
}

func (r *BookCategoryRepo) Delete(id uint) error {
	return r.DB.Delete(&models.BookCategory{}, id).Error
}
