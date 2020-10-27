# **CMS Version 2**

> 基于原项目[CMS Enhancement](https://github.com/htgolang/htgolang-20200919/tree/master/homework/day04-20201017/GO3020-beijing-jiangchen), 更改用户存储数据结构，由map[string]string 更改为 自定义结构体struct,给用户增添出生日期属性、密码属性，增加登录界面用户名+密码验证功能，增加新增用户以及修改用户的用户名校验功能

* ## **定义数据结构如下所示**

```go
//User ...
//Global data structure of User element.
type User struct {
	ID       int
	Name     string
	Tel      string
	Address  string
	Birthday time.Time
	Password [16]byte
}
```

****

* ## **增加登录界面用户名 + 密码验证功能，输入3次错误密码强制退出**


```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login username.
username : admin
please input password: 
password incorrect...
please input password: 
password incorrect...
please input password: 
password incorrect...
Bye :(
```

```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login username.
username : root
please input password: 

login success.

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > show
+------+-------+---------------+--------------------------------+------------+
|  ID  | NAME  |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+-------+---------------+--------------------------------+------------+
| 1001 | admin | +1 4406665321 | 2426 Wildwood Street, Medina,  | 1990-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1002 | root  | +1 4406665322 | 2427 Wildwood Street, Medina,  | 1992-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1003 | super | +1 4406665323 | 2428 Wildwood Street, Medina,  | 1994-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
```

****

* ## **添加新增用户校验功能，不能添加相同用户名的用户**

```bash
CMS > 
CMS > add
Please input Name > admin
Please input Tel Number > 
input Tel is blank, so take Tel "+1 2024561111" as default.
Please input Address > 
input Address is blank, so take Address "1600 Pennsylvania Avenue NW, Washington, DC 20500, United States" as default.
Please input Birthday(format: YYYY-MM-DD) > 1990-10-19
Please input password > 
name duplicate in CMS...Abort...
CMS > 
CMS > show
+------+-------+---------------+--------------------------------+------------+
|  ID  | NAME  |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+-------+---------------+--------------------------------+------------+
| 1001 | admin | +1 4406665321 | 2426 Wildwood Street, Medina,  | 1990-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1002 | root  | +1 4406665322 | 2427 Wildwood Street, Medina,  | 1992-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1003 | super | +1 4406665323 | 2428 Wildwood Street, Medina,  | 1994-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
CMS > 
CMS > add
Please input Name > new
Please input Tel Number > 
input Tel is blank, so take Tel "+1 2024561111" as default.
Please input Address > 
input Address is blank, so take Address "1600 Pennsylvania Avenue NW, Washington, DC 20500, United States" as default.
Please input Birthday(format: YYYY-MM-DD) > 1976-04-29
Please input password > 
Add User Finish.
CMS > 
CMS > show
+------+-------+---------------+--------------------------------+------------+
|  ID  | NAME  |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+-------+---------------+--------------------------------+------------+
| 1001 | admin | +1 4406665321 | 2426 Wildwood Street, Medina,  | 1990-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1002 | root  | +1 4406665322 | 2427 Wildwood Street, Medina,  | 1992-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1003 | super | +1 4406665323 | 2428 Wildwood Street, Medina,  | 1994-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1004 |  new  | +1 2024561111 |  1600 Pennsylvania Avenue NW,  | 1976-04-29 |
|      |       |               |  Washington, DC 20500, United  |            |
|      |       |               |             States             |            |
+------+-------+---------------+--------------------------------+------------+
```

****

* ## **添加修改用户校验功能，同一用户可保持用户名不修改，但不可将用户名修改为系统已存在的用户的用户名**

```bash
go run ./
---------------------
 __    _ __     ___
/ _|  | '  \   (_-<
\__|  |_|_|_|  /__/
---------------------

Welcome to CMS System.
Please input login username.
username : root
please input password: 

login success.

Commands:
show:     show current users.
add:      add new user.
delete:   delete user.
modify:   modify user.
query:    query user.
help:     print help messages.
exit:     exit.

CMS > show
+------+-------+---------------+--------------------------------+------------+
|  ID  | NAME  |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+-------+---------------+--------------------------------+------------+
| 1001 | admin | +1 4406665321 | 2426 Wildwood Street, Medina,  | 1990-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1002 | root  | +1 4406665322 | 2427 Wildwood Street, Medina,  | 1992-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1003 | super | +1 4406665323 | 2428 Wildwood Street, Medina,  | 1994-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
CMS > modify
please input the ID of User you want to modify > 1003
********
Find User:
********
+------+-------+---------------+--------------------------------+------------+
|  ID  | NAME  |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+-------+---------------+--------------------------------+------------+
| 1003 | super | +1 4406665323 | 2428 Wildwood Street, Medina,  | 1994-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
Do you want to modify this User? (y/n) > y
Please input new Name > root
Please input new Tel Number > 
input Tel is blank, so keep origin Tel as default.
Please input new Address > 
input Address is blank, so keep origin Address as default.
Please input new Birthday(format: YYYY-MM-DD) > 1990-11-11
Please input new password > 
name duplicate in CMS...Abort...
CMS > show
+------+-------+---------------+--------------------------------+------------+
|  ID  | NAME  |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+-------+---------------+--------------------------------+------------+
| 1001 | admin | +1 4406665321 | 2426 Wildwood Street, Medina,  | 1990-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1002 | root  | +1 4406665322 | 2427 Wildwood Street, Medina,  | 1992-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1003 | super | +1 4406665323 | 2428 Wildwood Street, Medina,  | 1994-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
CMS > modify
please input the ID of User you want to modify > 1002
********
Find User:
********
+------+------+---------------+--------------------------------+------------+
|  ID  | NAME |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+------+---------------+--------------------------------+------------+
| 1002 | root | +1 4406665322 | 2427 Wildwood Street, Medina,  | 1992-01-01 |
|      |      |               |              Ohio              |            |
+------+------+---------------+--------------------------------+------------+
Do you want to modify this User? (y/n) > y
Please input new Name > root2
Please input new Tel Number > 
input Tel is blank, so keep origin Tel as default.
Please input new Address > 
input Address is blank, so keep origin Address as default.
Please input new Birthday(format: YYYY-MM-DD) > 1990-11-11
Please input new password > 
Modify User Finish.
CMS > s
input error...please try again.
CMS > 
CMS > show
+------+-------+---------------+--------------------------------+------------+
|  ID  | NAME  |      TEL      |            ADDRESS             |  BIRTHDAY  |
+------+-------+---------------+--------------------------------+------------+
| 1001 | admin | +1 4406665321 | 2426 Wildwood Street, Medina,  | 1990-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1002 | root2 | +1 4406665322 | 2427 Wildwood Street, Medina,  | 1990-11-11 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
| 1003 | super | +1 4406665323 | 2428 Wildwood Street, Medina,  | 1994-01-01 |
|      |       |               |              Ohio              |            |
+------+-------+---------------+--------------------------------+------------+
```