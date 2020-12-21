 CREATE TABLE user (
                    id BIGINT PRIMARY KEY AUTO_INCREMENT,
                    name VARCHAR(32) NOT NULL DEFAULT '',
                    password VARCHAR(1024) NOT NULL DEFAULT '',
                    sex BOOLEAN NOT NULL DEFAULT TRUE,
                    born DATE NOT NULL,
                    address TEXT NOT NULL DEFAULT '',
                    cell VARCHAR(32) NOT NULL DEFAULT '',
                    INDEX IDX_NAME(name),
                    INDEX IDX_BIRTHDAY(born)
                ) ENGINE=INNODB DEFAULT CHARSET UTF8MB4 COMMENT="";
                -- user: web 
                -- pass: web
