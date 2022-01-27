package main

import (
	"fmt"
	"strings"
)

func index_string() {
	i := strings.Index("飞雪无情", "飞雪")
	fmt.Println(i)
}

func if_demo() {
	if i := 6; i > 10 {
		fmt.Println("i > 10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5 < i < 10")
	} else {
		fmt.Println("i < 5")
	}
}

// 从上自下 一旦成功 直接执行并返回
func switch_demo() {
	switch i := 6; {
	case i > 10:
		fmt.Println("i>10")
	case i > 5 && i <= 10:
		fmt.Println("5 < i <= 10")
	default:
		fmt.Println("i <= 5")
	}

}

func main() {
	//index_string()
	//if_demo()
	switch_demo()
}
