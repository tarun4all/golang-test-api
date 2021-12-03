package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	name string
}

func main() {
	fmt.Println("Mysql test")

	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("DB connected")

	insert, err := db.Query("INSERT INTO users VALUES('TARUN')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Data inserted")

	results, err := db.Query("SELECT name from users")

	for results.Next() {
		var user User

		err = results.Scan(&user.name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.name)
	}
}
