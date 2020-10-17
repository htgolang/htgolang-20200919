package main

/*
	1. go标准包
	2. 第三方包
	3. 本地导入
*/
// 绝对导入
// import (
// 	"fmt"

// 	"github.com/imsilence/testmath"
// )

// 相对导入
/*
import (
	"./xxx/xxxx/gpkg"
)
gpkg.Var
*/
// 点导入
/*
import (
	. "fmt"


	. "github.com/imsilence/testmath"
)

函数()
*/
// 别名导入

// "github.com/imsilence/testmath"

// 下划线导入(初始化包)
import (
	"fmt"
	gtestmath "github.com/imsilence/testmath"
	_ "testimport/testmath"
)

func main() {
	fmt.Println(gtestmath.Add(1, 2))
	// fmt.Println(testmath.Mul(2, 3))
}
