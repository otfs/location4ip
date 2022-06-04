package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// parse command line args
	port := flag.Int("p", 8080, "-p <port>")
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)

	http.HandleFunc("/location", GetIpLocationHandle)
	log.Printf("listen addr %s", addr)
	err := http.ListenAndServe(addr, nil)
	log.Fatalln(err)
}
