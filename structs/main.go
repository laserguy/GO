package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo  /* Writing this means 'contactInfo contactInfo'*/
}


func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "a@a.com",
			zipCode: 954410,
		},
	}

	jim.print()
	jim.updateName("Jimmy")  // Golang allows this shortcut for the pointers
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v\n", p)
}