package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//mapDemo()
	//mapDemo1()
	mapDemo2()
	//mapDemo3()
	//stringDemo()
}

func stringDemo() {
	s := "123456小王呀"
	bytes := []byte(s)
	// 6 + 3 * 3 = 15
	fmt.Println("bytes ==>", bytes, "\nlen ==>", len(s))
	fmt.Println(s[0], s[1], s[2], s[12])

	for i, r := range s {
		fmt.Println(i, r)
	}

	// 字符长度
	fmt.Println("utf8 len", utf8.RuneCountInString(s))
}

func mapDemo3() {
	name1 := map[string]int{"hi": 12}
	name1["hi1"] = 13
	name1["hi2"] = 14
	name1["hi3"] = 15
	name1["hi4"] = 16
	fmt.Println("name1 ==>", name1)
	for k, v := range name1 {
		fmt.Println("key ==>", k, "  v ==>", v)
	}
	fmt.Println("len ==>", len(name1))

	for k1 := range name1 {
		fmt.Println(k1)
	}
}

func mapDemo2() {
	name1 := map[string]int{"hi": 12}
	fmt.Println(name1)
	name, good := name1["hi"]
	fmt.Println("name ==>", name, "\ngood ==> ", good)
	if good {
		fmt.Println("1 good")
	}

	if _, good1 := name1["hi"]; good1 {
		fmt.Println("2 good1 ==>", good1)
	}

	name2, good2 := name1["hi1"]
	fmt.Println("name2 ==>", name2, "\ngood2==> ", good2)
}

func mapDemo1() {
	name1 := map[string]int{"hi": 12}
	fmt.Println(name1)
	name1["hi"] = 142
	fmt.Println(name1)
	name2 := name1["hi"]
	fmt.Println(name2)
	name3 := name1["hi1"]
	fmt.Println(name3)
}

func mapDemo() {
	name := make(map[string]int)
	name["hi"] = 12
	fmt.Println(name)
	// 相当于下面
	name1 := map[string]int{"hi": 12}
	fmt.Println(name1)
}
