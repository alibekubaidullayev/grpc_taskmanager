package main

import (
	"log"
	"net"
	"tm/src/core"
	"tm/src/pb"
	"tm/src/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := core.GetDB()
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err.Error())
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to create tcp server")
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterTaskManagerServer(s, &server.Server{Db: db})
	if err := s.Serve(l); err != nil {
		log.Fatalln("failed to server:", err)
	}
}
