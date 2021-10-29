package main

import ( 
	"fmt"
	"io/ioutil"
)

func readTxt()  {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Printf("Error reading ", err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// contents, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	fmt.Printf("Error reading ", err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
}

func grade(score int) string {
	g := ""
	switch {
		case score < 0 || score > 100:
			panic(fmt.Sprintf(
				"Wrong score: %d", score))
		case score < 60:
			g = "F"
		case score < 80:
			g = "C"
		case score < 90:
			g = "B"
		case score < 100:
			g = "A"
		default:
			panic(fmt.Sprintf(
				"Wrong score: %d", score))
	}
	fmt.Println("g ==>", g)
	return g
}

func main() {
	// readTxt()
	grade(99)
}