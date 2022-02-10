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
    Show() string
}

func (u user) Show() string {
	fmt.Println("the username is -->",u.userName," |id is -->",u.id)
	return "ok"
}

func printString(u Stringer) {
	fmt.Println(u.Show())
}

func structDemo()  {
	u := user{userName:"小明", id: 1, addr : address{
		city: "美国",
	}}
	printString(u)
	
}

func main() {
	structDemo()
}