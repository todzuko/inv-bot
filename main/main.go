package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	address := os.Getenv("ADDR")
	mux := http.NewServeMux()
	mux.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	if err = http.Serve(listener, mux); err != nil {
		log.Fatal(err)
	}
}
