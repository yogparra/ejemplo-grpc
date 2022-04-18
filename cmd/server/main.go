package main

import (
	"log"

	"github.com/yogparra/ejemplo-grpc/internal/db"
	"github.com/yogparra/ejemplo-grpc/internal/juego"
	"github.com/yogparra/ejemplo-grpc/internal/transport/grpc"
)

func Run() error {

	log.Println("Starting Up Juego gRPC Service")

	jgStore, err := db.New()
	if err != nil {
		return err
	}

	jgService := juego.New(jgStore)

	jgHandler := grpc.New(jgService)

	if err := jgHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
