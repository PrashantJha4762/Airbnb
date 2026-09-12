-- +goose Up
create table if not exists user_roles (
    id int primary key auto_increment,
    role_id int not null,
    users_id bigint unsigned not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    foreign key (role_id) references role(id),
    foreign key (users_id) references users(id)
);



-- +goose Down
drop table if exists user_roles;