package pg

import (
	"app_ink/app/model"
	"context"
)

var sqlUsers = createTableSql("users",
	"id         SERIAL PRIMARY KEY",
	"created_at TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"updated_at TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"username   TEXT NOT NULL",
	"password   TEXT NOT NULL",
	"role       TEXT",
	"nickname   TEXT",
	"avatar     TEXT",
)

func (pg *Pg) UserFindByUsername(ctx context.Context, username string) (credentials model.Credentials, err error) {
	err = pg.Conn.QueryRow(ctx, `
		SELECT username, password 
		FROM users
		WHERE username = $1
	`, username).Scan(&credentials.Username, &credentials.Password)

	return credentials, err
}

func (pg *Pg) UserInsert(ctx context.Context, credentials model.Credentials) (err error) {
	_, err = pg.Conn.Exec(ctx, `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
	`, credentials.Username, credentials.Password)

	return err
}
