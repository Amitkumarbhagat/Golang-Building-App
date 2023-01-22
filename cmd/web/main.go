//package main
//
//import (
//	"fmt"
//	"github.com/amitkumarbhagat/golang/pkg/config"
//	"github.com/amitkumarbhagat/golang/pkg/handlers"
//	"github.com/amitkumarbhagat/golang/pkg/render"
//	"net/http"
//)
//
//const portNumber = ":8080"
//
////main is application function
//func main() {
//	var app config.AppConfig
//
//	tc, err := render.CreateTemplateCache()
//	http.HandleFunc("/", handlers.Home)
//	http.HandleFunc("/about", handlers.About)
//
//	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
//	_ = http.ListenAndServe(portNumber, nil)
//}
//
//

package main

import (
	"fmt"
	"github.com/amitkumarbhagat/golang/pkg/config"
	"github.com/amitkumarbhagat/golang/pkg/handlers"
	"github.com/amitkumarbhagat/golang/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
