package repository

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type BunRepository struct {
	db *bun.DB
}

func NewBunRepository(conn *sql.DB) *BunRepository {
	return &BunRepository{
		db: bun.NewDB(conn, pgdialect.New()),
	}
}
