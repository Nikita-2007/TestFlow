package model

import "time"

// Структура данных пользователей
type User struct {
	Id        int
	Name      string
	Avatar    string
	Password  string
	Email     string
	DataBirth string
	Course    string
	DataReg   string
	Status    string
	Rate      int
}

// Вывод всех пользователей
func AllUser() *[]User {
	db := connect()
	query := "SELECT * FROM User"
	data, _ := db.Query(query)
	users := []User{}
	for data.Next() {
		user := User{}
		data.Scan(
			&user.Id,
			&user.Name,
			&user.Avatar,
			&user.Password,
			&user.Email,
			&user.DataBirth,
			&user.Course,
			&user.DataReg,
			&user.Status,
			&user.Rate,
		)
		users = append(users, user)
	}
	db.Close()

	return &users
}

// Вывод данных пользователя
func GetUser(id int) *User {
	db := connect()
	query := "SELECT * FROM User WHERE id = $1"
	data := db.QueryRow(query, id)
	user := User{}
	data.Scan(
		&user.Id,
		&user.Name,
		&user.Avatar,
		&user.Password,
		&user.Email,
		&user.DataBirth,
		&user.Course,
		&user.DataReg,
		&user.Status,
		&user.Rate,
	)
	course := Course{}
	query = `SELECT Name FROM Course WHERE id = $1`
	data = db.QueryRow(query, &user.Course)
	data.Scan(&course.Name)
	user.Course = course.Name

	db.Close()
	return &user
}

// Добавление пользователя
func AddUser(name string, password string, email string, dataBirth string) int {
	db := connect()
	dataReg := time.Now().Format("2006-01-02")
	avatar := "./view/img/avatar.png"
	query := "INSERT INTO User (Name, Avatar, Password, Email, DataBirth, Course, DataReg, Status, Rate) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	result, _ := db.Exec(query, avatar, name, password, email, dataBirth, 2, dataReg, "active", 0)
	db.Close()
	id, _ := result.LastInsertId()
	return int(id)
}

// Обновление пользователя
func UpdateUser(id int, name string, password string, email string, dataBirth string) {
	db := connect()
	query := "UPDATE User SET Name=?, Password=?, Email=?, DataBirth=? WHERE id = ?"
	db.Exec(query, name, password, email, dataBirth, id)
	db.Close()
}

// Удаление пользователя
func DeleteUser(id int) {
	db := connect()
	query := "DELETE FROM User WHERE id = $1"
	db.Exec(query, id)
	db.Close()
}
