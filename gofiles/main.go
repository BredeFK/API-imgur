// Author: Brede F. Klausen

// TODO : Comment code
// TODO : Make code more pretty and add more error handling
// TODO : Figure out how to add title and stuff when POSTing
// TODO : Maybe set up db and store som stuff there (like imageDeleteHash)
// TODO : Make more helpingFunctions
// TODO : Change it so I have to use postman to post and delete
// TODO : Always run fmt, vet and lint before pushing!
// TODO : MAke more todos

package main

import (
	"net/http"
)

func main() {

	// Handle info
	http.HandleFunc("/", HandleInfo)

	// Handle Images
	http.HandleFunc("/image/", HandleImage)

	// Handle Comments
	http.HandleFunc("/comment/", HandleComment)

	// Listen and serve for address "localhost:8080"
	http.ListenAndServe("localhost:8080", nil)
}
