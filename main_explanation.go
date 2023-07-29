package main

import "fmt"

type Item struct {
	title string
	body string
}
// The Item type is a simple data structure with a title and body

type API int 
// used to elevate all our functions to methods

// API is a type (an integer), but in this context it doesn't represent 
// an actual integer value. It's used as a receiver to attach methods to 
// (like GetByName, AddItem, etc.), essentially creating a group of methods
// that can be registered as a service with the rpc.Register function.


var database []Item
// database is a slice of Item type, acting as the "database" storing items.


    // receiver for the API pointer
func (a *API) GetByName(title string, reply *Item) error {
	                    // caller     // results 
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}

	*reply = changed
	return nil
}


func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:] ...)
			del = item
			break
		}
	}

	*reply = del
	return nil
}

func main() {
	/* this code is about setting up an RPC server. 
	It creates an instance of a type, tries to register that instance for RPC 
	so its methods can be called remotely, and checks for any errors in that process. 
	If everything is successful, the methods of API can now be called from other machines.
	*/

	// creating an RPC server with Go's built-in net/rpc package. RPC stands for Remote Procedure Call, which is a way to execute code on a remote server.
	var api = new(API)
	// create a server to conncet the methods we wrote 
	// creates a new instance of the API type and assigns it to the variable api. 
	// API is a custom type and it's been defined elsewhere in the code. 
	// It will contain methods that we want to expose as RPC functions.

	err := rpc.Register(api)
	// register the type so we can call its methods remotely
	// The rpc.Register function registers api as a service for RPC, 
	// which means that all the methods of api that fit the RPC format 
	// can now be called remotely. 

	// rpc.Register will return an error if the registration failed 
	// (for example, if api doesn't have any methods that can be used for RPC).

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()
	/*The rpc.HandleHTTP function in Go is used to register an HTTP handler 
	for RPC messages on http.DefaultServeMux. This means it sets up the 
	default HTTP server to understand and route incoming HTTP requests with
	RPC-encoded messages.

	It specifically registers the endpoint /debug/rpc which allows users to 
	view a list of the registered services, methods, and some statistics 
	about the server. It also registers the endpoint / to handle standard 
	RPC messages.

	This is useful because it allows your server to listen and respond to 
	incoming RPC calls over HTTP. This allows clients to use regular HTTP 
	protocols to make RPC calls, instead of having to use a specific RPC 
	protocol.*/


	listener, err := net.Listen("tcp", ":4040")
	// open the connection 

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving", err)
	}
	/*The http.Serve function in Go's net/http package is used to start an 
	HTTP server with a given listener and handler.
	listener: This is a net.Listener object which listens for incoming connections on a network address that you've specified. For example, this could be localhost:8080.

	nil: This is the HTTP handler. If you pass nil, the server uses the 
	default ServeMux in Go's http package. ServeMux is an HTTP request 
	multiplexer (or router) that matches the URL of incoming requests 
	against a list of registered patterns and calls the associated handler 
	for the pattern that matches the URL the most.

	err = http.Serve(listener, nil): This starts the HTTP server and it 
	will block and keep running until it's shut down or it encounters an 
	unrecoverable error. If an error occurs when trying to start the server 
	(for example, if the specified network address isn't available), 
	http.Serve will return an error and this is captured in err.





	*/





}