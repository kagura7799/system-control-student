package main

import (
	"fmt"
)

type Student struct {
	ID         int
	Name       string
	Surname    string
	Patronymic string
}

var (
	name, surname, patronymic, command string
	ID                                 int
)

func main() {
	fmt.Println(
		`+-------------------------------+
	|                        		 |
	| Система управления студентами  |
	|                                |
	+--------------------------------+`,
	)
	students := make(map[int]Student)

	students[1] = Student{ID: 1, Name: "Максим", Surname: "Котельников", Patronymic: "Викторович"}
	students[2] = Student{ID: 2, Name: "Александр", Surname: "Воронов", Patronymic: "Андреевич"}
	students[3] = Student{ID: 3, Name: "Борис", Surname: "Аристов", Patronymic: "Петрович"}
	students[4] = Student{ID: 4, Name: "Федор", Surname: "Астафьев", Patronymic: "Максимович"}
	students[5] = Student{ID: 5, Name: "Настя", Surname: "Шмыкова", Patronymic: "Игорьевна"}

	fmt.Println("Выберите действие: del/add/print/find")
	for {
		_, err := fmt.Scan(&command)

		if err != nil {
			fmt.Println("Произошла ошибка: ", err)
			return
		}

		switch command {

		case "del":
			RemoveStudents(students)
		case "add":
			AddStudents(students)
		case "print":
			PrintStudents(students)
		case "find":
			FindStudents(students)

		default:
			fmt.Println("Неверная команда, попробуйте еще раз.")
		}
	}
}

func AddStudents(students map[int]Student) {
	fmt.Println("Пожалуйста, укажите ID, ФИО студента:")
	_, err := fmt.Scanf("%d %s %s %s", &ID, &surname, &name, &patronymic)

	if err != nil {
		fmt.Println("Произошла ошибка: ", err)
		return
	}

	fmt.Printf("Студент %s %s %s успешно добавлен в базу.\n", surname, name, patronymic)

	newStudent := Student{ID: ID, Name: name, Surname: surname, Patronymic: patronymic}
	students[newStudent.ID] = newStudent
}

func RemoveStudents(students map[int]Student) {
	if len(students) == 0 {
		fmt.Println("Список студентов пуст.")
	} else {
		fmt.Println("Пожалуйста, укажите ID студента:")

		_, err := fmt.Scan(&ID)

		studentFound := false
		for _, val := range students {
			if ID == val.ID {
				studentFound = true
				break
			}
		}

		if !studentFound {
			fmt.Printf("Студента с ID %d не существует.\n", ID)
		} else {
			fmt.Printf("Студент %s %s %s успешно удален из базы.\n", students[ID].Surname, students[ID].Name, students[ID].Patronymic)
			delete(students, ID)
		}

		if err != nil {
			fmt.Println("Произошла ошибка: ", err)
		}
	}
}

func FindStudents(students map[int]Student) {
	if len(students) == 0 {
		fmt.Println("Список студентов пуст.")
	}

	fmt.Println("Введите ID студента которого собираетесь найти:")
	_, err := fmt.Scanf("%d", &ID)

	if err != nil {
		fmt.Println("Произошла ошибка: ", err)
	}

	studentFound := false
	for _, val := range students {
		if ID == val.ID {
			fmt.Printf("Студент под ID %d: %s %s %s\n", val.ID, val.Surname, val.Name, val.Patronymic)
			studentFound = true
			break
		}
	}

	if !studentFound {
		fmt.Printf("Студент с ID %d не найден.\n", ID)
	}
}

func PrintStudents(students map[int]Student) {
	if len(students) == 0 {
		fmt.Println("Список студентов пуст.")
	}

	for id, student := range students {
		fmt.Printf("ID: %d, %s %s %s\n", id, student.Surname, student.Name, student.Patronymic)
	}

}
