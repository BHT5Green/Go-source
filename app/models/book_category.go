package models

import (
	"time"

	"gorm.io/gorm"
)

type BookCategory struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Books       []Book `json:"books" gorm:"foreignKey:CategoryID"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"user"`
}

func (BookCategory) TableName() string {
	return "book_categories"
}

type BookCategoryInput struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Image       string `file:"image" json:"image"`
	Description string `form:"description" json:"description"`
}

type UpdateBookCategory struct {
	ID          uint   `form:"id" json:"id"`
	Name        string `form:"name" json:"name" binding:"required"`
	Image       string `file:"image" json:"image"`
	Description string `form:"description" json:"description"`
}

type BookCategoryResponse struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `form:"description" json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FilterBookCategoryRecord(book_categories *BookCategory) *BookCategoryResponse {
	return &BookCategoryResponse{
		ID:          book_categories.ID,
		Name:        book_categories.Name,
		Description: book_categories.Description,
		Image:       book_categories.Image,
		CreatedAt:   book_categories.CreatedAt,
		UpdatedAt:   book_categories.UpdatedAt,
	}
}
