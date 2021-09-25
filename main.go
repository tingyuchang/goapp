package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func main()  {
	var port string
	flag.StringVar(&port, "port", "8080", "http port")
	flag.Parse()
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, this is port "+port+"\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


