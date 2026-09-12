-- +goose Up
Create table if not exists role (
    id int primary key auto_increment,
    name varchar(255) not null unique,
    description varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);
insert into role (name, description) values
 ('admin', 'Administrator role with full access'),
 ('user', 'Regular user role with limited access'),
 ('manager', 'Manager role with elevated access'),
 ('guest', 'Guest role with minimal access');

-- +goose Down
Drop table if exists role;