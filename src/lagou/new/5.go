package main

import (
	"errors"
	"fmt"
)

func test() {
	// aa := [3][3]int{}
	// aa[0][0] = 1
	// aa[0][1] = 2
	// aa[0][2] = 3
	// aa[1][0] = 4
	// aa[1][1] = 5
	// aa[1][2] = 6
	// aa[2][0] = 7
	// aa[2][1] = 8
	// aa[2][2] = 9
	// fmt.Println(aa)
	// result, err := sum(-1, 22)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }
	result := sum("sdf", -1, 22, 22, 3, 5, 76, 879)
	fmt.Println(result)
}

// func sum(a int, b int) int {
func sum1(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	return a + b, nil
}
func sum2(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}
func sum(tip string, params ...int) int {
	fmt.Println(tip, params)
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

func test1() {
	// sum2 := func(a, b int) int {
	// 	return a + b
	// }
	// fmt.Println(sum2(1, 2))
    cl:=colsure()
    fmt.Println(cl())
    fmt.Println(cl())
    fmt.Println(cl())

}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Age uint
func (age Age) String(){
    fmt.Println("the age is",age)
}
func (age *Age) Modify(){
	*age = Age(22)
}

func test2()  {
	age := Age(233)
	age.String()
	// age.Modify()
	(&age).Modify()
	age.String()
}

func main() {
	// test()
	// test1()
	test2()
}
