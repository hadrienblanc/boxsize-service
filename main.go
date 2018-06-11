package main

import (
	"fmt"
	"log"
	"net/http"
  "flag"
)

const (
	DefaultPort = 3000
)

func main() {
	var port = flag.Int("port", DefaultPort, "Specify the port number")
	flag.Parse()

	server := http.NewServeMux()
	server.HandleFunc("/", BoxSizeHandler)

	fmt.Printf("BoxSize Server will run on port %d\n", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), server)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
