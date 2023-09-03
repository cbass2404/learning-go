package main

import "log"

// Loops and ranging over data

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	// _ is equal to a blank identifier when index is not needed
	for _, animal := range animals {
		log.Println(animal)
	}

	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"

	for animalType, animal := range animalsMap {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"

	// strings are a slice of runes in go, without further manipulation it will not show the actual letters
	// strings in go are immutable so reassigning to a new string value destroys the old one and adds the new one to the var
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "John@Smith.com", 30})
	users = append(users, User{"Mary", "Jones", "Mary@Jones.com", 20})
	users = append(users, User{"Sally", "Brown", "Sally@Brown.com", 45})
	users = append(users, User{"Alex", "Anderson", "Alex@Anderson.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
