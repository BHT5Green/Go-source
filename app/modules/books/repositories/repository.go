package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) (*models.Book, error)
	FindAll() ([]*models.Book, error)
	FindByUserID(userID uint) ([]*models.Book, error)
	FindByID(id uint) (*models.Book, error)
	Update(book *models.Book) (*models.Book, error)
	Delete(id uint) error
}

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepository {
	return &BookRepo{DB: db}
}

func (r *BookRepo) Create(book *models.Book) (*models.Book, error) {
	if err := r.DB.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepo) FindAll() ([]*models.Book, error) {
	var books []*models.Book
	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) FindByUserID(userID uint) ([]*models.Book, error) {
	var books []*models.Book
	if err := r.DB.Where("user_id = ?", userID).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := r.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepo) Update(book *models.Book) (*models.Book, error) {
	if err := r.DB.Save(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}
