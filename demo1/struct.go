package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	person := Person{name: name}
	person.age = 24
	return &person
}

func (this Person) getPerson() string {
	return this.name
}

func main() {
	s := Person{name: "tony", age: 50}
	fmt.Println(s.name)
	fmt.Println(newPerson("pug"))
}
