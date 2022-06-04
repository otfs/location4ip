package main

import (
	"location4ip/config"
	"log"
	"net/http"
)

func main() {
	config.Init()
	http.HandleFunc("/location", GetIpLocationHandle)

	log.Printf("listen addr %s", config.Settings.BindAddress)
	err := http.ListenAndServe(config.Settings.BindAddress, nil)
	log.Fatalln(err)
}
