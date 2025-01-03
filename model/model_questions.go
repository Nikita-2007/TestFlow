package model

type Questions struct {
	Id             int
	NumberQuestion int
	Question       string
	Correct        int
	Answer_1       string
	Answer_2       string
	Answer_3       string
	Answer_4       string
	Answer_5       string
	Answer_6       string
}

// Получаем список вопросов
func GetQuestions(id int) *[]Questions {
	db := connect()
	query := `
		SELECT Question_id
		FROM TestQuestion
		WHERE Testing_id = $1
	;`
	arr, _ := db.Query(query, id)
	questions := []Questions{}
	var idq string
	for arr.Next() {
		arr.Scan(&idq)
		query = `
		SELECT Id, NumberQuestion, Question, Correct,
				Answer_1, Answer_2, Answer_3, Answer_4, Answer_5, Answer_6
			FROM Questions
			WHERE Id = $1
		;`
		data, _ := db.Query(query, idq)
		for data.Next() {
			question := Questions{}
			data.Scan(
				&question.Id,
				&question.NumberQuestion,
				&question.Question,
				&question.Correct,
				&question.Answer_1,
				&question.Answer_2,
				&question.Answer_3,
				&question.Answer_4,
				&question.Answer_5,
				&question.Answer_6,
			)
			questions = append(questions, question)
		}
	}
	db.Close()

	return &questions
}
