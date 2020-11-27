1. 复制文件
   执行命令：`go run cp1.go -s test.go -d test1.go `
   - [cp1.go](./copyFile/cp1.go) 调用 io.Copy 直接完成。。。
   - [cp2.go](./copyFile/cp2.go) 
2. 打印文件
   执行命令：`go run ../copyFile/listFile1.go -s "/Users/aomine/Documents/projects/scripts/go/htgolang-20200919/homework/day06-20201031/GO3004-zhaoweiping/listFile/listFile1.go"`
   - [listFile1.go](./listFile/listFile1.go) 调用 ioutil.ReadDir 直接完成
   - [listFile2.go](./listFile/listFile2.go)  
   - [listFile3.go](./listFile/listFile3.go)  
      
3.  只打印 .go, .cgo 后缀的文件
   - [listFile4.go](./listFile/listFile4.go)  
4.  统计代码行数
   - [sumFile.go](./listFile/sumFile.go)
5.  用户管理
   - [CMS](./CMS/)