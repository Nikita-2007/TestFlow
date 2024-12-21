package model

import (
	"database/sql"
	"fmt"
)

var (
	Dbms string // Тип СУБД
	Path string // Путь к БД
)

// Создание БД(одноразово при запуске)
func Init(dbms, path string) {
	Dbms = dbms
	Path = path

	//Проверка на существование БД и ее создание при необходимости
	createTable() //ДОДЕЛАТЬ
}

func createTable() {
	db, err := sql.Open(Dbms, Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "CREATE TABLE IF NOT EXISTS User (Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT, Age INTEGER)"
	data, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.RowsAffected()) //Сообщение о создании БД или что она уже есть
}
