package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	address := "0.0.0.0:8084"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
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
