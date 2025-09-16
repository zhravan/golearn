package http_server

import (
	"fmt"
	"net/http"
	"strings"
)

// helloHandler responds with "Hello, {name}!" based on the URL path segment
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/hello/")
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func init() {
	http.HandleFunc("/hello/", helloHandler)
}

// StartServer registers the /hello/ route and starts an HTTP server on :8080
func StartServer() {
	http.ListenAndServe(":8080", nil)

}

func main() {
	StartServer()

}
