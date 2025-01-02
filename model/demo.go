package model

// Добавление пользователя
func Demo() {
	db := connect()

	//Пользователи
	query := `
	INSERT INTO User(Avatar, Name, Password, Email, DataBirth, Course_id, DataReg, Status, Rate) VALUES
		('../view/img/avatar.png', 'Вася Нечейный', '123456543211286756', 'vasya2000@gmail.com', '01-01-2001', 1, '01-01-2024', 'active', 0),
		('../view/img/avatar.png', 'Екатерина Иванова', 'StrongP@$d123', 'katya.i@exale.com', '15-03-1998', 2, '29-12-2024', 'active', 75),
		('../view/img/avatar.png', 'Александр Петров', 'PasswOrd@12', 'aleks.petrov@mail.ru', '22-04-1990', 3, '25-12-2024', 'active', 92),
		('../view/img/avatar.png', 'Ольга Смирнова', 'SecurePw!3', 'olga.smirnova@yandex.ru', '10-11-2001', 4, '01-01-2025', 'banned', 12),
		('../view/img/avatar.png', 'Иван Кузнецов', 'StrongPass123', 'ivan.kuznetsov@gmail.com', '05-07-1985', 5, '27-12-2024', 'active', 65),
		('../view/img/avatar.png', 'Анастасия Васильева', 'MySecretP@ss', 'anastasia.vas@mail.ru', '18-02-2003', 6, '26-12-2024', 'active', 34),
		('../view/img/avatar.png', 'Сергей Иванов', 'Secret123!', 'sergei.ivanov@example.com', '29-09-1978', 7, '30-12-2024', 'banned', 88),
		('../view/img/avatar.png', 'Елена Соколова', 'Pa$$wOrd1', 'elena.sokolova@yandex.ru', '01-05-1999', 8, '20-12-2024', 'active', 52),
		('../view/img/avatar.png', 'Дмитрий Попов', 'MySecureP@ss', 'dmitry.popov@gmail.com', '12-08-1992', 9, '28-12-2024', 'active', 19),
		('../view/img/avatar.png', 'Мария Лебедева', 'Secure123!', 'maria.lebedeva@mail.ru', '17-06-2000', 10, '21-12-2024', 'active', 77),
		('../view/img/avatar.png', 'Андрей Сидоров', 'SuperSecretPw', 'andrey.sidorov@example.com', '03-03-1980', 11, '22-12-2024', 'active', 9),
		('../view/img/avatar.png', 'Наталья Морозова', 'Passw0rd#123', 'natalia.morozova@gmail.com', '25-10-2004', 1, '29-12-2024', 'active', 45)
	;`
	data, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	print("База данных пользователей заполнена: ")
	println(data.RowsAffected())

	//Вопросы
	query = `INSERT INTO Questions(Testing_id, NumberQuestion, Question, Correct, Answer_1, Answer_2, Answer_3, Answer_4, Answer_5, Answer_6) VALUES
		(1, 1, 'Какая из перечисленных единиц измерения является единицей измерения скорости?', 3, 'метр', 'секунда', 'метр в секунду', 'ньютон', 'килограмм', 'паскаль'),
		(1, 2, 'Какой закон Ньютона описывает инерцию?', 1, 'Первый', 'Второй', 'Третий', 'Закон всемирного тяготения', 'Закон сохранения энергии', 'Закон сохранения импульса'),
		(1, 3, 'Какое уравнение описывает равномерное прямолинейное движение?', 4, 's = gt²/2', 'F = ma', 'E = mc²', 's = vt', 'v = at', 's = v₀t + at²/2'),
		(1, 4, 'Что такое ускорение?', 2, 'Изменение положения тела', 'Изменение скорости тела', 'Сила, действующая на тело', 'Масса тела', 'Плотность тела', 'Импульс тела'),
		(1, 5, 'Какая сила вызывает движение тела по окружности?', 3, 'Сила трения', 'Сила тяжести', 'Центростремительная сила', 'Сила упругости', 'Сила Архимеда', 'Сила Ампера'),

		(2, 1, 'Какой тип реакции описывается общим уравнением A + B → AB?', 1, 'Соединение', 'Разложение', 'Замещение', 'Обмен', 'Нейтрализация', 'Гидролиз'),
		(2, 2, 'Как называется реакция между кислотой и основанием?', 4, 'Окисление', 'Восстановление', 'Разложение', 'Нейтрализация', 'Гидратация', 'Полимеризация'),
		(2, 3, 'Какой коэффициент нужно поставить перед O₂ в уравнении 2H₂ + ?O₂ → 2H₂O, чтобы уравнять его?', 2, '1', '2', '3', '4', '0.5', '1.5'),
		(2, 4, 'Какой из перечисленных признаков не указывает на протекание химической реакции?', 5, 'Выделение газа', 'Образование осадка', 'Изменение цвета', 'Выделение тепла', 'Изменение объема', 'Появление запаха'),
		(2, 5, 'Что такое катализатор?', 3, 'Вещество, замедляющее реакцию', 'Продукт реакции', 'Вещество, ускоряющее реакцию', 'Вещество, изменяющее равновесие', 'Реагент', 'Растворитель'),

		(3, 1, 'Какая планета Солнечной системы является самой большой?', 3, 'Марс', 'Венера', 'Юпитер', 'Земля', 'Сатурн', 'Меркурий'),
		(3, 2, 'Как называется естественный спутник Земли?', 1, 'Луна', 'Фобос', 'Деймос', 'Европа', 'Ганимед', 'Титан'),
		(3, 3, 'Какой планетой Солнечной системы являются «красной планетой»?', 4, 'Юпитер', 'Венера', 'Уран', 'Марс', 'Нептун', 'Меркурий'),
		(3, 4, 'Что такое астероид?', 2, 'Комета', 'Небольшое каменное тело', 'Звезда', 'Планета', 'Газовое облако', 'Черная дыра'),
		(3, 5, 'Вокруг чего вращаются планеты в Солнечной системе?', 1, 'Солнце', 'Земля', 'Луна', 'Юпитер', 'Центр галактики', 'Другие планеты'),

		(4, 1, 'Что такое экспозиция в литературном произведении?', 1, 'Начало произведения, вводящее в курс дела', 'Развязка событий', 'Кульминация сюжета', 'Описание места действия', 'Диалог героев', 'Внутренний монолог героя'),
		(4, 2, 'Какой элемент композиции следует за экспозицией?', 2, 'Эпилог', 'Завязка', 'Кульминация', 'Развязка', 'Пролог', 'Отступление'),
		(4, 3, 'Что такое метафора?', 3, 'Сравнение двух предметов', 'Преувеличение', 'Перенос значения на основе сходства', 'Обращение к кому-либо', 'Гипербола', 'Олицетворение'),
		(4, 4, 'Что является главным элементом сюжета?', 4, 'Описание пейзажа', 'Второстепенные персонажи', 'Лирическое отступление', 'Последовательность событий', 'Авторские ремарки', 'Внутренние переживания героя'),
		(4, 5, 'Какой художественный прием используется в фразе "золотые руки"?', 5, 'Сравнение', 'Эпитет', 'Ирония', 'Гипербола', 'Метонимия', 'Анафора'),

		(5, 1, 'Какой органеллой является митохондрия?', 2, 'Рибосома', 'Энергетическая станция', 'Хлоропласт', 'Эндоплазматическая сеть', 'Ядро', 'Аппарат Гольджи'),
		(5, 2, 'Где в клетке хранится генетическая информация?', 5, 'Митохондрия', 'Рибосома', 'Лизосома', 'Цитоплазма', 'Ядро', 'Вакуоль'),
		(5, 3, 'Какую функцию выполняют рибосомы?', 3, 'Синтез липидов', 'Хранение воды', 'Синтез белков', 'Транспорт веществ', 'Деление клетки', 'Удаление отходов'),
		(5, 4, 'Какой органеллой растительной клетки является хлоропласт?', 4, 'Центросома', 'Аппарат Гольджи', 'Лизосома', 'Осуществляет фотосинтез', 'Эндоплазматическая сеть', 'Вакуоль'),
		(5, 5, 'Что такое цитоплазма?', 1, 'Внутренняя среда клетки', 'Клеточная стенка', 'Мембрана клетки', 'Оболочка ядра', 'Межклеточное вещество', 'Хромосома'),

		(6, 1, 'Какое устройство является основным "мозгом" компьютера?', 2, 'Оперативная память (RAM)', 'Центральный процессор (CPU)', 'Жесткий диск (HDD/SSD)', 'Видеокарта (GPU)', 'Материнская плата', 'Блок питания'),
		(6, 2, 'Что такое RAM?', 4, 'Постоянная память', 'Энергонезависимая память', 'Внешняя память', 'Оперативная память', 'Флеш-память', 'Архивная память'),
		(6, 3, 'Для чего используется жесткий диск (HDD/SSD)?', 1, 'Для постоянного хранения данных', 'Для обработки графики', 'Для временного хранения данных', 'Для вывода информации на экран', 'Для работы с сетью', 'Для охлаждения'),
		(6, 4, 'Какая часть компьютера отвечает за обработку графики?', 3, 'Центральный процессор (CPU)', 'Оперативная память (RAM)', 'Видеокарта (GPU)', 'Жесткий диск (HDD/SSD)', 'Звуковая карта', 'Сетевая карта'),
		(6, 5, 'Что является программным обеспечением?', 5, 'Монитор', 'Клавиатура', 'Мышь', 'Принтер', 'Операционная система', 'Процессор'),

		(7, 1, 'Какой тип алгоритма выполняется последовательно, шаг за шагом?', 1, 'Линейный', 'Разветвляющийся', 'Циклический', 'Рекурсивный', 'Параллельный', 'Вложенный'),
		(7, 2, 'Какой блок на блок-схеме обозначает условие?', 2, 'Прямоугольник', 'Ромб', 'Овал', 'Параллелограмм', 'Круг', 'Трапеция'),
		(7, 3, 'Какой тип алгоритма предполагает повторение одних и тех же действий?', 3, 'Линейный', 'Разветвляющийся', 'Циклический', 'Случайный', 'Рекурсивный', 'Последовательный'),
		(7, 4, 'Какой алгоритмической структурой является "если-то-иначе"?', 4, 'Линейной', 'Циклической', 'Рекурсивной', 'Разветвляющейся', 'Параллельной', 'Сложной'),
		(7, 5, 'Что такое блок-схема?', 5, 'Описание кода программы', 'Набор данных', 'Последовательность команд', 'Текстовое описание алгоритма', 'Графическое представление алгоритма', 'Список переменных'),

		(8, 1, 'Какая река является самой длинной в мире?', 3, 'Амазонка', 'Янцзы', 'Нил', 'Миссисипи', 'Обь', 'Енисей'),
		(8, 2, 'Какой горный хребет является самым высоким в мире?', 1, 'Гималаи', 'Анды', 'Кордильеры', 'Альпы', 'Уральские горы', 'Аппалачи'),
		(8, 3, 'Столицей какой страны является город Токио?', 4, 'Китай', 'Южная Корея', 'Индия', 'Япония', 'Вьетнам', 'Таиланд'),
		(8, 4, 'Какой океан является самым большим по площади?', 2, 'Индийский', 'Тихий', 'Атлантический', 'Северный Ледовитый', 'Южный', 'Антарктический'),
		(8, 5, 'Какая пустыня является самой крупной в мире?', 5, 'Калахари', 'Атакама', 'Гоби', 'Аравийская', 'Сахара', 'Намиб'),

		(9, 1, 'В каком году началась Вторая мировая война?', 2, '1937', '1939', '1941', '1945', '1914', '1918'),
		(9, 2, 'Какое событие произошло в 1492 году?', 4, 'Основание Рима', 'Падение Константинополя', 'Начало Столетней войны', 'Открытие Америки Колумбом', 'Французская революция', 'Первая публикация Библии'),
		(9, 3, 'В каком веке произошла Французская революция?', 3, 'XVII', 'XVIII', 'XIX', 'XX', 'XVI', 'XV'),
		(9, 4, 'Когда была провозглашена независимость США?', 1, '1776', '1789', '1812', '1861', '1865', '1918'),
		(9, 5, 'Какой год считается годом окончания Первой мировой войны?', 5, '1914', '1915', '1916', '1917', '1918', '1919'),

		(10, 1, 'Чему равно значение x в уравнении 2x + 5 = 11?', 3, '2', '4', '3', '6', '8', '1'),
		(10, 2, 'Какая функция является линейной?', 2, 'y = x²', 'y = 2x + 3', 'y = 1/x', 'y = √x', 'y = x³', 'y = eˣ'),
		(10, 3, 'Какой знак нужно поставить вместо * в неравенстве 5 * 3, чтобы оно было верным?', 1, '>', '<', '=', '≤', '≥', '≠'),
		(10, 4, 'Как называется выражение x² + 2xy + y²?', 4, 'Разность квадратов', 'Сумма кубов', 'Квадрат разности', 'Квадрат суммы', 'Куб суммы', 'Неполный квадрат разности'),
		(10, 5, 'Чему равна производная функции f(x) = x³?', 5, '3', 'x²', '2x', 'x³/2', '3x²', '1/x'),

		(11, 1, 'Что такое переменная в программировании?', 2, 'Тип данных', 'Именованная область памяти для хранения данных', 'Оператор', 'Функция', 'Цикл', 'Условие'),
		(11, 2, 'Какой оператор используется для присваивания значения переменной?', 1, '=', '==', '!=', '+', '-', '×'),
		(11, 3, 'Что такое цикл в программировании?', 4, 'Оператор присваивания', 'Оператор сравнения', 'Условный оператор', 'Многократное выполнение блока кода', 'Объявление переменной', 'Вывод на экран'),
		(11, 4, 'Какой из перечисленных вариантов не является языком программирования?', 5, 'Python', 'Java', 'C++', 'JavaScript', 'HTML', 'PHP'),
		(11, 5, 'Для чего предназначены функции в программировании?', 3, 'Хранение данных', 'Управление потоком выполнения', 'Для выделения блоков кода', 'Для определения типов переменных', 'Для выполнения математических операций', 'Для вывода на экран')
	;`
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	print("База данных вопросов заполнена: ")
	println(data.RowsAffected())

	//Тесты
	query = `INSERT INTO Testing(Name, Course_id, User_id, Description) VALUES
		('Механика и движение', 1, 1, 'Основы кинематики и динамики, законы Ньютона'),
		('Химические реакции', 2, 2, 'Основные типы реакций, баланс, свойства веществ'),
		('Солнечная система', 3, 3, 'Планеты, спутники, астероиды, строение и движение'),
		('Анализ литературного произведения', 4, 4, 'Сюжет, композиция, образы, средства выразительности'),
		('Клетка и её строение', 5, 5, 'Структура клеток, органеллы, процессы жизнедеятельности'),
		('Основы компьютерной архитектуры', 6, 6, 'Аппаратное и программное обеспечение, принципы работы ПК'),
		('Структуры алгоритмов', 7, 7, 'Линейные, разветвляющиеся, циклические алгоритмы, блок-схемы'),
		('Географическая номенклатура', 8, 8, 'Рельеф, климат, природные зоны, страны и столицы'),
		('Хронология событий', 9, 9, 'Основные исторические периоды, ключевые даты, личности'),
		('Основы алгебры', 10, 10, 'Уравнения, неравенства, функции, графики, преобразования'),
		('Основы языков программирования', 11, 11, 'Синтаксис, переменные, циклы, функции, операторы')
	;`
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	print("База данных тестов заполнена: ")
	println(data.RowsAffected())

	db.Close()
}
