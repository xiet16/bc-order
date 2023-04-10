package main

import (
	"bc-order/handler"
	pb "bc-order/proto"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"log"

	"github.com/gin-gonic/gin"
	httpServer "github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	serviceName   = "bc-order"
	version       = "latest"
	serverAddress = ":9085"
)

func main() {
	srv := httpServer.NewServer(
		server.Name(serviceName),
		server.Address(serverAddress),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// register router
	demo := handler.NewDemo()
	demo.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()),
	)
	service.Init()
	service.Run()
}

func main1() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterBcOrderHandler(srv.Server(), new(handler.BcOrder)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
