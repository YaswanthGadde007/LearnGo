package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	yaswanth := person{
		firstName: "yaswanth",
		lastName:  "gadde",
		contact: contactInfo{
			email:   "test@gmail.com",
			zipCode: 530044,
		},
	}
	yaswanthPointer := &yaswanth
	yaswanth.print()
	yaswanthPointer.updateFirstName("yashwanth")

	yaswanth.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
