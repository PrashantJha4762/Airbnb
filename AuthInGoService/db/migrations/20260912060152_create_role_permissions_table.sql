-- +goose Up
create table if not exists role_permissions (
    id int primary key auto_increment,
    role_id int not null,
    permission_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    foreign key (role_id) references role(id) on delete cascade,
    foreign key (permission_id) references permissions(id) on delete cascade
);

insert into role_permissions (role_id,permission_id)
select 1,id from permissions;
-- giving all permissions to admin role

insert into role_permissions (role_id,permission_id)
select 2,id from permissions where name in ('user:read','role:read');
-- giving read permissions to user role




-- +goose Down
drop table if exists role_permissions;