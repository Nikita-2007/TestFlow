package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Главная</h1><div><a href='/home'>Домой</a></div>")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<b>Сервер работает</b>")
}

func control() {
	fmt.Println("Запуск сервера")
	//Урлы страниц
	http.HandleFunc("/", index)
	http.HandleFunc("/home/", home)

	//Слушатель порта
	http.ListenAndServe("localhost:8888", nil)
}

func main() {
	control()
}
