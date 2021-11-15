package main

import "fmt"

func enums()  {
	const (
		cpp = iota
		_
		python
		golang
		js
	)
	fmt.Println(cpp,js,python,golang)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	enums()
}