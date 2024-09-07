-- +goose Up
-- +goose StatementBegin
CREATE TABLE shop_roles (
    shop_id varchar(50) not null ,
    role_id int not null ,
    primary key  (shop_id, role_id),
    foreign key (shop_id) references shops(id),
    foreign key (role_id) references roles(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists shop_roles;
-- +goose StatementEnd
