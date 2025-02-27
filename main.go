package main

import (
	"fmt"
	"log"
	"m1thrandir225/cicd2025/api"
)

func main() {
	fmt.Println("Hello World")
}

func runGinServer() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("Cannot create server")
	}
	err = server.Start("http://localhost:8080")
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
