package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var i int = 10
	// 自动推导类型
	// var (
	// 	j = 1
	// 	k = 12
	// 	i = 122
	// )

	// fmt.Print("i => ", i)
	// fmt.Print("j => ", j)
	// fmt.Print("k => ", k)

	// var f32 float32 = 2.2
	// var f64 float64 = 10.3456
	// fmt.Println("f32 is",f32,",f64 is",f64)

	// var bf bool =false

	// var bt bool = true

	// fmt.Println("bf is",bf,",bt is",bt)

	// var s1 string = "Hello"

	// var s2 string = "世界"

	// fmt.Println("s1 is",s1,",s2 is",s2)
	// fmt.Println("s1 + s2 is ", s1+s2)

	// var zi int

	// var zf float64

	// var zb bool

	// var zs string

	// fmt.Println(zi,zf,zb,zs)

	i:= 10
	// bf:= false
	// s4:= "hello"
	// fmt.Println(i,bf)
	// fmt.Println(s4)

	// 指针
	// pi:=&i
	// fmt.Println("&pi ==> ", pi)
	// fmt.Println("pi ==> ", *pi)

	i = 20
	fmt.Println("i的新值是",i)

	// 常量
	const name = "小米"
	fmt.Println("name ==> ", name)

	// 
	// const(

	// 	one = 1
	
	// 	two = 2
	
	// 	three =3
	
	// 	four =4
	
	// )

	// iota 是一个常量生成器
	const (

		one = iota+1
	
		two
	
		three
	
		four
	
	)
	
	fmt.Println(one,two,three,four)
	
	// 字符串和数字互转
	i2s:= strconv.Itoa(i)
	s2i, err:=strconv.Atoi(i2s)
	fmt.Println(s2i, i2s, err)

	f64:= 12.222222
	i2f:=float64(i)

	f2i:=int(f64)

	fmt.Println(i2f,f2i)

	//判断s1的前缀是否是H
	s1:="123rerth"
	fmt.Println(strings.HasPrefix(s1,"H"))

	//在s1中查找字符串o

	fmt.Println(strings.Index(s1,"o"))

	//把s1全部转为大写

	fmt.Println(strings.ToUpper(s1))



}