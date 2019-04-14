package main

import (
	"github.com/fpawel/mproducto/internal/auth"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"testing"
)

func TestRpcClient(t *testing.T) {
	c := jsonrpc2.NewHTTPClient("http://localhost:3001/rpc")
	defer c.Close()

	// Synchronous call using positional params and HTTP.
	var reply string
	err := c.Call("Auth.Login", auth.Credentials{"user1", "password1"}, &reply)
	log.Println(reply, err)

	log.Println(c.Call("Auth.ValidateNewUsername", [1]string{"user1"}, &reply), reply)
	log.Println(c.Call("Auth.ValidateNewUsername", [1]string{"user3"}, &reply), reply)
}
