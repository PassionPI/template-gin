package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pg struct {
	Pool *pgxpool.Pool
}

func New(uri string) *Pg {
	background := context.Background()

	Pool, err := pgxpool.New(background, uri)

	if err != nil {
		log.Fatal("Unable to connect to database: %v\n", err)
	}

	// 检查连接
	err = Pool.Ping(background)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Postgres Successfully connected!")

	ensureTableExists(background, Pool, []string{
		sqlUsers,
		createIndexSQL("users", "username"),
		createIndexSQL("users", "role"),
		createIndexSQL("users", "nickname"),
		sqlTodo,
		createIndexSQL("todo", "username"),
	})

	return &Pg{
		Pool,
	}
}

// 检查表是否存在并创建表
func ensureTableExists(ctx context.Context, db *pgxpool.Pool, sql []string) {
	chunk := ""
	for _, s := range sql {
		chunk += s
	}
	_, err := db.Exec(ctx, chunk)
	if err != nil {
		log.Fatalf("Error Exec Initial SQL: %s", err)
	}
}

func createIndexSQL(tableName string, column string) string {
	return fmt.Sprintf(
		`CREATE INDEX IF NOT EXISTS idx_%s_%s ON %s(%s);`,
		tableName,
		column,
		tableName,
		column,
	)
}

func createTableSQL(tableName string, kv ...string) string {
	start := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (`,
		tableName,
	)

	columns := ""
	last := len(kv) - 1

	for i := 0; i < len(kv); i++ {
		columns += kv[i]
		if i != last {
			columns += ","
		}
	}

	end := ");"

	return start + columns + end
}
