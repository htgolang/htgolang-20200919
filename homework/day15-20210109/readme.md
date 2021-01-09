1. Beego知识
2. 用户管理 删除 -> 改为逻辑删除
    deleted_at => now()
    is null 未删除

    新建 => name重复
    编辑 => 已经删除，name重复
    查询 => 已经删除不能查询出来
3. 挑战
    访问控制

    用户管理
    角色管理（展示列表）
        有哪些角色，角色对应哪些URL有权限
        展示对应菜单


    访问日志：
        上传访问日志 => 访问日志TOP

    无页面相关操作：URL,菜单


    角色 => URL (ContollerName.ActionName)

    管理员
        所有URL进行操作

    操作员：
        只能进行访问日志相关URL的操作权限

    菜单，URL都要限制
    管理员：看到所有菜单，所有功能成功操作
    操作员：只能看到访问日志菜单，通过手动输入URL不能操作成功

    Prepare()
    User -> Role -> Menus => Ctx.Data[Menus]
                    Roles -> Urls => Controller.Action

    URL -> 当前Controller Action => in Urls -> 在 允许访问
                                             不在Abort到无权限页面