// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: supers.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getSuperByUserId = `-- name: GetSuperByUserId :one
SELECT supers.id, user_id, username, password, name FROM users
INNER JOIN supers ON supers.user_id = users.id
WHERE user_id = $1
`

type GetSuperByUserIdRow struct {
	ID       pgtype.UUID
	UserID   pgtype.UUID
	Username string
	Password string
	Name     string
}

func (q *Queries) GetSuperByUserId(ctx context.Context, userID pgtype.UUID) (GetSuperByUserIdRow, error) {
	row := q.db.QueryRow(ctx, getSuperByUserId, userID)
	var i GetSuperByUserIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Name,
	)
	return i, err
}

const getSuperByUsername = `-- name: GetSuperByUsername :one
SELECT supers.id, user_id, username, password, name FROM users
INNER JOIN supers ON supers.user_id = users.id
WHERE username = $1
`

type GetSuperByUsernameRow struct {
	ID       pgtype.UUID
	UserID   pgtype.UUID
	Username string
	Password string
	Name     string
}

func (q *Queries) GetSuperByUsername(ctx context.Context, username string) (GetSuperByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getSuperByUsername, username)
	var i GetSuperByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Name,
	)
	return i, err
}

const getSupers = `-- name: GetSupers :many
SELECT supers.id, user_id, username, password, name FROM supers
INNER JOIN users ON supers.user_id = users.id
`

type GetSupersRow struct {
	ID       pgtype.UUID
	UserID   pgtype.UUID
	Username string
	Password string
	Name     string
}

func (q *Queries) GetSupers(ctx context.Context) ([]GetSupersRow, error) {
	rows, err := q.db.Query(ctx, getSupers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSupersRow
	for rows.Next() {
		var i GetSupersRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Username,
			&i.Password,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSuper = `-- name: UpdateSuper :exec
UPDATE users SET username = $2, name = $3 WHERE id = $1
`

type UpdateSuperParams struct {
	ID       pgtype.UUID
	Username string
	Name     string
}

func (q *Queries) UpdateSuper(ctx context.Context, arg UpdateSuperParams) error {
	_, err := q.db.Exec(ctx, updateSuper, arg.ID, arg.Username, arg.Name)
	return err
}