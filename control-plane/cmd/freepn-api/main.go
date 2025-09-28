package main

import (
    "fmt"
    "log"
    "net/http"
)

// This is a very small stub for the FreePN control plane API.  A
// production implementation should allocate tunnel IP addresses,
// validate client tokens and program gateway nodes via the wgctrl
// library.  For now it only responds to GET requests at /health.
func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        _, _ = fmt.Fprintln(w, "freepn-api is running")
    })
    log.Println("Starting freepn-api on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
