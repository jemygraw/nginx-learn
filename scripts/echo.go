package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to listen")
	flag.Parse()

	endPoint := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(os.Stdout, req.Body)

		path := req.RequestURI
		if path == "/health" {
			w.Write([]byte("ok"))
		} else if path == "/css" {
			w.Header().Set("Content-Type", "text/css")
			w.Write([]byte("css"))
		} else if path == "/js" {
			w.Header().Set("Content-Type", "text/javascript")
			w.Write([]byte("js"))
		} else {
			w.Write([]byte("server on " + endPoint))
		}
	})
	http.ListenAndServe(endPoint, nil)
}
