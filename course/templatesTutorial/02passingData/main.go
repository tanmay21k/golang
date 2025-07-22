package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dataFile.gohtml"))
}

// {{}} contains . where the data is replaced
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "dataFile.gohtml", "tanmay") // need to use "" for the string
	if err != nil {
		panic(err)
	}
}
