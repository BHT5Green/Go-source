package server

import (
	"github.com/1rhino/clean_architecture/config"
	"github.com/1rhino/clean_architecture/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server struct
type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config *config.Config
}

// NewServer function
func NewServer(cfg *config.Config) *Server {
	return &Server{
		Router: gin.Default(),
		DB:     db.Init(cfg),
		Config: cfg,
	}
}

func (server *Server) Start() error {
	// CORS middleware
	server.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	SetupRoutes(server)

	return server.Router.Run(":" + server.Config.HTTP.Port)
}
