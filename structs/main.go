package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(name string) {
	p.Name = name
}

func main() {
	myPerson := NewPerson("hurtsec", 33)
	myPerson.changeName("J")

	a := 7
	b := &a
	*b = 9
	fmt.Println(*b)
	fmt.Println(a)

	mySlice := []int{
		1, 2, 3,
	}

	for index := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)
}
