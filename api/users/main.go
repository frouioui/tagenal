package main

import (
	"log"

	"github.com/frouioui/tagenal/api/users/server"
)

// func (s *server) Compute(cxt context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
// 	result := &pb.UserResponse{}
// 	result.Result = r.A + r.B

// 	logMessage := fmt.Sprintf("A: %d   B: %d     sum: %d", r.A, r.B, result.Result)
// 	log.Println(logMessage)

// 	return result, nil
// }

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	usersrv, err := server.NewUserServerAPI()
	if err != nil {
		log.Fatal(err.Error())
	}

	waiter := make(chan error)

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
