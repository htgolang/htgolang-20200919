1. 复习
2. 作业
3. 新内容
    a. redis
        字符串string
            "kk"
        列表list
            类似go中切片

            "kk","向宇",""
        散列hash
            map
            {"id":"", "name":""}
        集合set
            map key所有组成的元素
            ("1", "2", "3")
        有序集合zset
            "1" 1.0
            "2" 2.0
            "5" 3.0
            "3" 4.0

        key: value
        key命名: 字符串 英文单词:英文单词 [a-z][a-z0-9]*

        操作:
            key:
                keys
                exists
                del
                ttl
                expire
                type
            对于不同的数据类型有不同的操作
            字符串:
                set:
                    set EX: set+expire
                    set EX NX: 分布式锁
                    key: locker:master => tagname

                    写入数据 SET EX 30 NX
                    processA: a
                        < 10s
                        for {
                            1. SET EX NX:
                                设置成功:
                                    获取锁
                                设置失败:
                                    获取locker:master
                                        GET
                                    检查是否为当前标识
                                    是: 获取锁
                                    不是: 没获取锁
                            2. 获取锁=>续时间, 执行逻辑
                               未获取锁 => 休眠

                        }
                    processB: b
                        for {

                        }
                    processC: c
                        for {

                        }

                mset:
                get:
                mget:
                decr:
                decrby:
                incr:
                incrby:
            list：
                r(l)push
                l(r)pop
                llen
                lrange
                br(l)pop
            hash:
                hkey: hvalue
                hset:
                hget
                hmset
                hmget
                hgetall
                hexsist
                hdel
            set:
                sadd
                smembers
                sinmember
                srem
                scard

                suion
                sdiff
                sinter
            zset
                zadd
                zrange
                zrevrange
                zrem
        ...
    b. go redis
        1. 连接 //
        2. 执行 Do
        3. 结果获取(类型转换)
            redis.Bool
    c. cobra
        web => 认证 session
        api => 不需要启动 session

        1. web和api单独启动
        web: 8080启动
            web.conf
            sessionon=true
        api: 8081启动
            api.conf
            sesionon=false

        需要分别编译成两个程序

        如何修改配置文件
        beego.LoadAppConfig(type, path)


        2. flag 添加使用的配置文件路径
            不知道需要分别针对web/api配置文件启动进程

        3. os.Args
            cmdb web => web
            cmdb api => api

            docker image ls/remove
                   container
                   volume

            子命令的方式
        4. cobra
            Command
            在某个命令中和子命令公共的参数
            PersistentFlags()
            某个子命令特有的参数
            Flags()

    d. gitlab
        gitlab github 私有代码仓库, ci/cd

        gitlab-ce: 社区版
        gitlab-ee: 商业版

        开发 => 测试 => 运维
        ci/cd:
        ci: 持续集成
            开发 -> git ->
            master => 最终交付代码
            feature分支 -> 自测 -> 合并到master => 进行测试
        cd: 持续交付/持续部署
            持续交付：
            开发, 测试
            tag/release => 运维上线(手动)

            持续部署：
            开发，测试，运维

        用户：oa系统自动添加
                gitlab api调用


        gitlab与自动化 工具之间的通信

        token

    e. webhook
    f. ci/cd
        代码触发后续流程
        执行命令:
            步骤1
            步骤2
            步骤3
            步骤4
            ...
            步骤n

        开发提交(起点):
            打 tag
        测试(单元测试):
            go test ./...
        构建:
            go build .
            docker image build .
            // docker image push => tag.gz


            git clone repo
                认证问题
            git checkout tagname
            go build.

        上线(直接部署目标机器):
            A, B, C
            // docker image pull => 远程拷贝 scp
            docker container stop/rm
            docker container run


        流程
        a. 事件
            验证token
            验证tag_push
            验证为创建tag
            获取仓库信息，项目信息, tag信息
        b. 通过项目查找相关信息
            是否自动构建
            构建流程配置
            是否自动部署
            部署流程配置
            部署在哪些机器上(路径，用户名)

        c. (已开启自动构建)构建
            git clone (密钥)
            git checkout
            打包 => tag.gz


        d. 发布包(放在某个位置)
        e. 部署:
            密钥
            scp
            ssh

        问题：
            ssh指纹忽略
            /etc/ssh/ssh_config
            Host *
                StrictHostKeyChecking no

            权限:
                可以读项目的git权限

    g. gitlab api
        腾讯云api/阿里云/aws/...
        http/sdk
            认证:
            协议: request/response
