package main

import "fmt"

func main() {
	var banyak int
	fmt.Print("Enter expected max loop : ")
	fmt.Scan(&banyak)
	for i := 1; i <= banyak; i++ {
		fmt.Printf("%d ", i)
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println()
		}
	}
}
