package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Pg struct {
	Conn *pgx.Conn
}

func New(uri string) *Pg {
	background := context.Background()
	// 连接数据库
	Conn, err := pgx.Connect(background, uri)

	if err != nil {
		log.Fatal("Unable to connect to database: %v\n", err)
	}

	// 检查连接
	err = Conn.Ping(background)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Postgres Successfully connected!")

	ensureTableExists(background, Conn, []string{
		sqlUsers,
		createIndexSql("users", "username"),
		createIndexSql("users", "role"),
		createIndexSql("users", "nickname"),
		sqlTodo,
		createIndexSql("todo", "username"),
	})

	return &Pg{
		Conn,
	}
}

// 检查表是否存在并创建表
func ensureTableExists(ctx context.Context, db *pgx.Conn, sql []string) {
	chunk := ""
	for _, s := range sql {
		chunk += s
	}
	_, err := db.Exec(ctx, chunk)
	if err != nil {
		log.Fatalf("Error Exec Initial SQL: %s", err)
	}
}

func createIndexSql(tableName string, column string) string {
	return fmt.Sprintf(
		`CREATE INDEX IF NOT EXISTS idx_%s_%s ON %s(%s);`,
		tableName,
		column,
		tableName,
		column,
	)
}

func createTableSql(tableName string, kv ...string) string {
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
