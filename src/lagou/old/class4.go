package main

import "fmt"

func arrayDemo() {
	array := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(array[3])
	array1 := [...]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(array1[3])
	for i := 0; i < len(array1); i++ {
		fmt.Print(array1[i], "  ")
	}
	fmt.Println()
	for _, v := range array1 {
		fmt.Printf("array1值:%s\n", v)
	}
	fmt.Println()
	for i, v := range array1 {
		fmt.Printf("array1数组索引:%d,array1值:%s\n", i, v)
	}
}

func sliceDemo() {
	array := [...]string{"a", "b", "c", "d", "e", "f"}
	slice := array[2:3]
	fmt.Println(slice)
	slice = array[2:6]
	fmt.Println(slice)
}

func sliceDemo1() {
	array := [...]string{"a", "b", "c", "d", "e", "f"}
	slice := array[:3]
	fmt.Println("array[ : 3] ==>", slice)
	slice = array[2:]
	fmt.Println("array[2 : ] ==>", slice)
	slice = array[:]
	fmt.Println("array[ : ] ==>", slice)
	// fmt.Println(len(array))
}

func sliceDemo2()  {
	array := [...]int{1, 2, 3, 4, 5, 6}
	slice:=array[3:5]
	fmt.Println(slice)
	slice[1] = 11
	fmt.Println(array)
	array[4] = 5
	fmt.Println(array)
}

func sliceDemo3()  {
	slice1:=make([]int,4)
	fmt.Println(slice1,"  长度 --> ", len(slice1), cap(slice1))
	slice1 = make([]int,4,8)
	fmt.Println(slice1,"  长度 --> ", len(slice1), cap(slice1))
}

func appendDemo()  {
	array := []int{1}
	slice1:=make([]int,4)
	slice2 := append(slice1, 12)
	fmt.Println("slice1 ==>", slice1)
	fmt.Println("slice2 ==>", slice2)
	fmt.Println("array ==>", array)
	slice2 = append(slice1, 14, 16, 15)
	fmt.Println("slice1 ==>", slice1)
	fmt.Println("slice2 ==>", slice2)
	fmt.Println("array ==>", array)
	slice3 := append(slice1, array...)
	fmt.Println("slice3 -->", slice3)
	fmt.Println("array -->", array)

}

func main() {
	appendDemo()
	// sliceDemo3()
	// arrayDemo()
	// sliceDemo()
	// sliceDemo1()
	// sliceDemo2()
}
