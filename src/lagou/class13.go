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

func printStr(u Stringser) string {
	//fmt.Println(u.show())
	return u.show()
}

func structShow() {
	//u := users{usersName: "小明", id: 1, addr: addr{
	//	city: "美国",
	//}}

	address := addr{city: "美国"}

	printStr(address)
	fmt.Println("-------------------1-----------------------")
	printStr(&address)
	//printStr(&u)
	fmt.Println("-------------------2-----------------------")
	sip := printStr(address)
	fmt.Println(sip)
	fmt.Println("-------------------3-----------------------")

	sip1 := &printStr(address)
	fmt.Println(sip1)
	fmt.Println("-------------------4-----------------------")

}

func main() {
	structShow()
}
