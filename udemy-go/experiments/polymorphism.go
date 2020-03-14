package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) makeSound() {
	fmt.Println("Making Sound")
}

func (a Animal) move() {
	fmt.Println("Moving...")
}

type Dog struct {
	a Animal
}

func main() {

	d := Dog{a: Animal{"dog"}}

	d.a.move()
}
