package main

import "fmt"

// method merupakan suatu function yang biasanya melekat pada suatu tipe data biasanya pada sebuah struct

type Person struct {
	name string
	age  int
}

// bisa juga memasukan struct dalam bentuk pointer
func (p Person) Introduce(msg string) string {
	return fmt.Sprintf("%s My Name is %s and I'm %d years old", msg, p.name, p.age)
}

func main() {
	var person = Person{name: "Syah Fadel", age: 23}
	fmt.Println(person.Introduce("Hello Everyone !!!"))
}
