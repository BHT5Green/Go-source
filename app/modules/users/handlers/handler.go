package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	user "github.com/1rhino/clean_architecture/app/modules/users/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandlers struct {
	userUseCase user.UseCase
}

func NewUserHandlers(userUseCase user.UseCase) *UserHandlers {
	return &UserHandlers{userUseCase: userUseCase}
}

// SignUp User
func (h *UserHandlers) SignUpUser(c *gin.Context) {
	payload := models.SignUpInput{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.userUseCase.SignUpUser(c, &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
}

// Login User
func (h *UserHandlers) LoginUser(c *gin.Context) {
	user := models.SignInInput{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedInUser, err := h.userUseCase.LoginUser(c, &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    loggedInUser.ID,
		"email": loggedInUser.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h *UserHandlers) GetUserProfile(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found in context"})
		return
	}

	userClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID, ok := userClaims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	profile, err := h.userUseCase.GetUserProfile(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *UserHandlers) LogoutUser(c *gin.Context) {
	token := strings.TrimSpace(c.GetHeader("Authorization"))
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	fmt.Printf("Logging out token: %s\n", token)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (h *UserHandlers) UpdateUser(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found in context"})
		return
	}

	userClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID, ok := userClaims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	var updatedUser models.UserUpdateInput

	if err := c.ShouldBind(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// file, err := c.FormFile("image")
	// if err == nil {
	// 	savePath := "assets/image/" + file.Filename
	// 	if err := c.SaveUploadedFile(file, savePath); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
	// 		return
	// 	}

	// 	updatedUser.Image = savePath
	// }

	file, err := c.FormFile("image")
	if err == nil {
		uploadedURL, uploadErr := middleware.HandleUploadImage(file)
		if uploadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}
		updatedUser.Image = uploadedURL
	}

	updatedUserResponse, err := h.userUseCase.UpdateUser(uint(userID), &updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, updatedUserResponse)
}

func (h *UserHandlers) DeleteUser(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found in context"})
		return
	}

	userClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID, ok := userClaims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	err := h.userUseCase.DeleteUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
