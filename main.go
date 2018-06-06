package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	DefaultPort = 3000
)

func ArgsToPortNumber(args []string) int {
	if len(args) > 1 {
		port, err := strconv.ParseInt(args[1], 10, 64)
		if err == nil {
			return int(port)
		}
	}

	return DefaultPort
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", BoxSizeHandler)

	port := ArgsToPortNumber(os.Args)
	fmt.Printf("BoxSize Server will run on port %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
