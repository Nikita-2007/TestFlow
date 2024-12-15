package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	Name string
	Age  uint16
	Arr  []string
}

func (u *User) getUserName() string {
	return fmt.Sprintf(u.Name)
}

func GetUser() []User {
	path := "./model/data/db.sqlite"
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Println(err)
	}
	data := db.QueryRow(`
		SELECT * FROM 'User'
	`)
	fmt.Println(data)
	return data
}
