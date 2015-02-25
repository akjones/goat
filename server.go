package main

import (
	"io"
	"log"
	"net/http"
)

func handleMessages(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `["somthing anything"]`)
}

func main() {
	http.HandleFunc("/messages", handleMessages)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
