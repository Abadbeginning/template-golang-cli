package main

import "fmt"

func updateSlice(s []int) {
	s[2] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
	fmt.Println("=================================")
	s1 := arr[2:]
	s2 := arr[:]
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Extengding slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))
	//fmt.Println("s2 = ", s2)
	s3 := append(s2, 100)
	s4 := append(s3, 101)
	s5 := append(s4, 101)
	fmt.Println("s3 s4 s5 = ", s3, s4, s5)
	fmt.Println("arr = ", arr)

}
