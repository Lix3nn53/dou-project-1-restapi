package storage

import (
	"fmt"
	"goa-golang/internal/logger"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"goa-golang/app/model/choiceModel"
	"goa-golang/app/model/employeeModel"
	"goa-golang/app/model/surveyModel"
	"goa-golang/app/model/userModel"
	"goa-golang/app/model/voteModel"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
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

	sqlDB, err := db.DB()

	if err != nil {
		logger.Fatalf(err.Error())
		return nil
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime * time.Second)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	retryCount := 30
	for {
		err := sqlDB.Ping()
		if err != nil {
			if retryCount == 0 {
				logger.Fatalf("Not able to establish connection to database")
			}
			logger.Infof(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", retryCount))
			retryCount--
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}
	if err = sqlDB.Ping(); err != nil {
		return nil
	}

	db.AutoMigrate(
		&choiceModel.Choice{},
		&voteModel.Vote{},
		&userModel.User{},
		&employeeModel.Employee{},
		&surveyModel.Survey{},
	)

	return &DbStore{
		db,
	}
}
