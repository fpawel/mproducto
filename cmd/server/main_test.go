package main

import (
	"github.com/fpawel/mproducto/internal/rpcapi"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"testing"
)

func TestNewUser(t *testing.T) {
	var (
		token   string
		profile rpcapi.AuthProfileResult
	)

	newUser := rpcapi.AuthRegisterArg{Name: "newuser112", Email: "binf1611@.mailnnu.ru", Pass: "wtf111222333444"}

	if err := cli.Call("Auth.Register", newUser, &token); err != nil {
		t.Error(err)
	}
	log.Println("new user token:", token)

	if err := cli.Call("Auth.Profile", [1]string{token}, &profile); err != nil {
		t.Error(err)
	}
	log.Println("new user profile:", profile)

	if err := cli.Call("Auth.Unregister", [1]string{token}, new(struct{})); err != nil {
		t.Error("new user delete: ", err)
	} else {
		log.Println("new user delete - OK ")
	}

}

func TestAuthLoginAndProfile(t *testing.T) {
	var token string
	err := cli.Call("Auth.Login", rpcapi.AuthLoginArg{"pawel1", "11111111"}, &token)
	if err != nil {
		t.Error(err)
	}
	log.Println("TOKEN:", token)

	var profile rpcapi.AuthProfileResult

	if err = cli.Call("Auth.Profile", [1]string{token}, &profile); err != nil {
		t.Error(err)
	}
	log.Println("PROFILE:", profile)
}

var cli = jsonrpc2.NewHTTPClient("http://localhost:3001/rpcapi")
