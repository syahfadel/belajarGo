package main

import (
	"database/sql"
	"fmt"

	// menggunakan tanda _ karena tidak akan menggunakan syntax apapun dari github.com/lib/pq
	// mengimport driver postgresql untuk meregistrasikan dengan mengimport
	// supaya database/sql mengetahui driver apa yang dipakai
	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "db-go-sql"
)

// berupa variabel global
var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	// menggabungkan seluruh info dari postgresql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, port, user, password, dbname)

	// menggunakan function Open dari database/sql yang menerima 2 parameter yaitu nama driver yang dipakai
	// dan string berupa informasi bagaimana cara database/sql membangun koneksi dengan database kita
	// function sql.Open akan mengembalikan nilai berupa pointer dari struct sql.DB
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// db.Ping merupakan method dari struct sql.DB yang direturn sql.Open
	// method Ping sangat penting untuk membangun koneksi ke database sekaligus memeriksa jika string
	// panjang berupa info yang kita berikan merupakan info yang 100% valid
	// function open tidak membangun koneksi ke database, melainkan hanya untuk memvalidasi argumen yang diberikan
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	GetEmployees()
	// UpdateEmployee()
	// GetEmployees()
	// DeleteEmployee()
	// GetEmployees()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees(full_name,email,age,division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	// db.QueryRow method untuk mengeksekusi query sql, dan hanya untuk satu return atau satu row
	err = db.QueryRow(sqlStatement, "Syah Fadel Putra Dwingga", "syahfadel@email.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
		// Scan untuk mendapatkan nilai balikan dari statement karena menggunakan Returning *
		// disimpan kedalam field dari variable employee

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	//method query biasa digunakan untuk mendapatkan banyak data dari suatu table pada database
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// rows.Next() akan memberikan nilai true jika masih ada data
	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas: ", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, division = $4, age =$5
	WHERE id = $1;
	`
	// QueryRow juga dapat digunakan namun dianjurkan untuk method Exec untuk insert update delete karena Exec
	// tidak mendapatkan data yang baru saja diupdate maupun dimuat.
	res, err := db.Exec(sqlStatement, 1, "Syah Fadel", "syah@email.com", "FinTech", 23)

	if err != nil {
		panic(err)
	}

	// RowsAffected untuk mengetahui jumlah row atau data yang diupadte
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated data amount", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employees
	WHERE id = $1
	`

	res, err := db.Exec(sqlStatement, 2)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data Amount: ", count)
}
