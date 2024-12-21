package model

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	Id   int
	Name string
	Age  uint16
}

func AddUser(name string, age int) {
	db, err := sql.Open(Dbms, Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Exec("INSERT INTO User (Name, Age) VALUES ($1, $2)", name, age)
	if err != nil {
		panic(err)
	}
	fmt.Print(data.RowsAffected())
	fmt.Println(data.LastInsertId())
}

func AllUser() []user {
	path := "./model/data/db.sqlite"
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	data := []user{}

	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data = append(data, u)
	}
	return data
}

// Данные пользователя  ПЕРЕДОДЕЛАТЬ
func GetUser(id int) *user {
	db, err := sql.Open(Dbms, Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT * FROM User WHERE id = $1"
	data := db.QueryRow(query, id)

	user := user{}
	err = data.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		panic(err)
	}

	return &user
}
