package main

import (
	"errors"
	"fmt"
)

type commonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能为负数")
	} else {
		return a + b, nil
	}
}

func add1(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{
			errorCode: 1,
			errorMsg:  "a或者b不能为负数"}
	} else {
		return a + b, nil
	}
}

func addDemo() {
	// err本质上是一个commentError指针类型，而commentError指针类型又实现了error接口
	//sum, err := add1(-1, 2)
	//cm,ok:=err.(*commonError)
	//cm,ok:=err.(*commonError)
	//if ok{
	//	fmt.Println("错误代码为:",cm.errorCode,"，错误信息为：",cm.errorMsg)
	//} else {
	//	fmt.Println("sum -->",sum)
	//}
	sum, err := add1(-1, 2)
	//cm,ok:=err.(*commonError)
	var cm *commonError
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}
}

func warpDemo() {
	e := errors.New("那种鸡")
	w := fmt.Errorf("那个蛋:%w", e)
	fmt.Println(w)

	//fmt.Println(errors.Unwrap(w))
	// 鸡 蛋代表嵌套error，鸡代表原来的错误。
	//	例如Is(a, b) --> a包含b则为true，(鸡 蛋，鸡)，而false，(鸡， 鸡 蛋)肯定不包含呀，判断第一个参数是否包含第二个参数。
	fmt.Println(errors.Is(w, e))
	fmt.Println(errors.Is(e, w))

}

func main() {
	//warpDemo()
	addDemo()
}
