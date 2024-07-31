package usecase

import (
	"errors"

	"github.com/1rhino/clean_architecture/app/models"
	users "github.com/1rhino/clean_architecture/app/modules/users/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	SignUpUser(ctx *gin.Context, payload *models.SignUpInput) (*models.UserResponse, error)
	LoginUser(ctx *gin.Context, user *models.SignInInput) (*models.User, error)
	GetUserProfile(userID uint) (*models.UserResponse, error)
	UpdateUser(userID uint, updatedUser *models.UserUpdateInput) (*models.UserResponse, error)
	DeleteUser(userID uint) error
}

type UserUseCase struct {
	userRepo users.UserRepoInterface
}

func NewUserUseCase(userRepo users.UserRepoInterface) UseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u UserUseCase) SignUpUser(ctx *gin.Context, payload *models.SignUpInput) (*models.UserResponse, error) {
	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	if u.userRepo.CheckEmailExisting(payload.Email) {
		return nil, errors.New("email existing, please choose another email")
	}

	createdUser, err := u.userRepo.CreateUser(payload)
	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(createdUser), nil
}

func (u *UserUseCase) LoginUser(ctx *gin.Context, user *models.SignInInput) (*models.User, error) {
	foundUser, err := u.userRepo.GetByEmail(user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return foundUser, nil
}

func (u *UserUseCase) GetUserProfile(userID uint) (*models.UserResponse, error) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return models.FilterUserRecord(user), nil
}

func (u *UserUseCase) UpdateUser(userID uint, updatedUser *models.UserUpdateInput) (*models.UserResponse, error) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Image != "" {
		user.Image = updatedUser.Image
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(user), nil
}

func (u *UserUseCase) DeleteUser(userID uint) error {
	err := u.userRepo.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}
