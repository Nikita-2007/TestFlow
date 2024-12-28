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
		createTable()
	} else {
		backupDB()
	}
}

func createTable() {
	db, err := sql.Open(Dbms, Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS User (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		Avatar TEXT,
		Name TEXT,
		Email TEXT,
		DataBirth TEXT,
		Course TEXT,
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
