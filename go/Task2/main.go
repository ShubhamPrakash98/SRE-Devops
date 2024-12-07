package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	student1 := addStudent("amar")
	addStudent("akhbar")
	addStudent("anthony")

	showAllStudents()
	deleteStudentByID(student1.ID.String())
	fmt.Println("-----------After deletion---------")
	showAllStudents()

}

var students []*student

type student struct {
	ID    uuid.UUID
	Name  string
	Marks map[string]int
}

var idToStudent = make(map[uuid.UUID]*student)

func addStudent(name string) *student {
	subjectToMarks := map[string]int{
		"English":   0,
		"Maths":     0,
		"Physics":   0,
		"Chemistry": 0,
		"Biology":   0,
	}
	id := uuid.New()
	student := &student{
		ID:    id,
		Name:  name,
		Marks: subjectToMarks,
	}
	students = append(students, student)
	idToStudent[id] = student
	return student
}

func updateStudentName(id uuid.UUID, newName string) error {
	if _, ok := idToStudent[id]; !ok {
		return errors.New("the student not found!")
	}

	idToStudent[id].Name = newName
	return nil
}

func deleteStudentByID(id string) error {
	Uid := uuid.MustParse(id)
	if _, ok := idToStudent[Uid]; !ok {
		return errors.New("the student not found!")
	}

	var newArr []*student
	for _, student := range students {
		if student.ID != Uid {
			newArr = append(newArr, student)
		}
	}
	students = newArr
	return nil
}

func showAllStudents() {
	for _, student := range students {
		fmt.Println("-----------------")
		fmt.Println(student.ID)
		fmt.Println(student.Name)
		fmt.Println(student.Marks)
	}

}
