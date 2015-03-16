package main

import (
	"io"
	"net/http"
)

func handleMessages(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `["something","anything"]`)
}

func handleMessage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `"hello"`)
}
