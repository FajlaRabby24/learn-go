// enums
package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type OfficeStatus string

const (
	StatusOpen    OfficeStatus = "open"
	StatusHalf    OfficeStatus = "half_day"
	StatusClose   OfficeStatus = "closed"
	StatusInvalid OfficeStatus = "invalid"
)

func getWorkDayStatus(day Weekday) OfficeStatus {
	switch day {
	case Sunday, Monday, Tuesday, Wednesday:
		return StatusOpen
	case Thursday:
		return StatusHalf
	case Friday, Saturday:
		return StatusClose
	default:
		return StatusInvalid
	}
}

func main() {
	fmt.Println(getWorkDayStatus(Sunday))
	fmt.Println(getWorkDayStatus(Friday))
}
