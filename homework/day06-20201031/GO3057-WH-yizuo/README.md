
# Note

* **user_manager** CMS管理系统，使用csv作为后端持久化存储
* **cpFile** 复制工具
* **findFile** 打印该目录及所有子目录下所有文件的文件名
* **findGoFile**  仅打印该目录及所有子目录下以.go或者.cgo为后缀的文件名
* **countCodeLine**   统计.go或者.cgo文件的代码行数
    yizuo@yizuodeMacBook-Pro countCodeLine % go run countCodeLine.go path
    脚本路径:path/a/1/1.go  代码行数:9
    脚本路径:path/a/2/2.go  代码行数:9
    脚本路径:path/a/3/3.cgo 代码行数:9
