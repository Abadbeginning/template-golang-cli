package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "mouse",
		"courese": "go",
		"site":    "imooc",
		"quality": "notbad",
	}

	fmt.Println(m)

	// m2 == empty map
	m2 := make(map[string]int)

	fmt.Println(m2)

	// m3 == nil
	var m3 map[string]int

	fmt.Println(m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	courseName1 := m["coures"]
	fmt.Println(courseName1)

	courseName2, ok := m["course"]
	fmt.Println(courseName2, ok)

	if causeName4, ok1 := m["name"]; ok1 {
		fmt.Println(causeName4)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
