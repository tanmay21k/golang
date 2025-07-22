// parsing an reading the data from the file via executeTemplate

package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	// on parsing we get a pointer to the template
	//tpl, err := template.ParseFiles("index.gohtml") // we can send multiple paths as parameter

	// tpl, err := template.ParseGlob("template/*")
	// if err != nil {
	// 	panic(err)
	// }

	// .Execute expects writer and data
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}

	// tpl, err = template.ParseFiles("file.gothml")
	// if err != nil {
	// 	panic(err)
	// }

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}

	// executing specific template
	err = tpl.ExecuteTemplate(os.Stdout, "file.gohtml", nil)
	if err != nil {
		panic(err)
	}
	// executing txt template
	err = tpl.ExecuteTemplate(os.Stdout, "index.txt", nil)
	if err != nil {
		panic(err)
	}
}
