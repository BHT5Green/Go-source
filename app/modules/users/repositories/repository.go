package repository

import (
	"fmt"

	"github.com/1rhino/clean_architecture/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CheckEmailExisting(email string) bool
	CreateUser(data *models.SignUpInput) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{DB: db}
}

func (r UserRepo) CreateUser(data *models.SignUpInput) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	var user = &models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hashedPassword),
	}

	result := r.DB.Table(models.User{}.TableName()).Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r UserRepo) CheckEmailExisting(email string) bool {
	var user *models.User

	result := r.DB.Table(models.User{}.TableName()).Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println("err: ", result.Error)

		return false
	}

	return result.RowsAffected > 0
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepo) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
