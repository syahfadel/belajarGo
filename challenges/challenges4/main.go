package main

import (
	"fmt"
	"os"
)

type Data struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func (d Data) getAllData() {
	fmt.Printf("Nama\t\t: %s\n", d.nama)
	fmt.Printf("Alamat\t\t: %s\n", d.alamat)
	fmt.Printf("Pekerjaan\t: %s\n", d.pekerjaan)
	fmt.Printf("Alasan\t\t: %s\n", d.alasan)
}

func main() {
	absen := os.Args[1][0] - 48
	var data = []Data{
		{nama: "Fadel",
			alamat:    "Kota Tangerang",
			pekerjaan: "Freshgraduate",
			alasan:    "Mempelajari web service",
		},
		{nama: "Taufik",
			alamat:    "Pasaman",
			pekerjaan: "Freshgraduate",
			alasan:    "Mempelajari golang",
		},
		{nama: "Rival",
			alamat:    "Surabaya",
			pekerjaan: "Freshgraduate",
			alasan:    "Mendalami materi backend",
		},
	}
	data[absen].getAllData()
}
