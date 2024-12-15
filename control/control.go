package control

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  uint16
	Arr  []string
}

func index(w http.ResponseWriter, r *http.Request) {
	data := User{"Bob", 9, []string{"98", "93", "77", "82", "83"}}
	page, err := template.ParseFiles(
		"view/index.html",
		"view/inc/header.html",
		"view/inc/menu.html",
		"view/inc/footer.html",
	)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	page.ExecuteTemplate(w, "content", data)
}

func about(w http.ResponseWriter, r *http.Request) {
	data := 16
	page, err := template.ParseFiles(
		"view/about.html",
		"view/inc/header.html",
		"view/inc/menu.html",
		"view/inc/footer.html",
	)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	page.ExecuteTemplate(w, "content", data)
}

func contact(w http.ResponseWriter, r *http.Request) {
	data := 7
	page, err := template.ParseFiles(
		"view/contact.html",
		"view/inc/header.html",
		"view/inc/menu.html",
		"view/inc/footer.html",
	)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	page.ExecuteTemplate(w, "content", data)
}

func Control() {
	fmt.Println("Запуск сервера")

	//Статические файлы
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("./view/"))))

	//Урлы страниц
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)

	//Слушатель порта
	http.ListenAndServe("localhost:8888", nil)
}
