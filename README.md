# ToDoAppApi
Making a RESTful JSON API in Go

 1) A Basic Web server

 2) Adding a Router ( using Mux )

		i ) use go get to import external package

		ii) we create a basic router, adds the route "/"
           and assign the index handler to run when the endpoint is called

 3)  Adding additional basic router

 4)  A basic Todo model

		i) In Golang, a Struct will typically serve as model.

 5)  Creating idiomatic JSON

		i) By using the struct tags we can control exactly how our
		   struct will be marshalled to JSON

 6)  Code refactoring

 		i) now we utilizes a struct to contain more detailed information
 		   about the route. Specifically, we can specify the action, such as
 		   GET, POST, DELETE, etc

7)   Creating a Web log Pakcage

8)  Applying the logger decorator
	
		i) To apply the decorator, when we create the router, we will simply wrap 
			all our current routes in it by updating our NewRouter function. 

9)  Refactoring Routes file					

