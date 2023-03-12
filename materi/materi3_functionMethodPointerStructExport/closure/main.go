package main

import (
	"fmt"
	"strings"
)

func main() {
	// closure merupaakn sebuah anonymous function atau function tanpa nama yang dapat disimpan sebagai sebuah variable

	// Closure pada variable
	fmt.Println("\n================Closure pada Variabel================")
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 858, 46, 28, 62, 35}

	fmt.Println(evenNumbers(numbers...))

	// IIFE merupakan closure yang langsung dieksekusi saat dideklarasikan
	fmt.Println("\n================Closure IIFE================")
	var isPalindrom = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(str[i]) //str[i] mengembalikan kode ascii
		}

		return temp == str
	}("katak") // lansung diberikan nilai untuk argument

	fmt.Println(isPalindrom)

	fmt.Println("\n================Closure sebagai return================")
	var studentList = []string{"Fadel", "Putra", "Widya", "Irin", "Freya"}
	var find = findStudent(studentList)
	fmt.Println(find("Fadel"))

	fmt.Println("\n================Closure (Callback)================")
	var findOdd = findOddNumber(numbers, func(number int) bool {
		return number%2 == 1
	})
	fmt.Println("Total Odd Number: ", findOdd)
}

// closure sebagai nilai return
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s doesn't exist!!!", s)
		}

		return fmt.Sprintf("We found %s at position %d", student, position)
	}
}

// closure sebagai callback
// callback merupakan sebuah closure yang dijadikan sebagai paramter pada function lain
func findOddNumber(numbers []int, callback func(int) bool) int {
	var totalOddNumber int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumber++
		}
	}
	return totalOddNumber
}
