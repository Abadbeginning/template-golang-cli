package main

import (
	"errors"
	"fmt"
	"strconv"
)

func errorDemo() {
	i, err := strconv.Atoi("b")
	//i,err:=strconv.Atoi("10")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

func addHeight(heightMin, heightMax int) (int, error) {
	if heightMin < 150 || heightMax < 180 {
		return 150, errors.New("heightMin不能小于150,heightMax不能小于180")
	} else {
		return heightMin + heightMax, nil
	}
}

func addHeightShow() {
	hSum, error := addHeight(140, 181)
	if error == nil {
		fmt.Println(hSum)
	} else {
		fmt.Println(error)
	}
}

type universalError struct {
	// 错误状态
	errorStatus int
	// 错误信息
	errorMsg string
}

func (ue *universalError) Error() string {
	return ue.errorMsg
}

func addHeight1(heightMin, heightMax int) (int, error) {
	if heightMin < 150 || heightMax < 180 {
		return 150, &universalError{
			errorStatus: 400,
			errorMsg:    "heightMin不能小于150,heightMax不能小于180",
		}
	} else {
		return heightMin + heightMax, nil
	}
}

func assertAddHeight1() {
	sum, error1 := addHeight1(149, 181)
	var ue *universalError
	if errors.As(error1, &ue) {
		fmt.Println("错误状态 -->", ue.errorStatus, "\n错误信息 -->", ue.errorMsg)
	} else {
		fmt.Println(sum)
	}

	//sum, error := addHeight1(149,181)
	//ue, result := error.(*universalError)
	//if result {
	//	fmt.Println("错误状态 -->", ue.errorStatus,"\n错误信息 -->", ue.errorMsg)
	//} else {
	//	fmt.Println(sum)
	//}
}

func main() {
	assertAddHeight1()
	//fmt.Println(addHeight1(149,181))
	//addHeightShow()
	//fmt.Println(addHeight(150,181))
	//errorDemo()
}
