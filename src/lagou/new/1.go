package main

import (
	"fmt"
	"strconv"
)

func test() {
	i := 111
	pi := &i
	fmt.Println(pi)
	// const name = [1, 2, 3]
	i2s := strconv.Itoa(i)

	s2i, err := strconv.Atoi(i2s)

	fmt.Println(i2s, s2i, err)

}

func main() {
	test()

}
