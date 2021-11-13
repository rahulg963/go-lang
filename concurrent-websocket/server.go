package concurrent_websocket

import (
	"io"
	"log"
	"net/http"
)

func SocketServerStart() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Gophers")
}

func Pprof() {
	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		log.Fatalf("Pprof failed: %v", err)
	}
}
