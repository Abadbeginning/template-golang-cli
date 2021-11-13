package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func printArray1(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 5, 6, 7, 8}
	var grid [4][3]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 下标,数值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(&arr1)
	//printArray(arr2)
	//printArray(arr3)
	fmt.Println(arr1, arr3)
}
