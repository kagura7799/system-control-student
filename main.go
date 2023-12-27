package main

import (
	"fmt"
)

type Student struct {
	ID      int
	Name    string
	Surname string
}

var name, surname string
var ID int

func main() {
	fmt.Println(
		`+-------------------------------+
	|                        		 |
	| Система управления студентами  |
	|                                |
	+--------------------------------+`,
	)
	students := make(map[int]Student)

	var command string
	fmt.Println("Выберите действие: del/add/check")
	for {
		_, err := fmt.Scan(&command)

		if err != nil {
			fmt.Println("Произошла ошибка: ", err)
			return
		}

		switch command {
		case "del":
			RemoveStudent(students)
		case "add":
			AddStudent(students)
		case "check":
			PrintStudents(students)

		default:
			fmt.Println("Неверная команда, попробуйте еще раз.")
		}
	}
}

func AddStudent(students map[int]Student) {
	fmt.Println("Пожалуйста, укажите ID, Имя, Фамилию студента:")
	_, err := fmt.Scanf("%d %s %s", &ID, &name, &surname)

	if err != nil {
		fmt.Println("Произошла ошибка: ", err)
		return
	} else {
		fmt.Printf("Студент %s %s успешно добавлен в базу.\n", name, surname)
	}

	newStudent := Student{ID: ID, Name: name, Surname: surname}
	students[newStudent.ID] = newStudent
}

func RemoveStudent(students map[int]Student) {
	fmt.Println("Пожалуйста, укажите ID студента:")
	_, err := fmt.Scanf("%d", &ID)

	if err != nil {
		fmt.Println("Произошла ошибка: ", err)
	} else {
		fmt.Printf("Студент %s %s успешно удален из базы.\n", students[ID].Name, students[ID].Surname)
	}

	delete(students, ID)
}

func PrintStudents(students map[int]Student) {
	if len(students) == 0 {
		fmt.Println("Список студентов пуст.")
	}

	for id, student := range students {
		fmt.Printf("ID: %d, Имя: %s, Фамилия: %s\n", id, student.Name, student.Surname)
	}
}
