package main

import "fmt"

func ifDemo()  {
	i := 4
	if i > 10 {
		fmt.Println("i > 10")
	} else if i > 3 && i <= 10 {
		fmt.Println("3 < i <= 10")
	} else {
		fmt.Println("i <= 3")
	}
}

func switchDemo()  {
	switch i:=7;{
	case i>11:
		fmt.Println("i>11")
	case i>5 && i<=11:
		fmt.Println("5<i<=11")
	default:
		fmt.Println("i<=5")
	}
	
}

func switchDemo1()  {
	switch 3>1 {
	case true:
		fmt.Println("3>1")
	case false:
		fmt.Println("3<=1")
	}
}

func switchDemo2()  {
	switch j:=2;j {
	case 2:
		fallthrough
	case 3:
		fmt.Println("2")
	default:
		fmt.Println("没有匹配到")
	}
}

func forDemo()  {
	sum:=0
	for i:=1;i<=50;i++ {
		sum+=i
	}
	fmt.Println("the sum is",sum)
}

func forDemo1()  {
	sum := 0
	
	i:=50
	for {
		sum+=i
		i++
		if i>60 {
			break
		}
	}

	fmt.Println("the sum is",sum)
}

func main() {
	forDemo1()
	// forDemo()
	// switchDemo2()
	// switchDemo1()
	// switchDemo()
	// ifDemo()
}