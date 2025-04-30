package main 

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	age int
}

func (s student) sayHello() {
	fmt.Println("Hallo ", s.name)
	fmt.Printf("Umur saya %d tahun\n", s.age)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ") [i-1]
}

func main() {
	var s1 = student{"aaaaa yaudah iya", 21}
	s1.sayHello()

	var name = s1.getNameAt(1)
	fmt.Println("Name Panggilan: ", name)
}