package main

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/h-mall/hm-gateway/proxy"
)

func main() {
	h := server.New(server.WithHostPorts("0.0.0.0:8080"))

	// h.Use(middleware.AuthMiddleware())

	h.GET("/api/user/:id", proxy.GetUserByID)

	log.Println("API Gateway is running on :8080")

	if err := h.Run(); err != nil {
		log.Fatalf("gateway stopped with error: %v", err)
	}
}