// Code generated by sqlc. DO NOT EDIT.
// source: timeschemas.sql

package db

import (
	"context"
	"encoding/json"
)

const createTimeSchema = `-- name: CreateTimeSchema :one
insert into timeschemas (name, items, updated_at)
values ($1, $2, now())
returning id, created_at, updated_at, name, items
`

type CreateTimeSchemaParams struct {
	Name  string          `json:"name"`
	Items json.RawMessage `json:"items"`
}

func (q *Queries) CreateTimeSchema(ctx context.Context, arg CreateTimeSchemaParams) (Timeschema, error) {
	row := q.db.QueryRowContext(ctx, createTimeSchema, arg.Name, arg.Items)
	var i Timeschema
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Items,
	)
	return i, err
}

const deleteTimeSchema = `-- name: DeleteTimeSchema :one
delete from timeschemas where id = $1
returning id, created_at, updated_at, name, items
`

func (q *Queries) DeleteTimeSchema(ctx context.Context, id int64) (Timeschema, error) {
	row := q.db.QueryRowContext(ctx, deleteTimeSchema, id)
	var i Timeschema
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Items,
	)
	return i, err
}

const getTimeSchema = `-- name: GetTimeSchema :one
SELECT id, created_at, updated_at, name, items FROM timeschemas
WHERE id = $1
`

func (q *Queries) GetTimeSchema(ctx context.Context, id int64) (Timeschema, error) {
	row := q.db.QueryRowContext(ctx, getTimeSchema, id)
	var i Timeschema
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Items,
	)
	return i, err
}

const searchTimeSchema = `-- name: SearchTimeSchema :many
select id, created_at, updated_at, name, items from timeschemas
where name like $1
`

func (q *Queries) SearchTimeSchema(ctx context.Context, name string) ([]Timeschema, error) {
	rows, err := q.db.QueryContext(ctx, searchTimeSchema, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Timeschema
	for rows.Next() {
		var i Timeschema
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Items,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTimeSchema = `-- name: UpdateTimeSchema :one
UPDATE timeschemas SET name = $2, items = $3, updated_at = now()
WHERE id = $1
returning id, created_at, updated_at, name, items
`

type UpdateTimeSchemaParams struct {
	ID    int64           `json:"id"`
	Name  string          `json:"name"`
	Items json.RawMessage `json:"items"`
}

func (q *Queries) UpdateTimeSchema(ctx context.Context, arg UpdateTimeSchemaParams) (Timeschema, error) {
	row := q.db.QueryRowContext(ctx, updateTimeSchema, arg.ID, arg.Name, arg.Items)
	var i Timeschema
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Items,
	)
	return i, err
}