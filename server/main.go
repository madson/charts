package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	rootDir    = flag.String("rootDir", "./output", "file server directory")
	serverPort = flag.String("port", "8081", "server port number")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	err := startServer()
	if err != nil {
		return err
	}

	return nil
}

func startServer() error {
	server := newServer()
	server.configureRoutes()

	log.Println("running server at http://localhost:" + *serverPort)
	return http.ListenAndServe(":"+*serverPort, server.router)
}
