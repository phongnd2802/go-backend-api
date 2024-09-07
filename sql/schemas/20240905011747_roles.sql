-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
   id int not null AUTO_INCREMENT,
   role_name varchar(20) not null unique,
   role_note text not null,
   primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists roles;
-- +goose StatementEnd
