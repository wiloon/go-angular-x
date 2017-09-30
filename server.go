package main

import (
	"net/http"
)

type httpServer struct {
}

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"Status\":\"SUCCESS\"}"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./app0930/dist")))
	var server httpServer
	http.Handle("/foo", server)
	http.ListenAndServe(":3000", nil)
}
