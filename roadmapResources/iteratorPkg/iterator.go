package iteratorpkg

import "fmt"

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func RegularGetUsers() []user{
	var users []user
	for i := 0; i < 100; i++ {
		users = append(users, user{
			ID:   i,
			Name: fmt.Sprintf("User%v", i),
		})
	}
	return users
}

func Printer() {
	for _ , v := range RegularGetUsers() {
		fmt.Println(v)
	}
}