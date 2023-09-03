package main

import (
	"log"

	"github.com/cbass2404/my-package/helpers"
)

/*
* Go Packages/Libraries
 */

/*
* cmd line $ go mod init github.com/user/package-name
 */

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNumber = 123

	log.Println(myVar)
}
