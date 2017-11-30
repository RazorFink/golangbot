package main

import (
	"fmt"
)

//Employee is an employee
type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	fmt.Println("Employee 1", emp1)

	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 2", emp2)

	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Adreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	var emp4 Employee
	fmt.Println("Employee 4", emp4)

	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)

	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)

	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	emp7.age = 8
	fmt.Println("Employee 7:", emp7)

	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Last Name:", (*emp8).lastName)
	fmt.Println("Age:", emp8.age)
	fmt.Println("Salary:", emp8.age)

	type Person struct {
		string
		int
	}
	p := Person{"Naveen", 50}
	fmt.Println(p)

	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	type Address struct {
		city, state string
	}

	type RealPerson struct {
		name    string
		age     int
		address Address
	}

	var p3 RealPerson
	p3.name = "Naveen"
	p3.age = 50
	p3.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.address.city)
	fmt.Println("State:", p3.address.state)

	// Anonymous fields that are structs have their fields promoted to fields
	// in the parent struct
	type PromotedPerson struct {
		name    string
		age     int
		Address // Anonymous struct
	}

	var p4 PromotedPerson
	p4.name = "Naveen"
	p4.age = 50
	p4.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p4.name)
	fmt.Println("Age:", p4.age)
	fmt.Println("City:", p4.city)   //Promoted Field
	fmt.Println("State:", p4.state) //Promoted Field
}
