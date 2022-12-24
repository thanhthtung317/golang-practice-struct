package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact contactInfo
}

func main() {
	alex := person{
		firstname: "alex",
		lastname: "anderson",
		contact: contactInfo{
			email: "sample@email.com",
			zipCode: 7000000,
		},
	}

	alex.updateFirstname("alexon")

	// fmt.Println(alex)
	alex.print()
}

func (PoiterToPerson *person) updateFirstname(newFirstName string){
	(*PoiterToPerson).firstname = newFirstName
}

func (p person) print(){
	// fmt.Println(p)
	fmt.Printf("%+v", p)
}