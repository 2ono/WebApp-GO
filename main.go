package main

import "log"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "Test@gmail.com", 23})
	users = append(users, User{"mary", "Kim", "Test1@gmail.com", 20})
	users = append(users, User{"King", "Harry", "Test2@gmail.com", 10})
	users = append(users, User{"Quen", "Brow", "Test3@gmail.com", 50})
	users = append(users, User{"Roy", "Red", "Test4@gmail.com", 42})
	users = append(users, User{"Alex", "Jun", "Test5@gmail.com", 70})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
