package main

import "fmt"

// go tidak menyediakan tipe data class tetapi terdapat yang namanya struct

type Employee struct {
	name     string //permission label public pada go yaitu dengan mengawali variabel dengan huruf kapital Name
	age      int
	division string
}

// struct yang berisi struct disebut embeded struct
type Person struct {
	name string
	age  int
}

type Student struct {
	nim    int
	major  string
	person Person
}

func main() {
	var employee1 Employee
	employee1.name = "Syah Fadel"
	employee1.age = 23
	employee1.division = "Financial Service"

	var employee2 = Employee{name: "Putra", age: 22, division: "merchant"}

	fmt.Println(employee1.name)
	fmt.Println(employee1.age)
	fmt.Println(employee1.division)
	fmt.Printf("%+v\n", employee1)
	fmt.Println(employee2.name)
	fmt.Println(employee2.age)
	fmt.Println(employee2.division)
	fmt.Printf("%+v\n", employee2)

	// kita juga bisa menggunakan pointer pada sebuah struct
	var employee3 *Employee = &employee2
	fmt.Println("\nEmployee2 Name(before): ", employee2.name)
	fmt.Println("Employee3 Name(before): ", employee3.name)
	employee3.name = "Dwingga"
	fmt.Println("Employee2 Name(after): ", employee2.name)
	fmt.Println("Employee3 Name(after): ", employee3.name)

	var student Student
	student.person.name = "Syah Fadel"
	student.person.age = 23
	student.nim = 1810953019
	student.major = "Electrical Engineering"
	fmt.Printf("\n%+v", student)

	// anonymous struct
	var employee4 = struct {
		name     string
		age      int
		division string
	}{name: "Fadel", age: 21, division: "Backend"}
	fmt.Printf("\n%+v\n\n", employee4)

	// slice to struct
	var people = []Person{
		{name: "Syah", age: 21},
		{name: "Fadel", age: 22},
		{name: "Putra", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
