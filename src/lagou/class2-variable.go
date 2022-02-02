package main

import "fmt"

var i int = 20

var (
	a = 1
	b = 2
	c = 3
)


e := 12

func zeroVarDemo()  {
	var zv int
	var zv1 float64
	var zv2 bool
	var zv3 string
	
	fmt.Println(zv3,zv,zv1,zv2)
	fmt.Println(zv,zv1,zv2,zv3)
	
}

func variableDemo()  {
	var i int = 20
	var j = 20
	k := 20
	fmt.Println("i ==>", i," j ==>", j," k ==>", k)
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println("a ==>", a," b ==>", b," c ==>", c)
	e, f, d := 11, 12, 13
	fmt.Println("e ==>", e," f ==>", f,"d  ==>", d)
}

func main() {
	variableDemo()
	// zeroVarDemo()
}