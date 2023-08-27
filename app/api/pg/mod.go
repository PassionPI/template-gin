package pg

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func New() {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqlDB, pgdialect.New())
}
