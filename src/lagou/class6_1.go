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
    show() string
}

func (u user) show() string {
	fmt.Println("the username is -->",u.userName," |id is -->",u.id)
	return "ok"
}

func printString(u Stringer) {
	fmt.Println(u.show())
}

func structDemo()  {
	u := user{userName:"小明", id: 1, addr : address{
		city: "美国",
	}}
	printString(u)
	printString(&u)
	
}

func main() {
	structDemo()
}