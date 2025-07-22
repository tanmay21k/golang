package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("welcome to json video")
	EncodingJson()
}

type courses struct {
	Name     string `json:"Course Name"` // created alias for the json key
	Price    int
	Platform string
	Password string   `json:"-"`              // to not display this field
	Tags     []string `json:"Tags,omitempty"` // omit if empty. Always check the syntax
}

func EncodingJson() {
	lcoCourses := []courses{
		{"reactjs",
			13000,
			"Udemy",
			"Password@123",
			[]string{"frontend", "UI"},
		},
		{"MERN",
			19000,
			"Udemy",
			"Password@1223",
			[]string{"frontend", "UI", "frontend"},
		},
		{"Golang",
			2000,
			"Linkedin",
			"Password@4123",
			nil,
		},
	}

	// finalJson, err := json.Marshal(lcoCourses) // conversion to bytes
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // indenting the text as well (structName, prefix, postfix)

	if err != nil {
		panic(err)
	}

	fmt.Println(finalJson)

	fmt.Printf("%s\n", finalJson) // Converting to string
}
