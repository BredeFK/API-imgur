package main

import (
	"fmt"
	"net/http"
)

// HandleInfo handles info on main page
func HandleInfo(w http.ResponseWriter, r *http.Request) {

	// Prints info about how to navigate page
	fmt.Fprintln(w, "Hello!\n\nImage info:\troot/image/\nComment info:\troot/comment/\n\n\nStatus:", http.StatusOK)
}
