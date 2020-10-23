package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	//Person Person    instead of say that Android has a person
	Person //we can say that Android is a Person
	Model  string
}

func main() {
	a := new(Android)
	a.Person.Talk() //or only a.Talk()
}
