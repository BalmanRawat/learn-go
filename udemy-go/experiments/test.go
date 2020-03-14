package main

import "fmt"

type person struct {
	name string
}

type car struct {
	name string
}

func (p *person) update(n string) {
	p.name = n
}

func main() {

	var owner map[string]string
	p1 := person{name: "Balman Rawat"}
	c1 := car{
		name: "lamborgini",
	}
	fmt.Printf("Name: %s Car: %s\n", p1.name, c1.name)

	fmt.Printf("Owners: %+T", owner)
}
