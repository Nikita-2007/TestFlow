package model

// Добавление пользователя
func DemoUser() {
	db := connect()
	query := `
	INSERT INTO User(Avatar, Name, Password, Email, DataBirth, Course, DataReg, Status, Rate) VALUES
		('./view/img/avatar.png', 'Вася Нечейный', '123456543211286756', 'vasya2000@gmail.com', '01-01-2001', 'Физика', '01-01-2024', 'active', 0),
		('./view/img/avatar.png', 'Екатерина Иванова', 'StrongP@$d123', 'katya.i@exale.com', '15-03-1998', 'Алгоритмизация', '29-12-2024', 'active', 75),
		('./view/img/avatar.png', 'Александр Петров', 'PasswOrd@12', 'aleks.petrov@mail.ru', '22-04-1990', 'Математика', '25-12-2024', 'active', 92),
		('./view/img/avatar.png', 'Ольга Смирнова', 'SecurePw!3', 'olga.smirnova@yandex.ru', '10-11-2001', 'Литература', '01-01-2025', 'banned', 12),
		('./view/img/avatar.png', 'Иван Кузнецов', 'StrongPass123', 'ivan.kuznetsov@gmail.com', '05-07-1985', 'Физика', '27-12-2024', 'active', 65),
		('./view/img/avatar.png', 'Анастасия Васильева', 'MySecretP@ss', 'anastasia.vas@mail.ru', '18-02-2003', 'Химия', '26-12-2024', 'active', 34),
		('./view/img/avatar.png', 'Сергей Иванов', 'Secret123!', 'sergei.ivanov@example.com', '29-09-1978', 'Философия', '30-12-2024', 'banned', 88),
		('./view/img/avatar.png', 'Елена Соколова', 'Pa$$wOrd1', 'elena.sokolova@yandex.ru', '01-05-1999', 'Алгоритмизация', '20-12-2024', 'active', 52),
		('./view/img/avatar.png', 'Дмитрий Попов', 'MySecureP@ss', 'dmitry.popov@gmail.com', '12-08-1992', 'Математика', '28-12-2024', 'active', 19),
		('./view/img/avatar.png', 'Мария Лебедева', 'Secure123!', 'maria.lebedeva@mail.ru', '17-06-2000', 'Литература', '21-12-2024', 'active', 77),
		('./view/img/avatar.png', 'Андрей Сидоров', 'SuperSecretPw', 'andrey.sidorov@example.com', '03-03-1980', 'Физика', '22-12-2024', 'active', 9),
		('./view/img/avatar.png', 'Наталья Морозова', 'Passw0rd#123', 'natalia.morozova@gmail.com', '25-10-2004', 'Химия', '29-12-2024', 'active', 45);
	`
	db.Exec(query)
	db.Close()
}