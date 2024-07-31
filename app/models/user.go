package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string         `gorm:"type:varchar(255)" json:"name"`
	Email        string         `gorm:"type:varchar(255)" json:"email"`
	Password     string         `gorm:"type:varchar(255)" json:"password"`
	Image        string         `gorm:"type:varchar(255)" json:"image"`
	Books        []Book         `json:"books" gorm:"foreignKey:UserID"`
	BookCategory []BookCategory `json:"book_categories" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUser struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `json:"password"  validate:"required"`
}

type UserProfile struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateInput struct {
	Name  string `form:"name" json:"name" validate:"required"`
	Email string `form:"email" json:"email" validate:"required"`
	Image string `file:"image" json:"image"`
}

func FilterUserRecord(user *User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
