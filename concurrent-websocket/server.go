package concurrent_websocket

import (
	"io"
	"net/http"
)

func SocketServerStart() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Gophers")
}
