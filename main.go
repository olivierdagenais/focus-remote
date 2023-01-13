package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	_, ok := os.LookupEnv("REPLIT_DB_URL")
	var local bool
	if ok {
		// we're running under Replit
		local = false
		// TODO: use the Replit database
		// TODO: insist users authenticate
	} else {
		// we're running locally
		// TODO: use the in-memory database
		// TODO: don't bother checking for authorization
		local = true
	}
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "There is nothing interesting here, yet!")
		if !local {
			fmt.Fprintf(w, "<br />")
			userId := r.Header.Get("X-Replit-User-Id")
			if len(userId) > 0 {
				fmt.Fprintf(w, "You have been identified as Replit user %s", userId)
			} else {
				fmt.Fprintf(w, "Please log in")
				fmt.Fprint(w, "<div>\n<script authed=\"location.reload()\" src=\"https://auth.util.repl.co/script.js\"></script>\n</div")
			}
		}
	})

	// https://stackoverflow.com/a/43425461/98903
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Serving from http://localhost:%d in local mode? %v\n", listener.Addr().(*net.TCPAddr).Port, local)

	panic(http.Serve(listener, r))
}
