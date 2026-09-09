-- +goose Up
create table users(
id serial primary key,
username varchar(255) not null,
email varchar(255) not null,
password varchar(255) not null,
created_at timestamp not null default current_timestamp,
updated_at timestamp not null default current_timestamp
);

-- +goose Down
drop table users;
