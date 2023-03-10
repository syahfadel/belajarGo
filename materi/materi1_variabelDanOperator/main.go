package main

import "fmt"

func main() {
	// dekalrasi variable dengan data type
	var nama string = "Syah Fadel"
	var umur int
	umur = 23

	// deklarasi variabel tanpa data type
	var nim = 1810953019
	// dapat di shorthand dengan cara
	jurusan := "Teknik Elektro"

	var kota, provinsi string
	kota, provinsi = "Tangerang", "Banten"

	var a, b, c = "1", 2, "3"

	// variabel _ untuk menempatkan variabel yang tidak digunakan, karena pada golang semua variabel harus digunakan

	/*
		ada beberapa cara untuk  menampilkan tulisan pada console yaitu
		• fmt.Print untuk menampilkan tulisan
		• fmt.Println untuk menampilkan tulisan dengan enter
		• fmt.Printf untuk menampilkan tulisan dengan format

		pada fmt.Print perlu menggunakan verb seperti %v %T %d dsb untuk verb yang ada dapat dilihat di
		https://pkg.go.dev/fmt
	*/
	_, _, _, _, _, _, _ = umur, jurusan, kota, provinsi, a, b, c
	fmt.Printf("Nama\t:%s\nNIM\t:%d", nama, nim)

	/*
		terdapat 4 kategori yaitu
		• basic Type 	 : number, string, boolean
		• Aggregate Type : array dan sruct
		• Reference Type : slice, pointer, map, function, channel
		• Interface type : interface
	*/

	/*
		*int (bilangan bulat)
		• int8, int16, int32, int64, uint8, uint16, uint32, uint64
		perbedaan antara ada u dan tidak yaitu pada uint merupakan bilangan bulat yang hanya bernilai positif angka
		dibelakang merupakan banyak memory atau bit dari type tersebut. defauiltnya int32 atau int64 tergantung nilainya

		*float (bilangan desimal)
		float dibagi menjadi dua jenis yaitu float32 dan 64
		jika menampilkan float akan menampilkan 6 angka dibelakang koma
		jika ingin mengatur bisa diatur dengan %.3f dengan 3 merupakan banyak angka dibelakang koma

		bool bernilai true or false
		string untuk kalimat diapit "" bisa juga dengan tanda `` yang mengakibatkan semua nilai berupa string seperti \n juga dianggep string

		konsonan dibuat dengan keywoard
		const
		pada golang terdapat iota untuk membuat multi konsonan dengan jarak tertentu
		contoh
		const (
			c1 = iota + 2
			c2 = iota + 2
			c3 = iota + 2
		)
		nilai c1 = 0, c2 = 2, c3 = 4

		operator pada java sama seperti golang

	*/
}
