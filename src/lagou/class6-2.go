package main

import "fmt"

type Mover interface{
    eat()
}

type cat struct{}

func (c cat) eat1() {
    fmt.Println("猫吃鱼")
}

func show()  {
	var x Mover
	var c = cat{}
	// x = cat{}
	x = c         
	x.eat()
	var c1 =&cat{}
	x = c1   
	x.eat()
}

func main() {
	show()
}