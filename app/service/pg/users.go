package pg

import (
	"app-ink/app/model"
	"context"
)

var sqlUsers = createTableSQL("users",
	"id         SERIAL PRIMARY KEY",
	"created_at TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"updated_at TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"username   TEXT NOT NULL UNIQUE",
	"password   TEXT NOT NULL",
	"role       TEXT",
	"nickname   TEXT",
	"avatar     TEXT",
)

func (pg *Pg) UserFindByUsername(ctx context.Context, username string) (credentials model.Credentials, err error) {
	err = pg.Pool.QueryRow(ctx, `
		SELECT username, password 
		FROM users
		WHERE username = $1
	`, username).Scan(&credentials.Username, &credentials.Password)

	return credentials, err
}

func (pg *Pg) UserInsert(ctx context.Context, credentials model.Credentials) (err error) {
	_, err = pg.Pool.Exec(ctx, `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
	`, credentials.Username, credentials.Password)

	return err
}
