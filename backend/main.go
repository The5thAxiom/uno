package main

import (
	"backend/server"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	fmt.Println("UNO by Samridh and Sanjeev")

	server := server.Server {
		Port: port,
		Mux: http.NewServeMux(),
	}

	fmt.Printf("starting uno server at port %d", server.Port)

	err := server.Run()
	log.Fatal(err)
}