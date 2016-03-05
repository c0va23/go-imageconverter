package main

import (
	"net/http"
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write(([]byte)("OK"))
}

func main() {
	serverErr := http.ListenAndServe("127.0.0.1:5050", http.HandlerFunc(handler))
	if nil != serverErr {
		panic(serverErr)
	}
}
