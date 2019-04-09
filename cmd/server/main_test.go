package main

import (
	"github.com/fpawel/mproducto/internal/api"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"testing"
)

func TestRpcClient(t *testing.T) {
	c := jsonrpc2.NewHTTPClient("http://localhost:3000/rpc")
	defer c.Close()

	// Synchronous call using positional params and HTTP.
	var reply string
	err := c.Call("Auth.Login", api.UserPass{"a", "b"}, &reply)
	log.Println("Client Auth.Login:", reply, err)
}
