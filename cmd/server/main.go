package main

import (
	"github.com/muhammad21236/Go-gRPC-Service/internal/db"
	"github.com/muhammad21236/Go-gRPC-Service/internal/rocket"
)

func Run() error {
	// Responsible for initializing the gRPC server and starting it.
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		panic(err)
	}

}
