package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	nethttp "net/http"
	"os"
	"product_service/internal/config"
	"product_service/internal/delivery/http"
	"product_service/internal/infrastructure/db/sqlc"
	"product_service/internal/repository"
	"product_service/internal/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, using environment variables")
	}
}

func main() {
	// setup db
	dbConn, err := sql.Open("sqlite3", os.Getenv(config.DBDSN))
	if err != nil {
		log.Fatal(err)
	}

	// wire layers/ dependency injection
	q := sqlc.New()
	repo := repository.NewProductRepo(dbConn, q)
	uc := usecase.NewProductUsecase(repo)
	handler := http.NewProductHandler(uc)

	http.RegisterRoutes(handler)

	url := os.Getenv(config.ServerStartURL)
	log.Info("server started running on:" + url)
	err = nethttp.ListenAndServe(url, nil)
	if err != nil {
		log.Fatal(err)
	}
}
