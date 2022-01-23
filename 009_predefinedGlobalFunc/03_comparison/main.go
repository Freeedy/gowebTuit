package main

import (
	"log"
	"os"
	"text/template"
)

//https://egovaze.udemy.com/course/go-programming-language/learn/lecture/5995886#overview
//https://github.com/Freeedy/golang-web-dev
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
