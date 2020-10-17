# **CMS enhancement**

## **添加密码功能，默认密码为admin，连续输入3次错误则程序退出**

```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login password.
password : 
password : 
password : 
password incorrect...
password : 
password incorrect...
password : 
password incorrect...
Bye :(
```

****

## **查看帮助,打印现有用户,退出演示**


```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login password.
password : 

login success.

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > 
CMS > 
CMS > 
CMS > help

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > show
+------+--------+---------------+--------------------------------+
|  ID  |  NAME  |    CONTACT    |            ADDRESS             |
+------+--------+---------------+--------------------------------+
| 1001 | Alice  | +1 4406665321 | 2426 Wildwood Street, Medina,  |
|      |        |               |              Ohio              |
+------+--------+---------------+--------------------------------+
| 1002 | Norman | +1 6789143737 |  4548 Davis Street, Norcross,  |
|      |        |               |            Georgia             |
+------+--------+---------------+--------------------------------+
| 1003 | Connie | +1 2184173411 | 1485 Laurel Lee, Forest Lake,  |
|      |        |               |           Minnesota            |
+------+--------+---------------+--------------------------------+
| 1004 | David  | +1 8455462309 |       3627 Camden Place,       |
|      |        |               |     Poughkeepsie, New York     |
+------+--------+---------------+--------------------------------+
CMS > quit
input error...please try again.
CMS > exit
Bye :)
```

****

## **添加用户演示**

```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login password.
password : 

login success.

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > add
Please input Name > zhangsan
Please input Contact Number > 12345678
Please input Address > Greenland
Add Users Finish.
CMS > 
CMS > show
+------+----------+---------------+--------------------------------+
|  ID  |   NAME   |    CONTACT    |            ADDRESS             |
+------+----------+---------------+--------------------------------+
| 1001 |  Alice   | +1 4406665321 | 2426 Wildwood Street, Medina,  |
|      |          |               |              Ohio              |
+------+----------+---------------+--------------------------------+
| 1002 |  Norman  | +1 6789143737 |  4548 Davis Street, Norcross,  |
|      |          |               |            Georgia             |
+------+----------+---------------+--------------------------------+
| 1003 |  Connie  | +1 2184173411 | 1485 Laurel Lee, Forest Lake,  |
|      |          |               |           Minnesota            |
+------+----------+---------------+--------------------------------+
| 1004 |  David   | +1 8455462309 |       3627 Camden Place,       |
|      |          |               |     Poughkeepsie, New York     |
+------+----------+---------------+--------------------------------+
| 1005 | zhangsan |   12345678    |           Greenland            |
+------+----------+---------------+--------------------------------+
CMS > exit
Bye :)
```

****

## **删除用户演示**

```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login password.
password : 

login success.

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > delete
please input the ID of User you want to delete > 1001
********
Find User:
********
+------+-------+---------------+--------------------------------+
|  ID  | NAME  |    CONTACT    |            ADDRESS             |
+------+-------+---------------+--------------------------------+
| 1001 | Alice | +1 4406665321 | 2426 Wildwood Street, Medina,  |
|      |       |               |              Ohio              |
+------+-------+---------------+--------------------------------+
Do you want to delete this User? (y/n) > y
Remove User success.
CMS > show
+------+--------+---------------+--------------------------------+
|  ID  |  NAME  |    CONTACT    |            ADDRESS             |
+------+--------+---------------+--------------------------------+
| 1002 | Norman | +1 6789143737 |  4548 Davis Street, Norcross,  |
|      |        |               |            Georgia             |
+------+--------+---------------+--------------------------------+
| 1003 | Connie | +1 2184173411 | 1485 Laurel Lee, Forest Lake,  |
|      |        |               |           Minnesota            |
+------+--------+---------------+--------------------------------+
| 1004 | David  | +1 8455462309 |       3627 Camden Place,       |
|      |        |               |     Poughkeepsie, New York     |
+------+--------+---------------+--------------------------------+
CMS > exit
Bye :)
```

****

## **修改用户演示**

```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login password.
password : 

login success.

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > modify
please input the ID of User you want to modify > 1003
********
Find User:
********
+------+--------+---------------+--------------------------------+
|  ID  |  NAME  |    CONTACT    |            ADDRESS             |
+------+--------+---------------+--------------------------------+
| 1003 | Connie | +1 2184173411 | 1485 Laurel Lee, Forest Lake,  |
|      |        |               |           Minnesota            |
+------+--------+---------------+--------------------------------+
Do you want to modify this User? (y/n) > y
Please input new Name > Donald Trump
Please input new Contact Number > 
input Contact is blank, so keep origin Contact as default.
Please input new Address > USA
CMS > 
CMS > show
+------+--------------+---------------+--------------------------------+
|  ID  |     NAME     |    CONTACT    |            ADDRESS             |
+------+--------------+---------------+--------------------------------+
| 1001 |    Alice     | +1 4406665321 | 2426 Wildwood Street, Medina,  |
|      |              |               |              Ohio              |
+------+--------------+---------------+--------------------------------+
| 1002 |    Norman    | +1 6789143737 |  4548 Davis Street, Norcross,  |
|      |              |               |            Georgia             |
+------+--------------+---------------+--------------------------------+
| 1003 | Donald Trump | +1 2184173411 |              USA               |
+------+--------------+---------------+--------------------------------+
| 1004 |    David     | +1 8455462309 |       3627 Camden Place,       |
|      |              |               |     Poughkeepsie, New York     |
+------+--------------+---------------+--------------------------------+
CMS > 
CMS > bye
input error...please try again.
CMS > exit
Bye :)
```

****

## **查询演示**

```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login password.
password : 

login success.

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > query
Please input Query String > david
********
Users that contain the Query String:
+------+-------+---------------+--------------------------------+
|  ID  | NAME  |    CONTACT    |            ADDRESS             |
+------+-------+---------------+--------------------------------+
| 1004 | David | +1 8455462309 |       3627 Camden Place,       |
|      |       |               |     Poughkeepsie, New York     |
+------+-------+---------------+--------------------------------+
CMS > exit
Bye :)
```
