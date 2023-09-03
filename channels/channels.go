package main

import (
	"log"

	"github.com/cbass2404/my-package/helpers"
)

/*
* Channels
 */

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	// pass the randomNumber to the channel
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	// defer waits for the current function to finish and then executes the command
	// always close channels after use due to a limitation of the amount accessible at once
	defer close(intChan)

	// turn it into a go routine for concurrent running
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
