package model

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	Dbms string
	Path string
)

func Init(DBMS, PATH string) {
	Dbms = DBMS
	Path = PATH

	_, err := os.Stat(Path)
	if err != nil && os.IsNotExist(err) {
		tableCourse()
		DataCourse() //Заполняем
		tableUser()
	} else {
		backupDB()
	}
}

func tableUser() {
	db, err := sql.Open(Dbms, Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS User (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		Avatar TEXT DEFAULT "./view/img/avatar.png",
		Password TEXT,
		Email TEXT,
		DataBirth TEXT,
		Course INTEGER,
		DataReg TEXT,
		Status TEXT,
		Rate INTEGER
	)`
	data, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.RowsAffected()) //Сообщение о создании БД или что она уже есть
}

func tableCourse() {
	db, err := sql.Open(Dbms, Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS Course (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		Icon TEXT,
		Background TEXT,
		Description TEXT
	)`
	data, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.RowsAffected()) //Сообщение о создании БД или что она уже есть
}
