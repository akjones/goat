package main

import (
	"io"
	"log"
	"net/http"
)

func handleMessages(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `["something","anything"]`)
}

func handleMessage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `"hello"`)
}

func main() {
	http.HandleFunc("/messages", handleMessages)
	http.HandleFunc("/messages/something", handleMessage)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
