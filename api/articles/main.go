package main

import (
	"log"

	"github.com/frouioui/tagenal/api/articles/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	artsrv, err := server.NewArticleServerAPI()
	if err != nil {
		log.Fatal(err.Error())
	}

	waiter := make(chan error)

	go func() {
		err := artsrv.RunServerHTTP()
		waiter <- err
	}()

	go func() {
		err := artsrv.RunServerGRPC()
		waiter <- err
	}()

	server.SetReady(true)
	err = <-waiter
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("server shutdown")
}
