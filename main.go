package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "this is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is the about page and 2 + 2 is %d ", sum))
	_, _ = fmt.Fprintf(w, "this is the about page")
}

//addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

//main is application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
