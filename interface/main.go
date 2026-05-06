package main

import "fmt"

// ts ==> interface for data shape
// go ==> interface for behavior contract

type Animal interface {
	speek()
}

type Dog struct{}
type Cat struct{}
type Human struct {
	name string
}

func (d Dog) speek() {
	fmt.Println("Bark")
}
func (c Cat) speek() {
	fmt.Println("Meow")
}

func (h Human) speek() {
	fmt.Println("My name is", h.name)
}

func makeSound(a Animal) {
	a.speek()
}

// abstraction, polymorphism, inheritance and encapsulation

func main() {
	dexter := Dog{}
	tommy := Cat{}
	human := Human{
		name: "Fajla",
	}
	makeSound(dexter)
	makeSound(tommy)
	makeSound(human)
}
