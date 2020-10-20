1. 包
    包名: 标识符规范
          小写英文字母, 尽量短小
          同一个文件夹内所有go文件的包名必须一致
          包名与文件夹名尽量保持一致
2. 包管理
    gopath+vendor
        GOPATH环境变量 配置目录 项目开发目录
        bin => 二进制文件
        pkg => 库文件
        src => 源文件 main => win exe

        包名 main: 编译成二进制程序
                    main函数是程序入口
            非main:
                math => 库 用来提供功能的
                调用 main/非main

                .a
            链接 二进制程序

        使用第三方包
            下载到本地 GOPATH/src
                testpkg/math v1, v2 v3
            项目A v1 => GOPATH
            项目B v2 => GOPATH
            go 1.5 vendor

            项目目录下创建vendor目录 第三方包下载考本到vendor

            a包中使用 a/vendor => a/vendor
                     a/../vendor

                    GOPATH/src/vendor
                    GOPATH/src
                    GOROOT/src

        升级
    gomod
        go 1.11 => gomod

        是否开启gomod => GO111MODULE 1.11


        GO111MOUDLE = on => 必须使用go mod
                    => off => 必须使用gopath方式
                    => 空(默认) => 自己选择
                            不在 GOPATH 且 当前目录有 go.mod => gomod
                            否则 GOPATH
        1. 版本 => go.mod
        2. GOPROXY => go mod
        3. 使用GOPATH要不然放在固定的目录（GOPATH），否则需要配置新的GOPATH
        4. go mod代码可以随意放置在任何位置
        5. go replace
            google.cn/aaa -> github.com/aaa

        初始化模块
        go mod init 项目名称
            代码仓库的路径/项目名称
            一般命名方式 => github.com/imsilence/testgomod
                           testgomod

            go git/svn 第三方包 => 自动下载

        a. 使用第三方包
        b. 提供第三方包
        c. 自己正常使用

    外网不通
    gitlab/gitee
    1. 自己内网, 没有以来任何第三方库
    2. 依赖内部其他团队提供的库
    3. 依赖外网的第三方库
        a. 编译环境 => 目标站点
                        => acl
                        => proxy <=> 镜像
                        => 第三方包下载 => gitlab/gitee
                            replace
    gopath+vendor


    打包 => 部署


    点导入
    a Add
    b Add
    import (
        . "a"
        . "b"
    )

    Add()

    程序执行 main
            import => 其他程序

3. 标准包
    add 程序 可以接收多个int格式的字符串, 至少输入两个
    少于两个给提示，程序的使用方法为add 1 2 [3 4 5]
    输入不是数字格式的字符串时，提示 程序的使用方法为add 1 2 [3 4 5]
    计算所有数字总和，并输出结果

    从不知道-》知道
        a. 培训->60% ->
            ->40%
        b. 看文档
        c. 搜索引擎
        d. 参加交流/参加会议
        e. 练习

4. 第三方包
5. 单元测试