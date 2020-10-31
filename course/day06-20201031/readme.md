1. 复习
    type StructName struct{
        AttrName1 AttrType1
        AttrName2 AttrType2
        ...
        AttrNameN AttrTypeN
    }

    var varname struct{
        AttrName1 AttrType1
        AttrName2 AttrType2
        ...
        AttrNameN AttrTypeN
    }

    type StructName struct{
        AttrName01 AttrType01
        AttrName02 AttrType02
        ...

        AttrName0N AttrType0N
        StructType1
        StructType2
        ...
        StructTypeN
    }


    func (s *StructName) func() {
        *s = &StructName{}
    }

    地址 => 送货
    李彬 => 地址 北京市海淀区32路a信箱
    我 => 李彬
    * 我 -> 地址 => a信箱 => 苹果 => 李彬 => 能拿到
    我 => 卫智鹏

2. 作业(课堂不处理)
3. 新内容
    模块: version
        自定义包 => version

        包的提供者 => 打tag
        包的使用者 => 改版本
    文件
        你读一个文件:
            1. 找到文件 => 文件存储的位置 => 路径
            2. 双击
            3. 从前到后移动光标位置阅读文件内容
            4. 关闭文件
        你改一个文件: a->b
            1. 找到文件 => 路径
            2. 双击 打开
            3. 从前到后移动光标位置阅读文件内容，并进行修改
            4. 保存
            5. 关闭文件

        1. 路径
            相对路径
                cd /home/kk/Desktop

                vim a.txt vim当前执行的路径
                相对路径: 程序执行的路径

                vim
                    vim二进制文件的路径: /usr/bin/vim
                    执行的路径: /home/kk/Desktop

                    a.txt /home/kk/Desktop/a.txt
            绝对路径
                vim /home/kk/Desktop/a.txt
                vim c:/Users/kk/Desktop/a.txt
        2. 文件类型
            cat/记事本 => 文本内容（无乱码） => 文本文件 => string
                            有乱码(word, zip) => 二进制文件 => []byte


        基本操作(不带缓冲IO): 读、写
        标准输入、输出、错误
        带缓冲的IO
    目录
    文件/目录的属性信息 标准包
    编码格式
        csv
        gob
4. 问题处理
    a. vscode 提示工具更新、工具缺失 => 按提示安装
    b. vscode pakcages.Load error => 打开目录并非go模块, 使用vscode打开go项目目录
    c. vscode 写代码
        运行: 命令行里运行
            windows: cmd(上课) / powershell
            Linux: terminal
    d. 程序报错
        错误截图 => 代码截图（截出来行号）