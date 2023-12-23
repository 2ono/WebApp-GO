package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println(myString)

	changeUsingPointer(&myString)

	log.Println(myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue

}
