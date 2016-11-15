package main

import (
	//"fmt"
	"github.com/flosch/pongo2"
	"net/http"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

var tplExample *pongo2.Template
var user User
var names []string

func init() {
	tplExample = pongo2.Must(pongo2.FromFile("example.html"))
	names = []string{}
	names = append(names, "1111", "2222")
}
func examplePage(w http.ResponseWriter, r *http.Request) {
	// Execute the template per HTTP request
	err := tplExample.ExecuteWriter(pongo2.Context{"items": names}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", examplePage)
	http.ListenAndServe(":8080", nil)
}
