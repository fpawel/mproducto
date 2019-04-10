package main

import (
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"testing"
)

func TestRpcClient(t *testing.T) {
	c := jsonrpc2.NewHTTPClient("http://localhost:3000/rpc")
	defer c.Close()

	// Synchronous call using positional params and HTTP.
	var reply string
	//err := c.Call("Auth.Login", api.Credentials{"user1", "password1"}, &reply)
	//log.Println(reply, err)

	err := c.Call("Greet.Hello", [1]string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTU0ODgyMzkzfQ.50AIhOHJG7Fh8DIg94jozPn9GhlO-OKlgsjjSo4nu_4"}, &reply)
	log.Println(reply, err)
}
