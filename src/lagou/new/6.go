package main

import "fmt"

type person struct {
	name string
	age uint
	addr address
}

type address struct {
    province string
    city string
}

func test() {
	// var p person
	// p := person{"小李子", 23}
	p := person{name: "小李子"}
	// p := person{age: 23, name: "小李子"}
	fmt.Println(p)
	fmt.Println(p.name, p.age)
}

type Stringer interface {
	String() string
}

func (p person) String() string  {
	return fmt.Sprintf("the name is %s, age is %d ", p.name, p.age)
}

func printString(s fmt.Stringer)  {
	fmt.Println(s.String())
}

func test1()  {
	p := person {
		age: 30,
		name: "时代感",
		addr: address {
			province: "上升到",
			city: "三大个人",
		},
	}
	// addr_res := p.addr
	// fmt.Println(p, addr_res.province, addr_res.city)
	printString(p)
}

func test2()  {
	
}

func main() {
	// test()
	test1()
	// test2()
}
