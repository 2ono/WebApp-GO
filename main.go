package main

import "log"

func main() {

	myVar := "tiger"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	case "tiger":
		log.Println("cat is set to tiger")
	default:
		log.Println("everything is here")

	}

}
