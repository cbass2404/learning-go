package typesStructs

/*
* Private/Public access is determined by the capitalization of the first letter of the var
* user *= private access
* User *= public access to other modules
 */

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}


func main() {
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, "PhoneNumber:", user.PhoneNumber, "Age:", user.Age, "Birthdate:", user.BirthDate)
}