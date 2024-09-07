-- name: CreateToken :exec
INSERT INTO tokens (id, public_key, refresh_token, user_id) VALUES (?, ?, ?, ?);

-- name: GetTokenByUserID :one
SELECT * FROM tokens WHERE user_id = ? LIMIT 1;