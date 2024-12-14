package control

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	page := "<h1>Главная</h1><div><a href='/home'>Домой</a></div>"
	fmt.Fprint(w, page)
}

func home(w http.ResponseWriter, r *http.Request) {
	page := "<p><b>Сервер работает</b></p><p><a href='/'>На главную</a></p><a href='/about'>О проекте</a>"
	fmt.Fprint(w, page)
}

func about(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("view/about.html")
	page.Execute(w, page)
}

func Control() {
	fmt.Println("Запуск сервера")
	//Урлы страниц
	http.HandleFunc("/", index)
	http.HandleFunc("/home/", home)
	http.HandleFunc("/about/", about)

	//Слушатель порта
	http.ListenAndServe("localhost:8888", nil)
}
