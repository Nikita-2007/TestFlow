package main

import (
	"TestFlow/control"
	"TestFlow/model"
	"TestFlow/view"
	"encoding/json"
	"log"
	"os"
)

// Глобальные переменные (переделать на запрос с json)
/*
const (
	HOST string = "localhost"         //Адрес сервера
	PORT string = "8888"              //Порт сервера
	DBMS string = "sqlite3"           //Дравер СУБД
	PATH string = "./model/sqlite.db" //Расположение БД
	STAT string = "./view/"           //Расположение файлов
)
*/

type Config struct {
	HOST string
	PORT string
	DBMS string
	PATH string
	STAT string
}

func main() {
	// Читаем данные из файла setting.json
	jsonFile, err := os.ReadFile("setting.json")
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	// Создаем экземпляр структуры Config
	var config Config

	// Распарсиваем JSON в структуру config
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Fatalf("Ошибка распаковки JSON: %v", err)
	}
	//Создание или проверка базы данных
	model.Init(config.DBMS, config.PATH)

	//Статические файлы
	view.Static(config.STAT)

	//Запуск контроллеров
	control.Tg()
	t := control.MustToken()
	print(t)

	control.Web(config.HOST, config.PORT)
}
