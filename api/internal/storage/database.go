package storage

import (
	"dou-survey/internal/logger"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"dou-survey/app/model/choiceModel"
	"dou-survey/app/model/employeeModel"
	"dou-survey/app/model/questionModel"
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/model/userModel"
	"dou-survey/app/model/voteModel"
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
		&questionModel.Question{},
		&userModel.User{},
		&employeeModel.Employee{},
		&surveyModel.Survey{},
	)

	return &DbStore{
		db,
	}
}
