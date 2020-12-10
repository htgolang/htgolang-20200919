1. 复习
    重定向: http.Redirect(w, r, url, 302)
    3XX Response Header Localtion: url
2. 作业
3. 数据库
    关系型数据库
        MySQL
        PostgreSQL
        access
        Oracle
        SQLServer
        SQLlite
    非关系型数据库
        K-V: redis, etcd
        文档的: mongodb, es
        面向列: cassandra, hbase/hdfs
        ...

    MySQL mariadb > 5.5
    SQL
        DCL
            账号: 用户名@访问范围
                ip地址, cidr格式
                %
            create user 'golang'@'%' identified by 'golang@2020';
            grant all privileges on user.* to 'golang'@'%';
        DDL
            库:
                查看所有库
                    show databases;
                创建库
                    create database 库名 default charset utf8mb4;
                查看详情
                    show create database 库名;
                删除库
                    drop database 库名;
                使用库
                    use 库名;

            表:
                查看所有表
                    show tables;
                创建表
                    create table 表名(
                        列名 列的类型 [列的修饰]
                    )engine=innodb default charset utf8mb4;

                    库名，表名，列名: 小写英文字母, _
                    分表分库: log_数字/日期

                    id 整数 int => int32, bigint => int64
                    name 字符串类型
                        string []byte

                        char(length)
                        * varchar(length)
                        * text 65535字节
                        *    longtext 4G
                        blob

                    password
                    sex bool
                        true 男
                        false 女

                        boolean

                    birthday
                        time.time
                        datetime
                    addr
                    tel

                    create table user (
                        id bigint,
                        name varchar(32),
                        password varchar(1024),
                        sex boolean,
                        birthday datetime,
                        addr text,
                        tel varchar(32),
                        index idx_name(name)
                    ) engine=innodb default charset utf8mb4 comment="";
                查看详情
                    desc 表名;
                    show create table 表名;
                删除表
                    drop table 表名;

                数据类型：
                    数值类型
                        布尔类型    boolean
                        整形
                            tinyint 1
                            smallint 2
                            * int 4
                            mediumint 6
                            * bigint 8
                        浮点型
                            * float 4
                            double 8
                            * decimal(m, d) 16 m: 有效位数 小数点之前和小数点之后的数字的个数最大值
                                             d: 小数点后最大位数
                    字符类型
                        char
                        * varchar
                        * text 65535 tinytext 255 mediumtext 16M longtext 4G
                        blob
                        json => 关系型数据库支持文档格式
                        enum
                    日期类型
                        date 年月日
                        time 时分秒
                        * datetime 年月日时分秒
                        timestamp 年月日时分秒
                            更新时会自动更新为当前时间
                修饰
                    主键：PRIMARY KEY
                    主键自动增长: int/bigint AUTO_INCREMENT
                    唯一: UNIQUE KEY
                    非NULL: NOT NULL
                        默认值: DEFAULT value
                                DEFAULT 0
                                DEFAULT ''
                    注释: COMMENT '',

                修改表
                    列 重命名, 添加, 类型修改，删除列
                    alter table 表名 add/change/drop column 列名 列类型 列的修饰;
                                                        修列名 新列名 列类型 列的修饰;
                                                    列名;

                    alter table user add column created_at datetime not null;
            索引
                针对查询字段, 不会针对枚举类型创建
                    用户 status int 1: 未激活 2: 正常使用 3: 锁定 4: 删除

                create unique index 唯一索引名称 [using 索引类型](创建的列) on 表名;
                create index 索引名称 [using 索引类型] (创建的列) on 表名;
                列>1 联合索引
                hash * btree rtree

                create index idx_name on user(name);
                create index idx_birthday on user(birthday) ;
                create index idx_addr on user(addr);
                create index idx_tel on user(tel) ;

                drop index 索引名称 on 表名;


                create table user (
                    id bigint primary key auto_increment,
                    name varchar(32) not null default '',
                    password varchar(1024) not null default '',
                    sex boolean not null default true,
                    birthday date not null,
                    addr text not null default '',
                    tel varchar(32) not null default '',
                    index idx_name(name),
                    index idx_birthday(birthday)
                ) engine=innodb default charset utf8mb4 comment="";

            清空表/重建表: truncate table 表名; // 删除数据 DDL
        DQL
            select * from 表名;
            select * from 表名 WHERE 条件;
                条件: 列名 基准对象 比较
                        比较:
                            关系运算
                                > < = != >= <=

                            like
                                以某个字符串开头 like 'substr%'
                                以某个字符串结尾 like '%substr'
                                包含字符串 like '%substr%'

                                % 0个或任意个字符
                                _ 一个或者固定数量的任意字符 substr


                                % _

                            like binary
                                name kk, AK

                            [not] in (v1, v2, vn) 列表
                            [not] between start and end

                            is null 允许为NULL 值为null的所有记录
                            is not null 允许为NULL 值不为null的所有记录

                            布尔运算
                                and or not
                            小括号

                            addr 西安 或者北京，或者上海的
                            addr = 西安  or addr = 北京  or addr = 上海
                            []string{"西安", "上海", "北京"}
                            in_slice(addr, slice)

                            birthday == '2002' 出生的人
                            birthday >= '2002-01-01 00:00:00' and birthday < '2003-01-01 00:00:00'


                            字符串类型
                                关系运算
                                包含内容 like

                                函数
                                    trim
                                    length
                                    upper
                                    lower

                            数值类型
                                关系运算

                                函数
                                四则运算 + - * / %
                            时间类型
                                关系运算

                                函数
                                date_format(time, '%Y-%m-%d %H:%i:%s')
                                now()
                            bool类型
                限制查询
                    限制查询结果条数 limit N offset M;
                                    limit M, N;
                    分页:
                        pageSize 每页显示数量
                        pageNum 显示第几页 1, 2, 3,...
                        limit pageSize offset pageSize * (pageNum - 1)

                    select * from 表名 where 条件 limit N offset M;
                排序:
                    order by colname [asc/desc],colname [asc/desc];
                    where 条件 order by limit N offset M;

                聚合查询
                    select count(*) from 表名 where 条件;

                    group by
                    select 结果中只能由聚合类的列，以及聚合类函数计算结果

                    count(*)
                    sum(column)
                    avg(column)
                    max(column)
                    min(collumn)

                    accesslog
                    时间 IP URL 状态码
                    datetime -- ip 'GET URL HTTP1/1' status size useragent referer

                    select ip, count(*) from accesslog group by ip;
                    url
                    status
                    select status, count(*) from accesslog group by status;
                    ip status
                    select ip, status, count(*) from accesslog group by ip, status;

                    select 结果 from 表名 where 条件 group by 分组列;


            多表:
                join 多个表之间存在关联关系
                    password  shadow

                    password
                    id, name

                    shadow
                    id, password

                    id 相同的数据是同一个用户的数据

                    create table password(
                        id bigint,
                        name varchar(32)
                    );
                    insert into password values(1, "root"),(2, "kk");
                    create table shadow(
                        id bigint,
                        password varchar(32)
                    );
                    insert into shadow values(1, "123"), (3, "456");

                    id name password

                    select password.id, password.name, shadow.password from password join shadow on password.id=shadow.id
                    where password.id
                    shadow.id
                    order by
                    limit;

                分表:
                    每个月一张表
                    集合 交集，并集，差集
                    union

                    accesslog_01
                    accesslog_02
                    accesslog_03

                子查询
                    password, shadow
                    所有在password存在的用户的密码信息
                    select * from shadow where id in (select id from password)
            别名:
                表别名
                     select p.id, p.name, s.password
                     from password as p join shadow as s on p.id = s.id
                列别名

        DML
            添加
                insert into 表名(列名) values(值);
                insert into 表名 values(值); // 所有列都添加数据

                NOT NULL 但是为设置DEFAULT 插入数据时必须指定该列

                insert into users(name, passwd, brithday, tel) values
                ('abckk', '1231231', '2001-11-12', '1590000100'),
                ('kkabc', '1231231', '2011-11-12', '1590000200'),
                ('abckkabc', '1231231', '2011-11-12', '1590000300'),
                ('abcdef', '1231231', '2011-11-12', '1591000000'),
                ('xxx', '1231231', '2003-11-12', '1590220000'),
                ('ffff', '1231231', '2005-11-12', '1590010000'),
                ('xxxxx', '1231231', '2006-11-12', '1590500000'),
                ('cccc', '1231231', '2007-11-12', '1590000006'),
                ('aaaa', '1231231', '2003-11-12', '1590000011'),
                ('abbbbbckk', '1231231', '2002-11-12', '1590000010'),
                ('ccccc', '1231231', '2001-11-12', '1590000040');


                insert into user(name, password, birthday, addr) values
                ('abckk', '1231231', '2001-11-12', '北京'),
                ('kkabc', '1231231', '2011-11-12', '西安');
            修改
                update 表名
                SET 列名=值, 列名=值
                [WHERE 条件];
            删除
                delete from 表名 [WHERE 条件]

    GO操作MySQL
        database/sql => 定义对数据库操作接口， 未实现针对数据库操作功能
                     => 需要使用第三方包(驱动)
        github.com/go-sql-driver/mysql

        操作：
            查询
                DQL
                Query
            修改
                DDL， DCL, DML
                Exec

        1. 使用驱动
            a. 选择驱动
            b. 初始化导入驱动
        2. 打开数据库(连接池)
        3. 操作:
            a. 修改
                Exec
            b. 查询
                Query
                for + Rows.Next Rows.Scan
        4. 关闭资源
            a. 查询
                Rows.Close()
        5. 进程退出
            关闭数据库（连接池）

        ORM:
            gorm 2.0
            beego orm
            sqlx
    练习