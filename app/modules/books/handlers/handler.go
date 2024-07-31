package handlers

import (
	"net/http"
	"strconv"

	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	book "github.com/1rhino/clean_architecture/app/modules/books/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandlers struct {
	bookUseCase book.UseCase
}

func NewBookHandlers(bookUseCase book.UseCase) *BookHandlers {
	return &BookHandlers{bookUseCase: bookUseCase}
}

// create a new book
func (h *BookHandlers) CreateBook(c *gin.Context) {
	var bookInput models.BookInput

	if err := c.ShouldBind(&bookInput); err != nil {
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
		bookInput.Image = uploadedURL
	}

	createdBook, err := h.bookUseCase.CreateBook(c, &bookInput, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdBook})
}

// get list of books
func (h *BookHandlers) GetAllBooks(c *gin.Context) {
	books, err := h.bookUseCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// get list books by userID
func (h *BookHandlers) GetBooks(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	books, err := h.bookUseCase.GetBooks(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// get book detail
func (h *BookHandlers) GetBookDetail(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book category ID"})
		return
	}

	getBook, err := h.bookUseCase.GetBook(uint(bookID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getBook)
}

// update book
func (h *BookHandlers) UpdateBook(c *gin.Context) {
	var bookInput models.UpdateBook

	if err := c.ShouldBind(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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
		bookInput.Image = uploadedURL
	}

	bookInput.ID = uint(bookID)
	updatedBook, err := h.bookUseCase.UpdateBook(c, &bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedBook})
}

// delete Book
func (h *BookHandlers) DeleteBook(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book category ID"})
		return
	}

	err = h.bookUseCase.DeleteBook(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BookCategory deleted successfully"})
}