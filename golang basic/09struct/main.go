package main

import "fmt"

type User struct {
	name       string
	age        int
	isVerified bool
	email      string
}

func main() {
	Tanmay := User{"Tanmay", 11, true, "example@gmail.com"}
	fmt.Println(Tanmay)
	fmt.Printf("The detailed value are: %+v \n", Tanmay)
	Person1 := User{"Person1", 91, false, "example121@gmail.com"}
	fmt.Printf("The age of person1 is %v and the email is %v\n ", Person1.age, Person1.email)
}
