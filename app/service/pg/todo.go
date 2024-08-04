package pg

import (
	"app_ink/app/model"
	"context"
)

var sqlTodo = createTableSql("todo",
	"id SERIAL   PRIMARY KEY",
	"created_at  TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"updated_at  TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')",
	"dead_line   TIMESTAMPTZ",
	"username    TEXT NOT NULL",
	"title       TEXT NOT NULL",
	"done        BOOLEAN DEFAULT false",
	"description TEXT",
)

func (pg *Pg) TodoInsert(ctx context.Context, todo model.TodoCreateItem) (err error) {
	_, err = pg.DB.Exec(ctx, `
		INSERT INTO todo (username, title, description, dead_line)
		VALUES ($1, $2, $3)
	`, todo.Username, todo.Title, todo.Description, todo.DeadLine)

	return err
}

type TodoFindByUsernameParams struct {
	Username string
	model.Pagination
}

func TodoFindByUsernameParamsCreate() TodoFindByUsernameParams {
	return TodoFindByUsernameParams{}
}
func (param *TodoFindByUsernameParams) SetUsername(username string) *TodoFindByUsernameParams {
	param.Username = username
	return param
}
func (param *TodoFindByUsernameParams) SetPagination(pagination model.Pagination) *TodoFindByUsernameParams {
	param.Pagination = pagination
	return param
}

func (pg *Pg) TodoFindByUsername(
	ctx context.Context,
	param TodoFindByUsernameParams,
) (todos []model.TodoCreateItem, err error) {
	rows, err := pg.DB.Query(ctx, `
		SELECT username, title, description, dead_line
		FROM todo
		WHERE username = $1
		OFFSET $2
		LIMIT $3
		ORDER BY created_at DESC
	`,
		param.Username,
		param.Pagination.Page*param.Pagination.Size,
		param.Pagination.Size,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo model.TodoCreateItem
		err = rows.Scan(
			&todo.Username,
			&todo.Title,
			&todo.Description,
			&todo.DeadLine,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (pg *Pg) TodoFindById(ctx context.Context, id int) (todo model.TodoCreateItem, err error) {
	err = pg.DB.QueryRow(ctx, `
		SELECT username, title, description, dead_line
		FROM todo
		WHERE id = $1
	`, id).Scan(
		&todo.Username,
		&todo.Title,
		&todo.Description,
		&todo.DeadLine,
	)

	return todo, err
}

func (pg *Pg) TodoDeleteById(ctx context.Context, id int) (err error) {
	_, err = pg.DB.Exec(ctx, `
		DELETE FROM todo
		WHERE id = $1
	`, id)

	return err
}

func (pg *Pg) TodoUpdateById(ctx context.Context, id int, todo model.TodoCreateItem) (err error) {
	_, err = pg.DB.Exec(ctx, `
		UPDATE todo
		SET title = $1, description = $2, dead_line = $3
		WHERE id = $4
	`, todo.Title, todo.Description, todo.DeadLine, id)

	return err
}
