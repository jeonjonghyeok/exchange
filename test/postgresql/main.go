package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "testdb"
	DB_PASSWORD = "5556"
	DB_NAME     = "testdb"
)

func main() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO test (name, age) VALUES('Jackass',19)")
	if err != nil {
		panic(err)
	}

	cntAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Affected Rows:", cntAffected)

}
