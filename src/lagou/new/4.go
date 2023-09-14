package main

import "fmt"

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

	slice1 := make([]string, 4, 8)
	

	// array1:=[5]string{1:"b",3:"d"}
	// fmt.Println(array1)
}

func main() {
	// test()
	test1()
}
