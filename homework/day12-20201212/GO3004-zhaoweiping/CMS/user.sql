-- mysql:5.7.30

create database if not exists user default charset utf8mb4  COLLATE utf8mb4_general_ci;

create table if not exists user (
    id bigint primary key auto_increment,
    name varchar(32) not null default '',
    sex boolean not null default true,
    addr text,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp
) engine=innodb default charset utf8mb4;

