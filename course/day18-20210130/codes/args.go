package main

type testArgs struct {
	a, b int
	c, d string
}

func NewTestArgs() *testArgs {
	return &testArgs{
		a: 1,
		b: 2,
		c: "",
		d: "",
	}
}

func test(a, b int, c, d string) {
	// 函数形参太多
	// 无默认值

}

func testV2(args *testArgs) {

}

func main() {

}
