-- name: CreateToken :exec
INSERT INTO tokens (id, public_key, refresh_token, user_id) VALUES (?, ?, ?, ?);

