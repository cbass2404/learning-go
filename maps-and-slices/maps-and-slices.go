package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	/*
	* MAP *
	* Maps are generated using the built-in make func.
	* It follows this pattern make(map[storageKey]valueType)
	* Can use interface{} as the valueType but is not recommended due to casting requirements.
	 */
	// myMap := make(map[string]User)

	// me := User{
	// 	FirstName: "Cory",
	// 	LastName:  "Bass",
	// }
	// myMap["me"] = me

	// log.Println(myMap["me"].FirstName)

	// var myNewVar float32
	// myNewVar = 11.1

	/*
	* SLICE *
	* Slices are generated using var sliceName []type
	 */
	var mySlice []string

	mySlice = append(mySlice, "Cory")
	mySlice = append(mySlice, "Kristen")
	mySlice = append(mySlice, "Chloe")
	mySlice = append(mySlice, "Gus")

	sort.Strings(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers[0:2])

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)
}
