-- name: CreateRole :exec
INSERT INTO roles (role_name, role_note) VALUES (?, ?);
