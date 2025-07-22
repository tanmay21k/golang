package main

import "fmt"

type User struct {
	name       string
	age        int
	isVerified bool
	email      string
}

func (u User) printer() { // created an alias for the struct
	fmt.Printf("The person verfication return -> %v", u.isVerified)
}

func main() {
	Tanmay := User{"Tanmay", 11, true, "example@gmail.com"}
	fmt.Println(Tanmay)
	fmt.Printf("The detailed value are: %+v \n", Tanmay)
	Person1 := User{"Person1", 91, false, "example121@gmail.com"}
	fmt.Printf("The age of person1 is %v and the email is %v\n ", Person1.age, Person1.email)

	fmt.Printf("\n Using method \n")
	Tanmay.printer()
}
