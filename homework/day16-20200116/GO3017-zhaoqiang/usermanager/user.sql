create database if not exists usermanager character set utf8mb4;
use usermanager;
create table if not exists user (
    id bigint primary key auto_increment,
    name varchar(64) not null default '',
    sex boolean not null default true,
    addr varchar(1024) not null default '',
    tel varchar(11) not null default '',
    brithday date,
    passwd text,
    create_at datetime
) engine=innodb default charset=utf8mb4;

alter table user add column if not exists remark text not null default '';
