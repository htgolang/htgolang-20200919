1. 复制文件(只复制文件)
    cp命令 -s srcFile -d dstFile
    srcFile 文件再复制
    dstFile 存在 报错，提示用户是否覆盖，y => 覆盖

    srcf = open(src)
    dstf = create(dst)

    read srcf => ctx
    ctx => writer dstf

    srcf.close
    dstf.close

2. 给路径(文件，目录)
    打印出文件名/目录下的所有文件(包含子目录) [目录不打印]

3. 基于作业2, 只打印 .go, .cgo 后缀的文件

4. 基于作业3, 统计代码行数

5. 用户管理 => 持久化到文件
    gob
    csv
    程序退出的时候保存 => 问题 程序中断无法保存
    每次操作的时候 发生修改 => 保存 => 建议

    加载文件 => 程序启动时
    每次需要的时候 => 读取 => 建议