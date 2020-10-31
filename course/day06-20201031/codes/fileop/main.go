package main

func main() {
	// 文件
	// 创建 => os.Create
	// 读取 => os.Open
	// 获取属性 os.Open().Stat / os.Stat
	// 修改属性 => 权限，所属人
	// os.Chmod()
	// os.Chown()
	// 重命名
	// os.Rename("a.txt", "b.txt")
	// 删除文件
	// os.Remove("b.txt")

	// 目录
	// 创建
	// fmt.Println(os.Mkdir("a", os.ModePerm))
	// fmt.Println(os.MkdirAll("a/b/c", os.ModePerm))
	// 读取 => os.Open
	// 获取属性 os.Open().Stat/ os.Stat
	// 修改属性  => 权限，所属人
	// os.Chmod()
	// os.Chown()
	// 重命名
	// fmt.Println(os.Rename("b", "d:\\d"))
	// 删除文件夹
	// fmt.Println(os.Remove("a"))
	// os.RemoveAll("a")
}
