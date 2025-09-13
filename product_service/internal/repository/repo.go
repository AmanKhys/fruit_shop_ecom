package repository

import (
	"database/sql"
	"product_service/internal/infrastructure/db/sqlc"
)

type productRepo struct {
	q  *sqlc.Queries
	db *sql.DB
}

func NewProductRepo(q *sqlc.Queries, db *sql.DB) *productRepo {
	return &productRepo{}
}
