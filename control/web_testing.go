package control

import (
	"TestFlow/model"
	"net/http"
	"strconv"
)

func testing(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Tests *[]model.Testing
	}{
		Title: "Тестирование",
		Tests: model.AllTests(),
	}
	tmpl := tmplFiles("view/test/testings.html")
	tmpl.ExecuteTemplate(w, "content", data)
}

// Контролер запроса создания теста
func createTesting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := struct {
			Title   string
			Courses *[]model.Course
		}{
			Title:   "Создание теста",
			Courses: model.AllCourses(),
		}
		tmpl := tmplFiles("view/test/create-test.html")
		tmpl.ExecuteTemplate(w, "content", data)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		course := r.FormValue("course")
		user := User
		description := r.FormValue("description")
		id := model.CreateTest(user, name, course, description)
		http.Redirect(w, r, "/get-test/"+strconv.Itoa(id), http.StatusSeeOther)
	}
}

// Контролер запроса изменения теста
func updateTesting(w http.ResponseWriter, r *http.Request) {
	id := getId(r.RequestURI)
	if r.Method == "GET" {
		data := struct {
			Title   string
			Testing *model.Testing
			Course  *[]model.Course
		}{
			Title:   "Редактирование теста",
			Testing: model.GetTest(id),
			Course:  model.AllCourses(),
		}
		tmpl := tmplFiles("view/test/update-test.html")
		tmpl.ExecuteTemplate(w, "content", data)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		course := r.FormValue("course")
		user := User
		description := r.FormValue("description")
		model.UpdateTest(id, user, name, course, description)
		http.Redirect(w, r, "/get-test/"+strconv.Itoa(id), http.StatusSeeOther)
	}
}

// Контролер запроса теста
func getTesting(w http.ResponseWriter, r *http.Request) {
	id := getId(r.RequestURI)
	arr := model.GetQuestions(id)
	count := len(*arr)
	if r.Method == "GET" {
		data := struct {
			Title     string
			Testing   *model.Testing
			Questions *[]model.Questions
			Сount     int
		}{
			Title:     "Тестирование",
			Testing:   model.GetTest(id),
			Questions: arr,
			Сount:     count,
		}
		tmpl := tmplFiles("view/test/get-test.html")
		tmpl.ExecuteTemplate(w, "content", data)
	} else if r.Method == "POST" {
		r.ParseForm()
		count, _ := strconv.Atoi(r.FormValue("count"))
		arr := []int{}
		for i := 1; i <= count; i++ {
			a, _ := strconv.Atoi(r.FormValue(strconv.Itoa(i)))
			if a > 0 {
				arr = append(arr, a)
			} else {
				arr = append(arr, 0)
			}
		}
		model.SetResult(User, id, arr)
		//Редирект на результат
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Контролер запроса удаления теста
func deleteTesting(w http.ResponseWriter, r *http.Request) {
	id := getId(r.RequestURI)
	model.DeleteTest(id)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
