package main

import "fmt"

func main() {
	// kondisi
	// keunggulan kondisi pada golang dapat membuat variable lokal untuk if dan else
	// nya dengan mendeklarasikan nya di awal sebelum kondisi
	var currentYear = 2023
	if age := currentYear - 1998; age < 17 {
		fmt.Println("Anda Dapat Membuat Kartu SIM")
	} else {
		fmt.Println("Anda Belum bisa membuat SIM")
	}

	// switch
	// switch pada golang tidak memerlukan break karena otomatis break, jika kita ingin menjalankan case selanjunya maka bisa
	// menggunakan fallthrough

	var score = 8
	switch score {
	case 8:
		fmt.Println("Nilai anda 8")
		fallthrough
	case 9:
		fmt.Println("Nilai anda 9")
	case 10:
		fmt.Println("Nilai anda 10")
	default:
		fmt.Println("Nilai anda selain itu")
	}

	// selain itu switch case pada golang dapat juga berisikan kondisi
	switch {
	case score > 8:
		fmt.Println("Nilai anda lebih dari 8")
		fallthrough
	case score < 8:
		fmt.Println("Nilai anda kurang dari 8")
	case score == 8:
		fmt.Println("Nilai anda 8")
	default:
		fmt.Println("Nilai anda selain itu")
	}

	// looping pada golang hanya for namun bisa diterapkan seperti while bahkan for bisa tanpa parameter
	// pada looping golang terdapat label sehingga jika break label maka for yang ditandai label akan berhenti
inilabel:
	for i := 0; i < 5; i++ {
		j := 0
		for j < 5 {
			fmt.Println(i, j)
			if j == 4 && i == 2 {
				break inilabel
			}
			j++
		}
	}

	// Array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%#v\n", numbers)

	// untuk mengeluarkan semua nilai di array dapat menggunakan loop through element
	for i, v := range numbers {
		fmt.Printf("indeks ke-%d : %d\n", i, v)
	}
	println("Panjang array number :", len(numbers))

	// slice
	// slice merupakan array tanpa penetapan banyak element, jika mengopi satu slice ke
	// sclie lain maka saat satu slice diganti nilainya, nilai pada slice lain juga akan terganti

	var fruits = []string{"apple", "banana", "manggo"}
	// bisa juga dibuat dengan
	var buah = make([]string, 3)
	buah[0] = "apel"
	buah[1] = "pisang"
	buah[2] = "mangga"

	// untuk menambahkan nilai pada slice dapate menggunakan function append
	buah = append(buah, "durian")
	// Jika kita ingin memasukkan seluruh element-element pada suatu array ke dalam array lainnya,
	// dapat menggunakan tanda ellipsis (...)
	var makanan []string
	makanan = append(makanan, fruits...)
	makanan = append(makanan, buah...)

	fmt.Printf("fruits : %#v\n", fruits)
	fmt.Printf("buah : %#v\n", buah)
	fmt.Printf("makanan : %#v\n", makanan)

	// slice copy function
	var fruits2 = []string{"", ""}
	nn := copy(fruits2, fruits)
	// fruits2 = fruits
	// fruits2[0] = "rambutan" jika seperti ini nilai fruits juga akan berubah
	fmt.Printf("fruits : %#v\n", fruits)
	fmt.Printf("fruits2 : %#v\n", fruits2)
	fmt.Printf("copied element : %#v\n", nn)

	// terdapat slicing di golang
	fmt.Println(fruits[1:3])
	fmt.Println(fruits[1:])
	fmt.Println(fruits[:2])
	fmt.Println(fruits[:])

	// MAP
	// map pada golang berfungsi untuk mengubah index menjadi nama atau tipe tertentu. seperti array assosiative di php atau object di js
	// map perlu di deklarasi dan di inisiasi
	var person map[string]string //Deklarasi
	person = map[string]string{} //inisiasi
	person["name"] = "Syah Fadel"
	person["NIM"] = "1810953019"
	person["age"] = "23"

	fmt.Println("name: ", person["name"])
	fmt.Println("NIM: ", person["NIM"])
	fmt.Println("age: ", person["age"])

	fmt.Println("Before deteling: ", person)

	delete(person, "NIM")

	fmt.Println("after deteling: ", person)

	// kita bisa cek apakah suatu key memiliki value

	value, exist := person["NIM"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	// Aliase merupakan sebuah fitu pada bahasa Go yang digunakan sebagai nama alternative daritipe data yang sudah ada.
	// Tipe data byte merupakan tipe data aliase dari tipe data uint8, yang berarti mereka merupakantipe data yang sama dengan nama yang berbda
	// cara membuat aliase
	type second = uint
	var nilai second = 10
	println(nilai)
}
