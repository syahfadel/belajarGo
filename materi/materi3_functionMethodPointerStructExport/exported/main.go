package main

import "exported/helpers"

func main() {
	helpers.Greet() //method bisa diakses dari package berbeda karena dibuat dengan awal kapital
	// helpers.greet() tidak bisa karena diawali huruf kecil
	var person = helpers.Person{}
	// person.Name("Syah Fadel")
	// person.Age(23)
	// person.Address("Tangerang")

	person.Invokegreet() //dapat mengakses greet melalui method yang diawali kapital
}
