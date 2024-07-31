package handlers

import (
	"net/http"
	"strconv"

	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	book_category "github.com/1rhino/clean_architecture/app/modules/book_category/usecase"
	"github.com/gin-gonic/gin"
)

type BookCategoryHandlers struct {
	bookUseCase book_category.UseCase
}

func NewBookCategoryHandlers(bookUseCase book_category.UseCase) *BookCategoryHandlers {
	return &BookCategoryHandlers{bookUseCase: bookUseCase}
}

// create a new book category
func (h *BookCategoryHandlers) CreateBookCategory(c *gin.Context) {
	var bookCategoryInput models.BookCategoryInput

	if err := c.ShouldBind(&bookCategoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err == nil {
		uploadedURL, uploadErr := middleware.HandleUploadImage(file)
		if uploadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}
		bookCategoryInput.Image = uploadedURL
	}

	createdBookCategory, err := h.bookUseCase.CreateBookCategory(c, &bookCategoryInput, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdBookCategory})
}

// get list of book categories by user ID
func (h *BookCategoryHandlers) GetBookCategories(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	bookCategories, err := h.bookUseCase.GetBookCategories(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookCategories})
}

// get all book categories
func (h *BookCategoryHandlers) GetAllBookCategories(c *gin.Context) {
	bookCategories, err := h.bookUseCase.GetAllBookCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookCategories})
}

// get book category detail
func (h *BookCategoryHandlers) GetBookCategoryDetail(c *gin.Context) {
	bookCategoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book category ID"})
		return
	}

	getBookCategory, err := h.bookUseCase.GetBookCategory(uint(bookCategoryID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getBookCategory)
}

// update a book category
func (h *BookCategoryHandlers) UpdateBookCategory(c *gin.Context) {
	var bookCategoryInput models.UpdateBookCategory

	if err := c.ShouldBind(&bookCategoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookCategoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book category ID"})
		return
	}

	file, err := c.FormFile("image")
	if err == nil {
		uploadedURL, uploadErr := middleware.HandleUploadImage(file)
		if uploadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}
		bookCategoryInput.Image = uploadedURL
	}

	bookCategoryInput.ID = uint(bookCategoryID)
	updatedBookCategory, err := h.bookUseCase.UpdateBookCategory(c, &bookCategoryInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedBookCategory})
}

// delete book category
func (h *BookCategoryHandlers) DeleteBookCategory(c *gin.Context) {
	bookCategoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book category ID"})
		return
	}

	err = h.bookUseCase.DeleteBookCategory(uint(bookCategoryID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BookCategory deleted successfully"})
}
