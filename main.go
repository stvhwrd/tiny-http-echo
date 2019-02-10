package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var port string

func echoRequest(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))
}

func init() {
	flag.StringVar(&port, "port", "", "The port for this HTTP server to listen on, eg. -port=8080")
	flag.Parse()

	if port == "" {
		log.Printf("Error: you must provide port flag at runtime.\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	log.Println("Listening on http://localhost:" + port)

	http.HandleFunc("/", echoRequest)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
