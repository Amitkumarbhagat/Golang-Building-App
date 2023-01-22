package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//	check to see if we already have the template in our cache
	_, inMap := tc[t]

	if !inMap {
		//	need to create the template
	} else {
		//	we have template in the cache
	}
}