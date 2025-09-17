package http_server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// helloHandler handles requests to paths beginning with /hello/.
// It extracts the first path segment after /hello/ and responds
// with "Hello, {name}!\n". If no name is provided, it returns
// a 400 Bad Request response.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/hello/")

	// Only take the first segment if multiple are present (e.g., /hello/a/b → "a").
	if i := strings.IndexByte(name, '/'); i >= 0 {
		name = name[:i]
	}

	// Handle the case where no name is given (e.g., /hello/).
	if name == "" {
		http.Error(w, "name required", http.StatusBadRequest)
		return
	}

	// Respond with plain text greeting.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

// init registers the /hello/ route when the package is loaded.
// This ensures the handler is available even when tests start
// their own server without calling StartServer.
func init() {
	http.HandleFunc("/hello/", helloHandler)
}

// StartServer starts an HTTP server on port :8080 using the
// default multiplexer. Any server error is logged instead of
// being silently ignored.
func StartServer() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("http server stopped: %v", err)
	}
}
