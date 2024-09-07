-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens (
    id varchar(50) not null,
    public_key text not null,
    refresh_token text not null,
    refresh_token_used text,
    user_id varchar(50) not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (id),
    foreign key (user_id) references shops(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
