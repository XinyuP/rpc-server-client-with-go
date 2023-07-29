package main

import "fmt"

type Item struct {
	title string
	body string
}

type API int // used to elevate all our functions to methods

var database []Item

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
	var api = new(API)
	// create a server to conncet the methods we wrote 
	err := rpc.Register(api)
	// register the type so we can call its methods remotely

	if err != nil {
		log.Fatal(error registering API, err)
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