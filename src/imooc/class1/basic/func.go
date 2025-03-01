package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div1(a, b int) (q, r int) {
	return a / b, a % b
}

// int类型的函数 入参 函数 op(int, int) 返回值int (int, int) a, b
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function %s with args "+"(%d, %d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(eval(3, 5, "*"))
	s, b := div1(10, 2)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	a1, b1 := 13, 4
	a1, b1 = swap1(a1, b1)
	fmt.Println(a1, b1)

}
