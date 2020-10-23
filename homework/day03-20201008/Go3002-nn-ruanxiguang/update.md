# 👍
一部分共用的变量，实际上是可以放在函数之外进行全局声明的，这样可以减少代码量，还方便快速修改。
var (
	ID    string
	name  string
	phone string
	addr  string
	OK    string
)
func add(){ ID=1 ....}
func main() {}