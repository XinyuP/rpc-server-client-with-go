package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body string
}

type API int // used to elevate all our functions to methods

var database []Item

func (a *API) GetDB(emprt string, reply *[]Item) error {
	*reply = database
	return nil
}
    // receiver for the API pointer
func (a *API) GetByName(title string, reply *Item) error {
	                    // caller     // results 
	var getItem Item

	for _, val := range database {
		if val.Title == title {
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
		if val.Title == edit.Title {
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
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:] ...)
			del = item
			break
		}
	}

	*reply = del
	return nil
}

func main() {
	var api = new(API)
	// create a server to conncet the methods we wrote 
	err := rpc.Register(api)
	// register the type so we can call its methods remotely

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()
	// register an HTTP handler 

	listener, err := net.Listen("tcp", ":4040")
	// open the connection 

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)	
	// Start an HTTP server with the specified listener and the default 
	// request multiplexer, and if there's an error, assign it to err

	if err != nil {
		log.Fatal("error serving", err)
	}






	// fmt.Println("initial database: ", database)
	// a := Item{"first", "a test item"}
	// b := Item{"second", "a second item"}
	// c := Item{"third", "a third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)

	// fmt.Println("second database: ", database)

	// DeleteItem(b)
	// fmt.Println("third database: ", database)

	// EditItem("third", Item{"fourth", "a new item"})
	// fmt.Println("fourth database: ", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)
}