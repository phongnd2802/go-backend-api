-- +goose Up
-- +goose StatementBegin
CREATE TABLE shops (
    id varchar(50) NOT NULL,
    shop_name varchar(100) NOT NULL,
    shop_email varchar(100) NOT NULL UNIQUE,
    shop_password varchar(255) NOT NULL,
    is_active tinyint(1) NOT NULL DEFAULT 0,
    verify tinyint(1) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists shops;
-- +goose StatementEnd
