// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: token.sql

package database

import (
	"context"
)

const createToken = `-- name: CreateToken :exec
INSERT INTO tokens (id, public_key, refresh_token, user_id) VALUES (?, ?, ?, ?)
`

type CreateTokenParams struct {
	ID           string
	PublicKey    string
	RefreshToken string
	UserID       string
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) error {
	_, err := q.db.ExecContext(ctx, createToken,
		arg.ID,
		arg.PublicKey,
		arg.RefreshToken,
		arg.UserID,
	)
	return err
}

const getTokenByUserID = `-- name: GetTokenByUserID :one
SELECT id, public_key, refresh_token, refresh_token_used, user_id, created_at, updated_at FROM tokens WHERE user_id = ? LIMIT 1
`

func (q *Queries) GetTokenByUserID(ctx context.Context, userID string) (Token, error) {
	row := q.db.QueryRowContext(ctx, getTokenByUserID, userID)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.PublicKey,
		&i.RefreshToken,
		&i.RefreshTokenUsed,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
