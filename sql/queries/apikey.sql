-- name: CreateAPIKey :exec
INSERT INTO api_keys (api_key, description) VALUES (?, ?);

-- name: GetAPIKey :one
SELECT * FROM api_keys WHERE api_key = ? LIMIT 1;