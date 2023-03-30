package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "books"
)

var (
	Db  *sql.DB
	Err error
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, port, user, password, dbname)
	Db, Err = sql.Open("postgres", psqlInfo)
	if Err != nil {
		panic(Err)
	}
	Err = Db.Ping()
	if Err != nil {
		panic(Err)
	}

	fmt.Println("Successfully connected to database")
}
