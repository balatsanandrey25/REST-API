package main

import (
	"log"

	server "github.com/balatsanandrey25/REST-API"
	"github.com/balatsanandrey25/REST-API/pkg/handler"
)

const (
	port = "8080"
)

func main() {
	handler := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run(port, handler); err != nil {
		log.Fatalf("error occured while runnig http server %s", err.Error())
	}
}
