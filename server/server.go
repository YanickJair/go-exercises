package server

import (
	"fmt"
	"net/http"
)

// Root - begin serving
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Let's rock this server")
}

/*
http.HandleFunc("/", server.Root)
http.ListenAndServe(":8000", nil)
*/
