package main

import "fmt"

func main() {
	//demo1()
	//popDemo()
	popDemo1()
}

func popDemo1() {
	s := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	sLen := len(s) - 1
	tail := s[sLen]
	s1 := s[:sLen]
	fmt.Println("s Len ==>", sLen, "\npop tail ==>", tail, "\ns1 ==> ", s1)
}

func demo1() {
	s := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("s ==>", s)
	s1 := append(s[:3], s[4:]...)
	fmt.Println("del element s1 ==>", s1)
}
