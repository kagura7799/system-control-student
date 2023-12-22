package main
import (
	"fmt"
)
 
type Student struct {
	ID int
	Name string
}

func main() {
	fmt.Println(
	`+-------------------------------+
	|                        		 |
	| Система контроля за студентами |
	|                                |
	+--------------------------------+`,
)
	students := make(map[int]Student)

	// var command string;
	
	// switch command {
	// case "del": RemoveStudent()
	// case "add": AddStudent()
	// case "check": PrintStudents()
	// }
	// fmt.Println("Выберите действие: del/add/check")
	// fmt.Scan(&command)

    AddStudent(students, Student{ID: 1, Name: "Володя"})

    fmt.Println("Список студентов:")
    PrintStudents(students)
	fmt.Println()

    RemoveStudent(students, 2)

    fmt.Println("Обновленный список студентов:")
    PrintStudents(students)
}


func AddStudent(students map[int]Student, newStudent Student) {
    students[newStudent.ID] = newStudent
}

func RemoveStudent(students map[int]Student, ID int) {
	delete(students, ID)
}
  
func PrintStudents(students map[int]Student) {
	for id, student := range students {
		fmt.Printf("ID: %d, Имя: %s\n", id, student.Name)
	}
}