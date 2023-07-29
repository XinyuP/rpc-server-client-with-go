# rpc-server-client-with-go
Building a Basic RPC Server and Client with Go


RPC remote procedure call

We want to make it so that external computer can call to our program interface and execute various pieces of our code.

Create a basic CRUD application 


In Golang, the net/rpc package is often used to implement RPC (Remote Procedure Call). When using this package, there are certain format requirements for functions that can be exposed as RPC functions:

The function must be a method that is exported from a type. This means that the function must start with an uppercase letter (as identifiers starting with a lowercase letter are not exported in Go).

The function must take two arguments, both of which must be exported (start with uppercase letter) or built-in types.

The first argument represents the arguments provided by the client, and the second argument is a pointer type that the function can write the results to. This means that the client can obtain the result of the function.

The function must return an error.