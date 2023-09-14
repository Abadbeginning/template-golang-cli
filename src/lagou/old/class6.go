package main

import "fmt"

type user struct {
	id uint
	userName string
	addr address
}

type address struct {
	city string
}

type Stringer interface {
    String() string
}

func (u user) String() string {
	return fmt.Sprintf("the username is %s,id is %d",u.userName,u.id)
}

func printString(user fmt.Stringer){
    fmt.Println(user.String())
}

func structDemo()  {
	u := user{userName:"小明", id: 1, addr : address{
		city: "美国",
	}}

	// printString(u)

	// 值接收者和指针接收者
	printString(&u)
	// fmt.Println(u)
	// fmt.Println(u.addr.city)

	// var u user
	// fmt.Println(u)
	// u := user{userName:"小明", id: 1}

	// u := user{1, "小明"}
	// fmt.Println(u.id,u.userName)
	
}

func main() {
	structDemo()
}