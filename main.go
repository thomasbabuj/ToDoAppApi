/**
*  1) A Basic Web server
*  2) Adding a Router ( using Mux )
*		i ) use go get to import external package
*		ii) we create a basic router, adds the route "/"
*           and assign the index handler to run when the endpoint is called
*
 */

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
