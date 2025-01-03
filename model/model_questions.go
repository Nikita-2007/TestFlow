package model

type Questions struct {
	Id             int
	NumberQuestion int
	Question       int
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
	for arr.Next() {
		print(arr)
	}
	/*query := `
		SELECT q.Id, q.NumberQuestion, q.Correct,
			q.Answer_1, q.Answer_2, q.Answer_3, q.Answer_4, q.Answer_5, q.Answer_6
		FROM TestQuestion tq, Questions q
		WHERE tq.Testing_id = $1
	;`
	//WHERE tq.Testing_id = $1 AND qt.Question_id = q.Id
	data, _ := db.Query(query, id)
	questions := []Questions{}
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
	*/
	db.Close()

	return &questions
}
