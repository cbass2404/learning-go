package main

import "log"

// Decision Structures

func main() {
	var isTrue bool
	isTrue = false
	myNum := 100
	cat := "cat"

	/*
	* IF ELSE
	 */
	if isTrue {
		log.Println("isTrue is: ", isTrue)
	} else {
		log.Println("!isTrue is: ", isTrue)
	}

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else if myNum == 101 || isTrue {
		log.Println("2")
	} else if myNum > 1000 && !isTrue {
		log.Println("3")
	}

	/*
	* SWITCH/CASE
	 */
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	case "fish":
		log.Println("cat is set to fise")

	default:
		log.Println("cat is something else")
	}
}
