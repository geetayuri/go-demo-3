package main

import (
	"os"

	"github.com/nguitarpb/7-solutions/configs"
	"github.com/nguitarpb/7-solutions/modules/servers"
	"github.com/joho/godotenv"

	"sync"

)

func startFiberServer() {
	if err := godotenv.Load("./.env"); err != nil {
		panic(err.Error())
	}

	cfg := new(configs.Configs)
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	s := servers.NewServer(cfg)
	s.Start()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go startFiberServer()

	grpcServer := servers.NewGrpcServer()
	go grpcServer.StartGRPCServer(&wg)

	wg.Wait()
}
