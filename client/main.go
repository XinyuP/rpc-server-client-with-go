package main

import (
	"log"
	"fmt"
	"net/rpc"
)

type Item struct {
	Title string
	Body string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040") 
	// establish a connection with an RPC server on the local machine at 
	// port 4040, and if it's successful, it provides a client for 
	// communicating with that server. If it encounters any error during 
	// this process, it returns that error.

	if err != nil {
		log.Fatal("connection error", err)
	}

	a := Item{"first", "a test item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("database: ", db)

	client.Call("API.EditItem", Item{"second", "a new second item"}, &reply)
	client.Call("API.GetDB", "", &db)	
	fmt.Println("database: ", db)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("database: ", db)

	client.Call("API.GetByName", "first", &reply)
	fmt.Println("first item: ", reply)

}