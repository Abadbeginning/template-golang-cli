package main

import "fmt"

type users struct {
	id        uint
	usersName string
	addr      addr
}

type addr struct {
	city string
}

type Stringser interface {
	show() string
}

func (u *users) show() string {
	fmt.Println("the usersname is -->", u.usersName, " |id is -->", u.id)
	return "ok"
}

func (addr addr) show() string {
	fmt.Println("the addr is -->", addr.city)
	return "ok"
}

//func (u users) show() string {
//	fmt.Println("the usersname is -->",u.usersName," |id is -->",u.id)
//	return "ok"
//}

func printStr(u Stringser) {
	a, ok := u.(addr)
	if ok {
		fmt.Println("addr -->", a)
	} else {
		fmt.Println("不是addr")
	}

	b, ok1 := u.(*users)
	if ok1 {
		fmt.Println("*users -->", b)
	} else {
		fmt.Println("不是*users")
	}

	fmt.Println(u.show())
}

func structShow() {
	u := users{usersName: "小明", id: 1, addr: addr{
		city: "美国",
	}}

	address := addr{city: "美国"}

	printStr(address)
	printStr(&u)

}

func main() {
	structShow()
}
