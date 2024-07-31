package db

import (
	"fmt"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/1rhino/clean_architecture/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.BookCategory{},
	)

	if err != nil {
		panic(err.Error())
	}

	return db
}
