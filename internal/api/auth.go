package api

import (
	"github.com/powerman/rpc-codec/jsonrpc2"
	"io/ioutil"
	"log"
	"net"
)

type Auth struct {

}


type UserPass struct {
	User, Pass string
}

type UserPassArgContext struct {
	UserPass
	jsonrpc2.Ctx
}


func (x *Auth) Login(p UserPassArgContext, token *string) error {
	*token = "login user"

	r := jsonrpc2.HTTPRequestFromContext(p.Context())
	host,port,err :=  net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("From Auth.Login: host=%q port=%q\n", host, port)

	b,err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("body=%q", string(b))


	return nil
}

func (x *Auth) New(p UserPassArgContext, token *string) error {
	*token = "new user"

	r := jsonrpc2.HTTPRequestFromContext(p.Context())
	host,port,err :=  net.SplitHostPort(r.RemoteAddr)
	log.Printf("From Auth.New: host=%q port=%q err=%v\n", host, port, err)

	b,err := ioutil.ReadAll(r.Body)
	log.Printf("body=%q\nerr=%v", string(b), err)

	return nil
}