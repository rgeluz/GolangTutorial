package main

import "fmt"

type contactInfo struct { //to be used to embedd struct in struct, see person struct
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //field name is optional when using embedded struct, using just contactInfo == concact contactInfo also works, and then in declaration use contactInfo : contactInfo{ ....}
}

func main() {

	//method 1 for declaring and initializing a struct
	//alex := person{"Alex," "Anderson"}  //strict on order defined inside struct definition

	//method 2 for declaring and initializing a struct
	//alex := person{firstName: "Alex", lastName: "Anderson"}  //more full proof way, incase ordering in the definition is changed
	//fmt.Println(alex)

	//method 3 for declaring and initializing a struct
	/*var alex person
	fmt.Println(alex)       //output: { }
	fmt.Printf("%+v", alex) //print all the different field names and values. output: {firstName: lastName:}%

	fmt.Println("")
	*/
	//updated struct values
	/*alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)       //output: {Alex Anderson}
	fmt.Printf("%+v", alex) //output: {firstName:Alex lastName:Anderson}%
	*/

	//struct with embedded struct
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//fmt.Printf("%+v", jim) //ouput: {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:94000}}%

	jim.print() //alternative to above, uses a receiver function

	fmt.Println("")

	//update first name
	jim.updateName("jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

//Structs with Receiver Functions
func (p person) print() {
	fmt.Printf("%+v", p) //output: {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:94000}}%
}
