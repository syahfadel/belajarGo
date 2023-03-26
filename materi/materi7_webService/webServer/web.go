package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/", greet) //untuk keperluan routing yang menerima 2 parameter, parameter 1 merupakan pathrouting dan kedua merupakan
	//function dengan 2 parameter yaitu http.ResponseWriter dan http.Request

	http.ListenAndServe(PORT, nil)
	// digunakan untuk menjalankan server aplikasi. Menerima 2 parameter yaitu port yang digunakan dan http.Handler yang merupakan sebuah
	// interface. Namun jika tidak digunakan bisa menggunakan nil
}

/*
http.ResponseWriter
merupakan interface dengan berbagai method yang digunakan untuk mengirim response balik kepada client yang mengirim request
http.Request
merupakan sebuah struct yang digunakan untuk mendapat berbagai macam data request seperti form value, headers, url parameter dan lain-lain
*/

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello Worlds"
	fmt.Fprint(w, msg)
}
