package main

import "log"

type myStruct struct {
	FirstName string
}

/*
* Functions using a receiver as below can access anything inside the
* assocaited struct. Think of it like a class method assocaited to all
* created instances of this struct.
 */
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Cory"
	myVar2 := myStruct{
		FirstName: "Kristen",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar is set to", myVar2.printFirstName())
}
