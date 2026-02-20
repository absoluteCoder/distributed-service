package main

import (
	"log"

	"github.com/absoluteCoder/distributedservice/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
