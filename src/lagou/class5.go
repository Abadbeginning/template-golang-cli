package main

import (
	"errors"
	"fmt"
)

func sumDemo(a int, b int) int {
	return a + b
}

func sum(a, b int, c string) int {
	return a + b
}

func sum1(a, b int) (int, int, error) {
	if a < 0 || b < 0 || a+b < 3 {
		return 0, 0, errors.New("a,b不能为负数，它们的和不能小于1")
	}
	
	return a + b, a, nil
}

func sum2(a, b int) (result int, result1 int, err error) {
	if a < 0 || b < 0 || a+b < 3 {
		return 0, 0, errors.New("a,b不能为负数，它们的和不能小于1")
	}
	result = a + b
	result1 = 1
	err = nil
	return
}

func sum3(args ...int) int {
	result := 0
	for _, value := range args {
		result += value
	}

	return result
}

func one() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Height uint

func (height Height) UToString() {
	fmt.Println("the man/woman height is", height)
}

func (height *Height) Change() {
	*height = Height(22)
}

func main() {
	height := Height(300)
	height.UToString()
	height.Change()
	height.UToString()

	//height := Height(300)
	//height.UToString()
	//o := one()
	//fmt.Println(o())
	//fmt.Println(o())
	//fmt.Println(o())
	//result := func(a, b int) int {
	//	return  a + b
	//}
	//
	//fmt.Println(result(1,2))

	//fmt.Println(sum3(1,2,3,4,5,6))
	//c, a, error := sum2(1,2)
	//fmt.Println(c,a,error)
	//c, a, error := sum1(1,2)
	//fmt.Println(c,a,error)
	//c1, a1, error1 := sum1(1,1)
	//fmt.Println(c1,a1,error1)
	//c := sumDemo(10, 20)
	//fmt.Println(c)
	//c1 := sum(10, 20, "123")
	//fmt.Println(c1)
}
