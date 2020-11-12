# **CMS Version 4**

基于原项目[CMS Version 3](https://github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/GO3020-beijing-jiangchen), 新增后端持久化方式json与csv, 其中csv为滚动存储模式，仅保留最新的5次新增的数据。新增命令行参数`-db`与`-init`，前者传参使用的数据库后端(boltdb, csv, json, memory, 其中memory代表纯使用内存而不持久化)，后者传参是否进行数据库初始化

* ## **定义存储接口，如下所示**

```go
type Storage interface {
	Check() error
	Init() error
	SyncFromDBToMemory() error
	SyncFromMemoryToDB() error
	InsertToDB(*User) error
	DeleteFromDB(int) error
	ModifyFromDB(*User) error
	GetNonAdmin() error
}
```

### 目前boltdb, csv, json均实现了上述接口

* ## **定义内存运行时结构体**

```go
type MemoryUsersEntity struct {
	Users             *[]User
	PersistentStorage Storage
}
```

### Users为指向内存中用户切片的指针, PersistentStorage为Storage接口, 使用命令行参数指定特定的存储(boltdb, csv, json, memory)，并对接口进行赋值, 从而模仿多态

* ## **如何编译**

```go
~# go build
```

* ## **如何使用**

```go
~# ./CMS -h
  -db string
        persistent db type (boltdb, csv, json, memory), default is memory. (default "memory")
  -init
        whether init database or not, default is false.
  NOTE: if db is "all" combined with "-init", delete all database.
```