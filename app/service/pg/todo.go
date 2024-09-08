package pg

import (
	"app-ink/app/model"
	"context"
)

var sqlTodo = createTableSQL("todo",
	"id          BIGSERIAL PRIMARY KEY",
	"created_at  TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"updated_at  TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"deadline    TIMESTAMPTZ",
	"username    TEXT NOT NULL",
	"title       TEXT NOT NULL",
	"done        BOOLEAN DEFAULT false",
	"description TEXT",
)

func (pg *Pg) TodoInsert(ctx context.Context, username string, todo *model.TodoCreateItem) (err error) {
	_, err = pg.Pool.Exec(ctx, `
		INSERT INTO todo (username, title, description)
		VALUES ($1, $2, $3)
	`,
		username,
		todo.Title,
		// todo.DeadLine,
		todo.Description,
	)

	return err
}

type TodoFindByUsernameParams struct {
	Username string
	*model.Pagination
}

func (pg *Pg) TodoFindByUsernameParamsCreate() *TodoFindByUsernameParams {
	return &TodoFindByUsernameParams{}
}
func (param *TodoFindByUsernameParams) SetUsername(username string) *TodoFindByUsernameParams {
	param.Username = username
	return param
}
func (param *TodoFindByUsernameParams) SetPagination(pagination *model.Pagination) *TodoFindByUsernameParams {
	param.Pagination = pagination
	return param
}

func (pg *Pg) TodoFindByUsername(
	ctx context.Context,
	param *TodoFindByUsernameParams,
) (todos []model.TodoScanItem, err error) {
	rows, err := pg.Pool.Query(ctx, `
		SELECT id, title, done, updated_at, deadline, description
		FROM todo
		WHERE username = $1
		ORDER BY id DESC
		OFFSET $2
		LIMIT $3
	`,
		param.Username,
		param.Pagination.Page*param.Pagination.Size,
		param.Pagination.Size,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo model.TodoScanItem
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Done,
			&todo.UpdatedAt,
			&todo.DeadLine,
			&todo.Description,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (pg *Pg) TodoFindById(ctx context.Context, id int) (todo model.TodoScanItem, err error) {
	err = pg.Pool.QueryRow(ctx, `
		SELECT id, title, done, deadline, description
		FROM todo
		WHERE id = $1
	`, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Done,
		&todo.DeadLine,
		&todo.Description,
	)

	return todo, err
}

func (pg *Pg) TodoDeleteById(ctx context.Context, id int) (err error) {
	_, err = pg.Pool.Exec(ctx, `
		DELETE FROM todo
		WHERE id = $1
	`, id)

	return err
}

func (pg *Pg) TodoUpdateById(ctx context.Context, todo *model.TodoUpdateItem) (err error) {
	_, err = pg.Pool.Exec(ctx, `
		UPDATE todo
		SET 
			updated_at = NOW(),
			-- title = $1, 
			done = $1
			-- description = $3,
			-- deadline = $4
		WHERE id = $2
	`,
		// todo.Title,
		todo.Done,
		// todo.DeadLine,
		// todo.Description,
		todo.ID,
	)

	return err
}
