package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/messages", handleMessages)
	http.HandleFunc("/messages/something", handleMessage)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
