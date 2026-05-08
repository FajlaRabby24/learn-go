package main

/**
// * Student struct বানাও — Name, Marks (slice), Grade
// Marks এর average বের করো
// Average অনুযায়ী Grade দাও:
// 80+  → "A+"
// 70+  → "A"
// 60+  → "B"
// 50+  → "C"
// 50 এর নিচে → "F"
// 3জন student এর report print করো


//* user এর কাছ থেকে একটা sentence নাও (scan)
// প্রতিটা word কতবার আছে count করো (map use করো)
// output:
// "the" → 2
// "go"  → 1
// "is"  → 1

//* Shape interface বানাও — Area() float64
// Circle struct বানাও — Radius
// Rectangle struct বানাও — Width, Height
// Triangle struct বানাও — Base, Height
// তিনটার Area বের করো interface দিয়ে
// সবচেয়ে বড় area কোনটা সেটা print করো

//* Product struct — Name, Price, Quantity
// slice তে 5টা product রাখো
// range দিয়ে loop করো
// Total inventory value বের করো (price * quantity)
// সবচেয়ে দামি product টা print করো
// pointer দিয়ে একটা product এর price update করো

//* map[string]string দিয়ে contact book বানাও
// key = নাম, value = phone number
// for loop দিয়ে সব contact print করো
// scan দিয়ে user এর কাছ থেকে নাম নাও
// সেই নাম contact book এ আছে কিনা check করো
// থাকলে number দেখাও, না থাকলে "Not found"
*/

type Student struct {
	name  string
	grade float32
	marks []int
}

func main() {

	// * Report card
	// students := []Student{
	// 	{name: "Elma", grade: 4.32, marks: []int{67, 47, 89, 45, 65}},
	// 	{name: "Fariha", grade: 4.62, marks: []int{67, 56, 34, 47, 78}},
	// 	{name: "Faija", grade: 4.67, marks: []int{80, 49, 82, 44, 65}},
	// 	{name: "Hasina", grade: 4.82, marks: []int{47, 78, 86, 55, 95}},
	// 	{name: "Samsul", grade: 4.92, marks: []int{90, 68, 84, 75, 85}},
	// }

	// var studentsAvgMarks []int
	// for _, student := range students {
	// 	var studentTotalMarks int
	// 	for _, value := range student.marks {
	// 		studentTotalMarks += value
	// 	}
	// 	studentsAvgMarks = append(studentsAvgMarks, studentTotalMarks/len(student.marks))
	// }

	// var studentGrade []string
	// for _, value := range studentsAvgMarks {
	// 	if value > 79 && value < 101 {
	// 		studentGrade = append(studentGrade, "A+")
	// 	} else if value > 69 {
	// 		studentGrade = append(studentGrade, "A")
	// 	} else if value > 59 {
	// 		studentGrade = append(studentGrade, "B")
	// 	} else if value > 49 {
	// 		studentGrade = append(studentGrade, "C")
	// 	} else {
	// 		studentGrade = append(studentGrade, "F")
	// 	}
	// }

	// fmt.Println(studentsAvgMarks)
	// fmt.Println(studentGrade)

	//* Word counter

}
