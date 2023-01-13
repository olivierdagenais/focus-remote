package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "There is nothing interesting here, yet!")
	})

	// https://stackoverflow.com/a/43425461/98903
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Serving from http://localhost:%d\n", listener.Addr().(*net.TCPAddr).Port)

	panic(http.Serve(listener, r))
}
