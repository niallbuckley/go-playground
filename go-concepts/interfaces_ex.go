package main

import (
	"fmt"
)


type Person interface {
     getName() string
     isMale()  bool
}

type Customer struct {
	name     string
        male     bool
	isMember bool
}

type Employee struct {
	name        string
        male        bool
	employeeId  int
}

func (c Customer) getName() string {
	return c.name
}

func (c Customer) isMale() bool {
	return c.male
}

func (e Employee) isMale() bool {
	return e.male
}

func (e Employee) getName() string {
	return e.name
}

func identifyPerson(p Person){
	fmt.Printf("Name: %s, Male: %v\n", p.getName(), p.isMale())

	// Use type assertion to differentiate behavior between Employee and Customer
	if _, ok := p.(Employee); ok {
		fmt.Printf("%s is an Employee \n", p.getName())
	} else if _, ok := p.(Customer); ok {
		fmt.Printf("%s is a Customer \n", p.getName())
	} else {
		fmt.Println("Unknown Person Type")
	}
}

func main() {
	mikeC := Customer{name: "Mike", male: true, isMember: true}
	hanE := Employee{name: "Hannah", male: false, employeeId: 1234}

        identifyPerson(mikeC)
        identifyPerson(hanE)
}

