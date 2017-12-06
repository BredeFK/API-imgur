// Author: Brede F. Klausen

// TODO : Comment code
// TODO : Fix info text
// TODO : Make code more pretty and add more error handling
// TODO : Maybe set up db and store som stuff there (like imageDeleteHash)
// TODO : Make more helpingFunctions
// TODO : Always run fmt, vet and lint before pushing!
// TODO : MAke more todos

package main

import (
	"net/http"
)

func main() {

	// Check if there is environment variables that's not empty
	if err, _, _ := CheckEnv(); err != 0 {
		return
	}

	// Handle info
	http.HandleFunc("/", HandleInfo)

	// Handle Images
	http.HandleFunc("/image/", HandleImage)

	// Handle Comments
	http.HandleFunc("/comment/", HandleComment)

	// Listen and serve for address "localhost:8080"
	http.ListenAndServe("localhost:8080", nil)
}
