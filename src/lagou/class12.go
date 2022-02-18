package main

import "fmt"

func pointDemo()  {
	age := 18
	ageP := &age
	var age1 *int
	age1 = &age
	fmt.Println(age1,*age1)
	fmt.Println("age值 --> ", age)
	fmt.Println("age指针值 --> ", ageP)
	fmt.Println("*age值 --> ", *ageP)
	// var strP *string 
	// strP = &age
	

}

func main() {
	pointDemo()
}