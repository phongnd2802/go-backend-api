-- name: GetShopByEmail :one
SELECT * FROM shops WHERE shop_email = ? LIMIT 1;


-- name: Create :exec
INSERT INTO shops (id, shop_name, shop_email, shop_password)
VALUES (?, ?, ?, ?);

-- name: ActiveShopOTP :exec
UPDATE shops SET is_active = 1 WHERE shop_email = ?;

-- name: UpdatePassword :exec
UPDATE shops SET shop_password = ? WHERE shop_email = ?;