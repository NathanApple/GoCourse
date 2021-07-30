package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// & : Get memory address from value
	// * : Get value from memory address

	// Structs is value type
	// While things like slice is references type

	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()


	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

}

func (p *person) updateName(newFirstName string)  {
	//(*p).firstName = newFirstName
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}