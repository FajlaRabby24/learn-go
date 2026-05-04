package main

import (
	"fmt"
)

var x = 10

// variable scrope
func variable() {
	x = 30
	fmt.Println(x)
}

func variable2() {
	fmt.Println(x)
}

func double(num *int) {
	*num = *num * 2
}

func main() {
	//* variable scope
	variable()
	variable2()

	//* array
	var students = [5]int{23, 43, 21, 67, 54}
	var totalMarks int
	for i := 0; i < len(students); i++ {
		element := students[i]
		totalMarks += element
	}
	fmt.Printf("Total marks: %d\n", totalMarks)

	//* slice
	slice := []string{"rice", "oil", "sugar"}
	slice = append(slice, "salt")
	target := "oil"
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		if element == target {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(slice)

	//* for loop
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	//* if-else
	temparature := 1
	var result string
	if temparature < 0 {
		result = "Freezing"
	} else if temparature <= 15 {
		result = "cold"
	} else if temparature <= 30 {
		result = "Normal"
	} else {
		result = "Hot"
	}
	fmt.Println(result)

	//* switch
	weekDayNum := 5
	var switchResult string
	switch weekDayNum {
	case 1:
		switchResult = "Saturday"
	case 2:
		switchResult = "Sunday"
	case 3:
		switchResult = "Monday"
	case 4:
		switchResult = "Tuesday"
	case 5:
		switchResult = "Wednesday"
	case 6:
		switchResult = "Thursday"
	case 7:
		switchResult = "Friday"
	default:
		switchResult = "Invalid day"
	}
	fmt.Println(switchResult)

	//* struct
	type student struct {
		name string
		age  int
		gpa  float32
	}
	student1 := student{name: "Fajla", age: 50, gpa: 2.45}
	student2 := student{name: "Fajla", age: 51, gpa: 2.46}
	student3 := student{name: "Fajla", age: 52, gpa: 2.47}
	fmt.Println(student1, student2, student3)

	//* pointer
	a := 10
	double(&a)
	fmt.Println(a)

	//* scan
	var userName string
	var userAge int
	fmt.Print("Enter you name: ")
	fmt.Scan(&userName)
	fmt.Print("Enter you age: ")
	fmt.Scan(&userAge)

	fmt.Printf("Hello %s, you are %d years old", userName, userAge)

	//* struct + slice
	studentSlice := []student{}
	studentSlice = append(studentSlice, student{name: "Fajla", age: 50, gpa: 2.45})
	studentSlice = append(studentSlice, student{name: "John", age: 25, gpa: 3.80})
	studentSlice = append(studentSlice, student{name: "Jane", age: 22, gpa: 3.90})

	var higherGpaStudent student
	for i := 0; i < len(studentSlice); i++ {
		candidate := studentSlice[i]
		if candidate.gpa > higherGpaStudent.gpa {
			higherGpaStudent = candidate
		}
	}
	fmt.Println(higherGpaStudent)

}
