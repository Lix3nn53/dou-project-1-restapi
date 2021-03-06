package main

import (
	"dou-survey/internal/logger"
	"dou-survey/internal/route"
	"dou-survey/internal/storage"
	"errors"

	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var config string

func setupRouter() (*gin.Engine, logger.Logger) {
	flag.StringVar(&config, "env", "pro.env", "Environment name")
	flag.Parse()

	logger := logger.NewAPILogger()
	logger.InitLogger()

	logger.Info("Hello World!")

	if err := godotenv.Load(config); err != nil {
		logger.Fatalf(err.Error())
		os.Exit(1)
	}

	db := storage.InitializeDB(logger)
	if db == nil {
		error := errors.New("could not initialize database")
		logger.Fatal(error)
		os.Exit(1)
	}

	dbCache := storage.InitializeCache()
	router := route.Setup(db, dbCache, logger)

	return router, logger
}

func main() {
	router, _ := setupRouter()

	router.Run(":" + os.Getenv("APP_PORT"))
}
