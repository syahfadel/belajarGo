package main

import (
	"fmt"
)

func main() {
	var firstNumber int = 10
	var secondNumber *int = &firstNumber

	fmt.Println("First Number (Value): ", firstNumber)
	fmt.Println("First Number (Address Memory): ", &firstNumber)
	fmt.Println("Second Number (Pointer Value): ", *secondNumber)
	fmt.Println("Second Number (value): ", secondNumber)
	fmt.Println("Second Number (Address Memory): ", &secondNumber)

	println("\n==mengganti nilai first number melalui second number==\n")

	*secondNumber = 20

	fmt.Println("First Number (Value): ", firstNumber)
	fmt.Println("First Number (Address Memory): ", &firstNumber)
	fmt.Println("Second Number (Pointer Value): ", *secondNumber)
	fmt.Println("Second Number (value): ", secondNumber)
	fmt.Println("Second Number (Address Memory): ", &secondNumber)

	println("\n========Memanfaatkan pointer untuk function========\n")
	a := 10
	fmt.Println("Before :", a)
	changeValue(&a) //yang menjadi nilai alamat memory
	fmt.Println("After :", a)

}

//pointer dapat dimanfaatkan sebagai parameter function
// sehingga saat nilai pointer pada  function diubah maka nilai pada variabel diargumen juga berubah

func changeValue(number *int) {
	*number = 100
}
