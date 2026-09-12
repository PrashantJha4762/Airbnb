-- +goose Up
create table if not exists permissions (
    id int primary key auto_increment,
    name varchar(255) not null unique,
    description varchar(255) not null,
    resource varchar(255) not null,
    action varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);
insert into permissions (name, description, resource, action) values
 ('user:read', 'Permission to read user information', 'user', 'read'),
 ('user:write', 'Permission to write user information', 'user', 'write'),
 ('user:delete', 'Permission to delete user information', 'user', 'delete'),
 ('role:read', 'Permission to read role information', 'role', 'read'),
 ('role:write', 'Permission to write role information', 'role', 'write'),
 ('role:delete', 'Permission to delete role information', 'role', 'delete'),
 ('role:manage', 'Permission to manage role information', 'role', 'manage'),
 ('permission:manage', 'Permission to manage permission information', 'permission', 'manage');

-- +goose Down
Drop table if exists permissions;