package main

import (
	"fmt"
	"github.com/amitkumarbhagat/golang/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

//main is application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
