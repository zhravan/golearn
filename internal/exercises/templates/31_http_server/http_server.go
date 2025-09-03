package http_server

import (
    "net/http"
)

// TODO:
// - Implement a minimal HTTP server:
//   - helloHandler: respond with "Hello, {name}!" using the request path segment.
//   - StartServer: register /hello route and start listening on :8080.
// - Keep function names; tests validate handler response and basic server start.

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: write a greeting response based on the path
}

func StartServer() {
    // TODO: register routes and start listening on :8080
}

func main() {
    // (no-op in skeleton)
}
