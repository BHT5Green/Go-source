package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(255)" json:"name"`
	Image       string       `gorm:"type:varchar(255)" json:"image"`
	Author      string       `gorm:"type:varchar(255)" json:"author"`
	PublicDate  time.Time    `json:"public_date"`
	Description string       `gorm:"type:varchar(255)" json:"description"`
	CategoryID  uint         `json:"category_id"`
	Category    BookCategory `json:"category"`
	UserID      uint         `json:"user_id"`
	User        User         `json:"user"`
}

func (Book) TableName() string {
	return "books"
}

type BookInput struct {
	Name        string    `form:"name" json:"name" binding:"required"`
	Image       string    `file:"image" json:"image"`
	Author      string    `form:"author" json:"author"`
	CategoryID  uint      `form:"category_id" json:"category_id"`
	PublicDate  time.Time `form:"public_date" json:"public_date" time_format:"02-01-2006"`
	Description string    `form:"description" json:"description"`
}

type UpdateBook struct {
	ID          uint      `form:"id" json:"id"`
	Name        string    `form:"name" json:"name" binding:"required"`
	Image       string    `file:"image" json:"image"`
	Author      string    `form:"author" json:"author"`
	CategoryID  uint      `form:"category_id" json:"category_id"`
	PublicDate  time.Time `form:"public_date" json:"public_date" time_format:"02-01-2006"`
	Description string    `form:"description" json:"description"`
}

type BookResponse struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Author      string    `json:"author"`
	CategoryID  uint      `json:"category_id"`
	UserID      uint      `json:"user_id"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	PublicDate  time.Time `json:"public_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FilterBookRecord(books *Book) *BookResponse {
	return &BookResponse{
		ID:          books.ID,
		Name:        books.Name,
		Author:      books.Author,
		CategoryID:  books.CategoryID,
		UserID:      books.UserID,
		PublicDate:  books.PublicDate,
		Description: books.Description,
		Image:       books.Image,
		CreatedAt:   books.CreatedAt,
		UpdatedAt:   books.UpdatedAt,
	}
}
