package server

import (
	"github.com/1rhino/clean_architecture/app/middleware"
	handlerBookCategory "github.com/1rhino/clean_architecture/app/modules/book_category/handlers"
	repositoryBookCategory "github.com/1rhino/clean_architecture/app/modules/book_category/repositories"
	bookCategoryUseCase "github.com/1rhino/clean_architecture/app/modules/book_category/usecase"
	handlerBook "github.com/1rhino/clean_architecture/app/modules/books/handlers"
	repositoryBook "github.com/1rhino/clean_architecture/app/modules/books/repositories"
	bookUseCase "github.com/1rhino/clean_architecture/app/modules/books/usecase"
	handlerUser "github.com/1rhino/clean_architecture/app/modules/users/handlers"
	repositoryUser "github.com/1rhino/clean_architecture/app/modules/users/repositories"
	userUseCase "github.com/1rhino/clean_architecture/app/modules/users/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *Server) {
	// routes
	r := gin.Default()
	api := r.Group("/api/v1")

	// User
	userRepo := repositoryUser.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := handlerUser.NewUserHandlers(userUseCase)
	authMiddleware := middleware.AuthMiddleware("your_secret_key")

	user := api.Group("/users")
	user.POST("/signup", userHandler.SignUpUser)
	user.POST("/login", userHandler.LoginUser)
	user.GET("/profile", authMiddleware, userHandler.GetUserProfile)
	user.DELETE("/logout", authMiddleware, userHandler.LogoutUser)
	user.PATCH("/update", authMiddleware, userHandler.UpdateUser)
	user.DELETE("/delete", authMiddleware, userHandler.DeleteUser)

	// Books
	bookRepo := repositoryBook.NewBookRepo(server.DB)
	bookUseCase := bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := handlerBook.NewBookHandlers(bookUseCase)

	books := api.Group("/books")
	books.POST("/create", authMiddleware, bookHandler.CreateBook)
	books.GET("/lists", authMiddleware, bookHandler.GetAllBooks)
	books.GET("/user/lists", authMiddleware, bookHandler.GetBooks)
	books.GET("/detail/:id", authMiddleware, bookHandler.GetBookDetail)
	books.PATCH("/update/:id", authMiddleware, bookHandler.UpdateBook)
	books.DELETE("/delete/:id", authMiddleware, bookHandler.DeleteBook)

	// Book Category
	bookCategoryRepo := repositoryBookCategory.NewBookCategoryRepo(server.DB)
	bookCategoryUseCase := bookCategoryUseCase.NewBookCategoryUseCase(bookCategoryRepo)
	bookCategoryHandler := handlerBookCategory.NewBookCategoryHandlers(bookCategoryUseCase)

	bookCategories := api.Group("/book_categories")
	bookCategories.POST("/create", authMiddleware, bookCategoryHandler.CreateBookCategory)
	bookCategories.GET("/user/lists", authMiddleware, bookCategoryHandler.GetBookCategories)
	bookCategories.GET("/lists", authMiddleware, bookCategoryHandler.GetAllBookCategories)
	bookCategories.GET("/detail/:id", authMiddleware, bookCategoryHandler.GetBookCategoryDetail)
	bookCategories.PATCH("/update/:id", authMiddleware, bookCategoryHandler.UpdateBookCategory)
	bookCategories.DELETE("/delete/:id", authMiddleware, bookCategoryHandler.DeleteBookCategory)

	server.Router = r
}
