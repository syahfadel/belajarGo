package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Data yang akan diambil
type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

// Data Dummy
var employees = []Employee{
	{ID: 1, Name: "Fadel", Age: 23, Division: "FinTech"},
	{ID: 2, Name: "Putra", Age: 24, Division: "Infra"},
	{ID: 3, Name: "Dwi", Age: 22, Division: "Flight"},
	{ID: 4, Name: "Ingga", Age: 21, Division: "Hotel"},
}

// Port yang akan digunakan
var PORT = ":8080"

func main() {
	http.HandleFunc("/getEmployees", getEmployees)
	http.HandleFunc("/addEmployee", addEmployee)
	fmt.Println("Application is listening on port ", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	// menentukan bentuk data response yang akan dikirim kepada client, disini dalam bentuk JSON
	w.Header().Set("Content-Type", "Application/json")

	// untuk tampil apda html

	if r.Method == "GET" { //mengecek method yang digunakan
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// json.NewEncoder(w).Encode(employees)
		// mengonversi data employees ke bentuk json dengan Encode
		// lalu data tersebut dikirim dengan NewEncoder

		tpl.Execute(w, employees)
		return
	}

	// akan menampilkan error jika bukan get
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, _ := strconv.Atoi(age)

		// if err != nil {
		// 	http.Error(w, "Invalid age", http.StatusBadRequest)
		// 	return
		// }

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
