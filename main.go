/**
*  1) A Basic Web server
*  2) Adding a Router ( using Mux )
*		i ) use go get to import external package
*		ii) we create a basic router, adds the route "/"
*           and assign the index handler to run when the endpoint is called
* 3)  Adding additional basic router
* 4)  A basic Todo model
*		i) In Golang, a Struct will typically serve as model.

 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo // Todos is a slice of type Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// Send back some JSOn
	// Simulate a real response and mock out the TodoIndex with static
	// data
	todos := Todos{
		Todo{Name: "Write presenation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
