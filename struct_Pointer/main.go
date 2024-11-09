package main

import "fmt"

type person struct {
	name string
	age  int
}

type contact struct {
	email string
	phone int
}

type Employee struct {
	person_d  person
	contact_d contact
}

func main() {

	var p1 person
	p1.name = "virat"
	p1.age = 34

	p2 := person{
		name: "Rohit",
		age:  34,
	}

	var p3 = new(person)

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
	fmt.Println("p3=", *p3)

	var e1 Employee

	e1.person_d = person{
		name: "rahul",
		age:  26,
	}

	e1.contact_d = contact{
		email: "ajhjj",
		phone: 348684,
	}

	fmt.Println("e1=", e1)

}
