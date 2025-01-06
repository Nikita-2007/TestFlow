package model

import (
	"time"
)

type Result struct {
	User    string
	Testing string
	Date    string
	Rate    int
}

// Смотрим результаттеста
func GetResult(user, test int) []Result {
	println(user, test)
	arr := []Result{}

	return arr
}

func GetAllResult(user int) []Result {
	println(user)
	arr := []Result{}

	return arr
}

// Вносим результат прохождения теста
func SetResult(user, test int, arr []int) {
	db := connect()
	date := time.Now().Format("2006-01-02")
	//Запрос верных ответов теста
	correctArr := []int{}
	query := `
		SELECT q.Correct
		FROM TestQuestion tq
		JOIN Questions q ON q.Id = tq.Question_id
		WHERE tq.Testing_id = $1
  	;`
	data, _ := db.Query(query, test)
	for data.Next() {
		answer := 0
		data.Scan(&answer)
		correctArr = append(correctArr, answer)
	}
	correctAnswer := 0
	for i := range correctArr {
		if correctArr[i] == arr[i] {
			correctAnswer += 1
		}
	}
	rate := correctAnswer * 100 / len(correctArr)
	query = `INSERT INTO UserTest (User_id, Testing_id, Date, Result)
		VALUES ($1, $2, $3, $4)`
	db.Exec(query, user, test, date, rate)
	db.Close()
}
