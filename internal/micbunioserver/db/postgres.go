package db

import (
	"fmt"
	"github.com/MicBun/micbun.io-backend/internal/micbunioserver/blog/model"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var gormInstance *gorm.DB

func NewPostgres() (*gorm.DB, error) {
	if gormInstance == nil {
		dsn := os.Getenv("PG_INTERNAL_DATABASE")
		if dsn == "" {
			dsn = os.Getenv("PG_EXTERNAL_DATABASE")
		}
		if dsn == "" {
			host := os.Getenv("PG_HOST")
			port := os.Getenv("PG_PORT")
			user := os.Getenv("PG_USER")
			password := os.Getenv("PG_PASSWORD")
			database := os.Getenv("PG_DATABASE")

			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s",
				host,
				port,
				user,
				password,
				database,
			)
		}

		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		gormInstance = gormDB
	}

	if err := gormInstance.AutoMigrate(&model.Blog{}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := gormInstance.AutoMigrate(&model.Comment{}); err != nil {
		return nil, errors.WithStack(err)
	}

	return gormInstance, nil
}
