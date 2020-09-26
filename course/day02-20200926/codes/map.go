package main

import "fmt"

func main() {
	// G -> 名字
	// key type
	// value type
	// map[keytype]valuetype
	var names map[string]string = map[string]string{"Go3037": "李雄"} // nil

	// 一定要进行初始化
	// 字面量
	// map[keytype]valuetype{k: v, k1: v1, k2: v2, kn: vn}

	fmt.Printf("%T %v\n", names, names)
	// key 访问元素
	// key 添加或者修改元素
	fmt.Printf("%q\n", names["Go3005"])
	names["Go3005"] = "卫智鹏"
	names["Go3037"] = "lixiong"
	fmt.Println(names)
	names["Go0001"] = ""
	fmt.Println(names["Go0001"], names["Go0002"])
	v, ok := names["Go0001"] // v = k => value k不存在返回value type 0值， ok bool
	fmt.Println(v, ok)
	v, ok = names["Go0002"] // v = k => value k不存在返回value type 0值， ok bool
	fmt.Println(v, ok)

	// 删除元素
	delete(names, "Go3005")
	fmt.Println(names)
	for k := range names {
		fmt.Println(k, names[k])
	}

	for k, v := range names {
		fmt.Println(k, v)
	}

	var scores = make(map[string]int)
	scores["Go3028"] = 80
	scores["Go3027"] = 82
	scores["Go3026"] = 82
	scores["Go3025"] = 82
	scores["Go3029"] = 82
	fmt.Println("--------------")
	for k, v := range scores {
		fmt.Println(k, v)
	}
	scores["Go3024"] = 82
	scores["Go3128"] = 80
	scores["Go3127"] = 82
	scores["Go3126"] = 82
	scores["Go3125"] = 82
	scores["Go3129"] = 82
	scores["Go3124"] = 82

	fmt.Println("--------------")
	for k, v := range scores {
		fmt.Println(k, v)
	}
}
