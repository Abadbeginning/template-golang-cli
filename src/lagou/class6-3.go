package main

type users1 struct {
	id        uint
	usersName string
	addr      addr
}

type addr1 struct {
	city string
}

func NewPerson(name string) *users1 {
	return &users1{usersName: name}
}

func main() {

}
