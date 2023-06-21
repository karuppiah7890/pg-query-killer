package postgres

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Client struct {
	db *bun.DB
}

func NewClient(uri string) *Client {
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
	db := bun.NewDB(pgdb, pgdialect.New())
	return &Client{
		db: db,
	}
}
