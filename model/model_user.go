package model

// Структура данных пользователей
type User struct {
	Id           int
	Avatar       string
	Name         string
	Email        string
	DataBirth    string
	Course       string
	DataReg      string
	TestsPassed  string
	TestsCreated string
	Status       string
	Rate         int
}

// Вывод всех пользователей
func AllUser() *[]User {
	db := connect()
	query := "SELECT * FROM User"
	data, _ := db.Query(query)
	users := []User{}
	for data.Next() {
		user := User{}
		data.Scan(&user.Id, &user.Name, &user.DataReg)
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
		&user.Avatar,
		&user.Name,
		&user.Email,
		&user.DataBirth,
		&user.Course,
		&user.DataReg,
		&user.Status,
		&user.Rate,
	)
	db.Close()

	return &user
}

// Добавление пользователя
func AddUser(name string, age int) int {
	db := connect()
	query := "INSERT INTO User (Name, DataReg) VALUES ($1, $2)"
	result, _ := db.Exec(query, name, age)
	db.Close()
	id, _ := result.LastInsertId()
	return int(id)
}

// Обновление пользователя
func UpdateUser(id int, name string, age int) {
	db := connect()
	query := "UPDATE User SET Name=?, Age=? WHERE id = ?"
	db.Exec(query, name, age, id)
	db.Close()
}

// Удаление пользователя
func DeleteUser(id int) {
	db := connect()
	query := "DELETE FROM User WHERE id = $1"
	db.Exec(query, id)
	db.Close()
}
