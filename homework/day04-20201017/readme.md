1. 用户管理
    增/删/改/查 => 单独go

    循环 让用户从控制台输入指令
        add => 执行add功能
        delete => 执行delete功能
        modify => 执行modify功能
        query => 执行query功能
        exit => 退出
        help => 帮助信息

2. 添加密码功能
    启动程序时让用户输入密码 比较对象
    再程序中内置一个md5值

    计算用户输入的密码MD5值 与程序中md5比较
    输入失败3次退出程序，如果成功执行用户操作

3. 用户输出 tablewriter

4. 使用映射存储操作指令及调用函数关系(挑战)
    map[string]callback
    callback = get(input)
    callback()