1. 课堂练习，课堂笔记整理
2. map[string]string => User
    id int
    name string
    addr string
    tel string
    birthday time.Time
    password => md5存放

3. 登陆 password => 输入用户名和密码
    可以用不同的用户名和密码登陆
    预置用户admin/password自定

4. 在用户添加输入/修改的时候验证
    name不能重复

    编辑 => 1 aaa => aaa => 需要通过
            已经存在 1 aaa => 2 aaa => 不能通过
    添加 => 1 aaa => aaa => 不能通过