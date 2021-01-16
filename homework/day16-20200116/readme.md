1. prometheus整理
2. 开发
    http开发web user

    handler/handlerfunc

    暴露metrics

    总的请求数量
    每个请求出现的次数 counter
    每个请求请求时间 historygram/summary
    每个请求状态码的出现次数 counter
    存活例程数量

    使用prometheus采集

3. 实现web basic认证
    main.go handler/handlerfunc 认证通过返回当前使劲 unixtime

    用户名:密码 => passwd =>
    每行一个用户信息: username: password(bcrypt)