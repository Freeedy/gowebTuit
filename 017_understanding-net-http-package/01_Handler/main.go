package main

import (
	"fmt"
	"net/http"
)

//https://egovaze.udemy.com/course/go-programming-language/learn/lecture/6027430#overview
type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
