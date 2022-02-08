package main

import "fmt"

type user struct {
	id uint
	userName string
	addr address
	*other
}

type address struct {
	city string
}

type other struct {
	sex uint
}

func structDemo()  {
	u := user{userName:"小明", id: 1, addr : address{
		city: "美国",
	}}
	fmt.Println(u)
	fmt.Println(u.addr.city)

	// var u user
	// fmt.Println(u)
	// u := user{userName:"小明", id: 1}

	// u := user{1, "小明"}
	// fmt.Println(u.id,u.userName)
	
}

func main() {
	structDemo()
}