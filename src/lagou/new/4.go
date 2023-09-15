package main

import (
	"fmt"
	"unicode/utf8"
)

func test() {
	sum := 0

	for i := 1; i < 100; i++ {

		if i%2 != 0 {

			continue

		}

		sum += i

	}

	fmt.Println("the sum is", sum)
}

func test1() {
	array := [...]string{"1", "2", "3", "4", "5"}
	// for i, v := range array {
	// 	fmt.Println("对应值：%d, %s", i, v)
	// }
	slice := array[2:5]
	fmt.Println(slice)

	slice1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice1), cap(slice1))
	// slice1 := make([]string, 4, 8)
	slice2 := append(slice1, "ff")
	fmt.Println(slice2)
	slice2 = append(slice2, "abc", "sdafg")
	fmt.Println(slice2)
	slice2 = append(slice2, slice...)
	fmt.Println(slice2)

	// array1:=[5]string{1:"b",3:"d"}
	// fmt.Println(array1)
}

func test2() {
	nameAgeMap := make(map[string]int)
	nameAgeMap["士大夫"] = 20
	fmt.Println(nameAgeMap)
	nameAgeMap1 := map[string]int{"你说的": 12}
	fmt.Println(nameAgeMap1)
	fmt.Println("")
}

func test3() {
	nameAgeMap := make(map[string]int)

	nameAgeMap["飞雪无情"] = 20

	age, ok := nameAgeMap["飞雪无情1"]
	// fmt.Println(nameAgeMap)
	delete(nameAgeMap, "飞雪无情")
	// fmt.Println(nameAgeMap)
	if ok {
		fmt.Println(age)
	}
	nameAgeMap["飞雪无情"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22
	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}
	fmt.Println(len(nameAgeMap))
}

func test4() {
	s := "合励咯十多个格瑞德sd"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[15])
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i, r := range s {
		fmt.Println(i, r)
	}
}

func main() {
	// test()
	// test1()
	// test2()
	// test3()
	test4()
}
