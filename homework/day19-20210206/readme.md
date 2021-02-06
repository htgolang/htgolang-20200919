1. build/deploy流程完成
2. Project/Host页面配置
    KeyFile 增删改查
    <form enctype="">
        multipart/form-data method=POST
        input type=file

    Host增删改查(测试)
        Addr
        Port
        User
        WorkDir
        KeyFile

    Project 增删改查
        选择多个Host
        checkbox
        select multiple

3. 挑战
    build =>
        project packageFile 编译时的版本信息 output status 操作按钮(删除/重新编译)
    deploy
        project packageFile host=>status/output 操作(删除/重新部署)

    project => 构建/部署