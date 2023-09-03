package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
* Reading and Writing JSON
 */

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJSON := `[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	/*
	* Reading JSON to a struct
	 */
	// called unmarshalled because the package used in goland is called marshalled/unmarshalled
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJSON), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	/*
	* Writing JSON from a struct
	 */

	var mySlice []Person
	m1 := Person{FirstName: "Hal", LastName: "Jordan", HairColor: "black", HasDog: false}
	m2 := Person{FirstName: "Wally", LastName: "West", HairColor: "red", HasDog: true}
	mySlice = append(mySlice, m1, m2)

	// MarshalIndent formats the JSON during translation, with server use just use Marshal
	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
