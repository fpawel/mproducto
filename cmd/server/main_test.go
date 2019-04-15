package main

import (
	"github.com/fpawel/mproducto/internal/data"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"testing"
)

func TestRpcClient(t *testing.T) {
	c := jsonrpc2.NewHTTPClient("http://localhost:3001/rpc")
	defer c.Close()

	// Synchronous call using positional params and HTTP.
	var reply string
	err := c.Call("Auth.Login", data.Cred{"user1", "password1"}, &reply)
	log.Println(reply, err)

	err = c.Call("Auth.Login", data.Cred{"alexey", "22222222"}, &reply)
	log.Println(reply, err)
}
