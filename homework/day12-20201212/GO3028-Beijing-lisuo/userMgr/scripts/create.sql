 CREATE TABLE user (
                    id BIGINT PRIMARY KEY AUTO_INCREMENT,
                    name VARCHAR(32) NOT NULL DEFAULT '',
                    password VARCHAR(1024) NOT NULL DEFAULT '',
                    sex BOOLEAN NOT NULL DEFAULT TRUE,
                    born DATE NOT NULL,
                    address TEXT NOT NULL DEFAULT '',
                    cell VARCHAR(32) NOT NULL DEFAULT '',
                    created_at DATETIME NOT NULL,
                    updated_at DATETIME NOT NULL,
                    deleted_at DATETIME,
                    INDEX IDX_NAME(name),
                    INDEX IDX_BIRTHDAY(born)
                ) ENGINE=INNODB DEFAULT CHARSET UTF8MB4 COMMENT="";
                -- user: web 
                -- pass: web


INSERT INTO user 
  (name, password, sex, born, address, cell, created_at, updated_at) 
    VALUES 
  ('admin', password('admin123'), 1, '1995.03.04', 'Venus', '18811739999', NOW(), NOW()),
  ('jaccy', password('jaccy'), 1, '1895.06.04', 'London', '18811738998', NOW(), NOW()),
  ('leslie', password('imwhatim'), 1, '1975.03.04', 'HongKong', '18811738888', NOW(), NOW());
