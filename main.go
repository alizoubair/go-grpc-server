package main

import (
	"fmt"
	"net"

	"github.com/alizoubair/go-grpc-server/api/user"
	userGrpc "github.com/alizoubair/go-grpc-server/api/user/grpc"
	conf "github.com/alizoubair/go-grpc-server/config"
	"google.golang.org/grpc"
)

func main() {
	dbR, dbW, err := conf.InitDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbR.Close()
		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		err := dbW.Close()
		if err != nil {
			panic(err)
		}
	}()

	log := conf.InitLog()

	userRepo := user.NewUserRepository(dbR, dbW)
	userSvc := user.NewUserService(log, userRepo)

	list, err := net.Listen("tcp", ":"+conf.Configuration.Port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	userGrpc.NewUserServerGrpc(server, userSvc)

	fmt.Printf("Listening gRPC server on: %s\n", conf.Configuration.Port)
	err = server.Serve(list)
	if err != nil {
		panic(err)
	}
}
