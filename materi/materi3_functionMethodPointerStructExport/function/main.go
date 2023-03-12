package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n============Function tanpa Return============")
	greet("Syah Fadel", "Kota Tangerang", 23)

	fmt.Println("\n============Function dengan Return============")
	names := []string{"Syah", "Fadel", "Putra", "Dwingga"}
	msg := "Hallo perkenalkan nama saya"
	print := greetReturn(msg, names)
	fmt.Println(print)

	fmt.Println("\n============Function multiple Return============")
	var area, circumference = calculate(3, 5)
	fmt.Println("Area :", area)
	fmt.Println("Circumference :", circumference)

	fmt.Println("\n============Function Predefined Return Values============")
	var area2, circumference2 = calculatePredefined(5, 10)
	fmt.Println("Area :", area2)
	fmt.Println("Circumference :", circumference2)

	fmt.Println("\n============Function Varadic============")
	studentLists := studentList("Fadel", "Putra", "Rizky", "Putri", "Irin")
	fmt.Printf("%v\n", studentLists)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sum(numbers...)
	fmt.Println("Sum: ", sum)
}

// function tanpa return
func greet(name, kota string, age int) {
	fmt.Printf("Hallo, Perkenalkan saya %s dan umur saya saat ini %d tahun.\nSaya Tinggal di %s\n", name, age, kota)
}

// function dengan return. tipe data setelah tanda kurung merupakan tipe data yang dikembalikan
func greetReturn(msg string, names []string) string {
	joinStr := strings.Join(names, " ")          //string.Join digunakan untuk menggabungkan array string menjadi sebuah string
	result := fmt.Sprintf("%s %s", msg, joinStr) //untuk membuat string dengan format
	return result
}

// function multiple return
func calculate(length, width int) (int, int) {
	area := length * width
	circumference := 2 * (length + width)
	return area, circumference
}

// function multiple return
func calculatePredefined(length, width int) (area int, circumference int) {
	area = length * width
	circumference = 2 * (length + width)
	return
}

// varadic function
// merupakan suatu function yang dapat menerima argumen yang tidak terbatas
// dibuat dengan cara menggunakan parameter ...<tipe-data> cth ...string pada akhir parameter jika parameter lebih dari 1
func studentList(names ...string) []map[string]string {
	// name akan dalam bentuk slice
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("Student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

func sum(numbers ...int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}
