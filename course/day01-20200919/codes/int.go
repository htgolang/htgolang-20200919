package main

import "fmt"

func main() {
	var age int = 31
	var char byte = 'a'      // 97 => ascii编码 a => 整数值
	var codepoint rune = '我' // unicode编码 我 => 整数值

	fmt.Println(age, char, codepoint)
	fmt.Printf("%T %T %T\n", age, char, codepoint)
	fmt.Printf("%d %c %q\n", age, char, codepoint)
	fmt.Printf("%b\n", age)

	// 十进制, 八进制, 十六进制
	var (
		t10 = 10
		t8  = 012 // 010 => 8 002 => 2
		t16 = 0XA // 0X 0x => A 10
	)
	fmt.Println(t10, t8, t16)
	fmt.Printf("%d %x %o\n", t10, t10, t10)

	// 算数运算
	// + - * / %
	fmt.Println(1 + 2) // 3
	fmt.Println(1 - 2) // -1
	fmt.Println(1 * 2) // 2
	fmt.Println(1 / 2) // 0 int/int => 向下取整
	fmt.Println(1 % 2) // 1

	// 自增 ++ ，自减 --
	age++            // age = age + 1
	fmt.Println(age) // 32
	age--            // age = age - 1
	fmt.Println(age) // 31

	// 关系运算 => bool
	// > >= < <= != ==
	fmt.Println(1 > 2)  // false
	fmt.Println(1 >= 2) // false
	fmt.Println(1 < 2)  // true
	fmt.Println(1 <= 2) // true
	fmt.Println(1 != 2) // true
	fmt.Println(1 == 2) // false
	// 位运算
	// & | ^ << >> &^
	// 两个整数转为二进制进行计算 对应的位进行计算
	// & 两个为1 => 1
	// | 只要有一个为1 => 1
	// ^ 抑或 相同->0, 不同->1
	// << 10 => 2 ^ n
	// >> 0 => /2^n
	// &^ a &^ b 擦除 b(1->a->0)

	// 赋值
	// = += -= *= /= %= &= |= ^= <<= >>=
	// a += b => a = a + b
	// a -= b => a = a - b

	a := 10 // int
	a += 5  // a = a + 5
	a *= 3  // a = a * 3
	fmt.Println(a)

	// 不同数据类型之间是不能进行计算
	var b int8 = 5
	fmt.Println(a + int(b))
	fmt.Println(int8(a) + b)
	// int 64 int8 => 8 -128 - 127
	// int 128 => int8
	// 0000 1000 0000 => 1000 0000 => -128
	a = 128
	fmt.Println(int8(a))

	fmt.Printf("%10d %-10d %010d %d\n", a, a, a, a)
}
