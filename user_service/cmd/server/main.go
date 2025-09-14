package main

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	nethttp "net/http"
	"os"
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
	// setup db
	dbConn, err := sql.Open("sqlite3", os.Getenv(config.DBDSN))
	if err != nil {
		log.Fatal(err)
	}

	// wire layers/ dependency injection
	q := sqlc.New()
	repo := repository.NewUserRepo(dbConn, q)
	uc := usecase.NewUserUsecase(repo)
	handler := http.NewUserHandler(uc)

	// bootstrap admin
	ctx := context.TODO()
	_, err = uc.EnsureAdminExists(ctx, os.Getenv(config.AdminEmail), os.Getenv(config.AdminPassword))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("admin created/exists")
	}
	http.RegisterRoutes(handler)

	log.Info("Server running at :8080")
	err = nethttp.ListenAndServe(os.Getenv(config.ServerStartURL), nil)
	if err != nil {
		log.Fatal(err)
	}
}
