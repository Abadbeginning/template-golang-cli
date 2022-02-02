package main

import "fmt"

func int_demo()  {
    fmt.Println("Hello 整型")
	var a int64 = -10
	var b uint32 = 190
	fmt.Print(" a => ", a, "\n b = ", b)
}

// func int_demo1()  {
//     fmt.Println("Hello 整型")
// 	var a int64 = 10
// 	var b uint32 = -1
// 	fmt.Print(" a => ", a, "\n b = ", b)
// }

func float_demo()  {
	var a float32 = 8.88
	var b float64 = 18.88888
	fmt.Println("a == ",a,",b ==",b)
}

func strDemo()  {
	var s string = "Hello"
	var s1 string = "Go"
	fmt.Println("s ==",s,",s1 ==",s1)
	var s2 = s + s1
	fmt.Println("s2 ==",s2)

}

func boolDemo()  {
	var b bool = true
	var b1 bool = false
	fmt.Println("b ==",b,",b1 ==",b1)

}

func main() {
	// int_demo()
	// int_demo1()
	// float_demo()
	// strDemo()
	boolDemo()
}