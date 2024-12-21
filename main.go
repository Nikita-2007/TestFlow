package main

import (
	"TestFlow/control"
	"TestFlow/view"
)

// Глобальные переменные(Переделать на json)
const (
	HOST string = "localhost"         //Адрес сервера
	PORT string = "8888"              //Порт сервера
	DBMS string = "sqlite3"           //Дравер СУБД
	PATH string = "./model/sqlite.db" //Расположение БД
	STAT string = "./view/"           //Расположение файлов
)

func main() {
	//Создание или проверка базы данных
	//model.Init(DBMS, PATH)

	//Статические файлы
	view.Static(STAT)

	//Запуск контроллеров
	control.Tg()
	control.Web(HOST, PORT)
}
