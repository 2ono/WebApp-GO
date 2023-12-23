package main

import "fmt"

func main() {
	fmt.Println("Hello World.")

	var whatToSay string = "Good Byecruel world"

	var i int
	i = 29

	i = 30
	whatToSay = "Good"
	fmt.Println(whatToSay)
	fmt.Println(i)

	i = 8

	fmt.Println("i is set to ", i)

	wwhatWasSaid, theotehrThingThatwasSaid := saySomething()

	fmt.Println("The Function returned", wwhatWasSaid, theotehrThingThatwasSaid)

}

func saySomething() (string, string) {
	return "something", "else"
}
