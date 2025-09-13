package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	nethttp "net/http"
	"os"
	"user_service/handlers"
	"user_service/internal/config"
	"user_service/internal/delivery/http"
	"user_service/internal/infrastructure/db/sqlc"
	"user_service/internal/repository"
	"user_service/internal/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, using environment variables")
	}
}

func main() {
	dbConn, err := sql.Open("sqlite3", os.Getenv(config.DBDSN))
	if err != nil {
		log.Fatal(err)
	}
	q := sqlc.New()
	repo := repository.NewUserRepo(dbConn, q)
	uc := usecase.NewUserUsecase(repo)
	handler := handlers.NewUserHandler(uc)

	http.RegisterRoutes(handler)

	log.Info("Server running at :8080")
	err = nethttp.ListenAndServe(os.Getenv(config.ServerStartURL), nil)
	if err != nil {
		log.Fatal(err)
	}
}
