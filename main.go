package main

import "fmt"

type person struct{
	firstName string
	lastName string
	contact contactInfo
	job
}

type job struct{
	name string
	desc string
}

type contactInfo struct{
	email string
	zipCode int
}

func main() {
	// alex:=person{"Alex", "Anderson"}   // method 1 to declare
	// alex:=person{
	// 	firstName:"Alex",
	// 	lastName: "Blue",
	// }									// method 2 to declare
	// fmt.Println(alex);

	// var alex person;
	// alex.firstName = "Alex"
	// alex.lastName = "Blue"
	// fmt.Printf("%+v", alex);                // method 3 , notice the specifier %+v

	alex:=person{
		firstName: "Alex",
		lastName: "Blue",
		contact: contactInfo{
			email: "a@mail.com",
			zipCode: 110001,
		},
		job: job{
			name: "Developer",
			desc: "create apps",
		},
	}

	alex.print();
	// alex.updateName("John")
	// alex.print()
}

func (p person) print(){
	fmt.Printf("%+v", p);
}

// func (p person) updateName(newFirstName string){
// 	p.firstName = newFirstName
// }