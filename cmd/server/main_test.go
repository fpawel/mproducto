package main

import (
	"github.com/fpawel/mproducto/internal/api"
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

	var i api.UserInfo
	err = c.Call("Auth.UserInfo", [1]string{reply}, &i)
	log.Println(i, err)

	i = api.UserInfo{}
	err = c.Call("Auth.UserInfo", [1]string{"??"}, &i)
	log.Println(i, err)
}
