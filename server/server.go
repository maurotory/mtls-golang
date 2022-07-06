// server/server.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// set up handler to listen to root path
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("new request")
		fmt.Fprintf(writer, "hello world \n")
	})

	// serve on port 9090 of local host
	server := http.Server{
		Addr:    ":9090",
		Handler: handler,
	}

	fmt.Println("Listening on port 9090")
	// serve the endpoint with tls encryption
	if err := server.ListenAndServeTLS("cert/server.crt", "cert/server.key"); err != nil {
		log.Fatalf("error listening to port: %v", err)
	}

}
