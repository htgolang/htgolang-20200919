create table if not exists user (
    id bigint primary key auto_increment,
    name varchar(32) not null default '',
    sex boolean not null default true,
    addr text not null,
	tel varchar(32) not null default '',
	birthday date not null default '1990-01-01',
	password varchar(32) not null default ''
) engine=innodb default charset utf8mb4;