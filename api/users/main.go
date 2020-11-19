package main

import (
	"log"

	"github.com/frouioui/tagenal/api/users/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	usersrv, err := server.NewUserServerAPI()
	if err != nil {
		log.Fatal(err.Error())
	}

	waiter := make(chan error)

	go func() {
		err := usersrv.RunServerHTTP()
		waiter <- err
	}()

	go func() {
		err := usersrv.RunServerGRPC()
		waiter <- err
	}()

	err = <-waiter
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("server shutdown")
}
