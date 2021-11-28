package storage

import (
	"goa-golang/internal/logger"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DbStore ...
type DbStore struct {
	*gorm.DB
}

// InitializeDB Opening a storage and save the reference to `Database` struct.
func InitializeDB(logger logger.Logger) *DbStore {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		logger.Fatalf(err.Error())
		return nil
	}

	return &DbStore{
		db,
	}
}
