package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// HandleComment handles all comment requests
func HandleComment(w http.ResponseWriter, r *http.Request) {

	// Split the url by the '/'
	parts := strings.Split(r.URL.Path, "/")

	// Get length of url after root/ - 1
	length := len(parts) - 1

	// Make switch on method
	switch parts[2] {

	// If the method is POST
	case "POST":

		// Handle POST
		CommentPOST(parts[2], parts[3], parts[4], length, w)

	// If the method is GET
	case "GET":

		// Give error
		http.Error(w, "Not yet implemented", http.StatusNotImplemented)

	// If the method is DELETE
	case "DELETE":

		// Give error
		http.Error(w, "Not yet implemented", http.StatusNotImplemented)

	// If the method is neither of the above
	default:

		// Give info about how to navigate comments
		fmt.Fprintln(w, "Path: root/comment/\n\nTo POST an comment:\troot/comment/POST/{{imageHash}}/{{comment}}\nTo GET an image:\troot/comment/GET/{{commentId}}\nTo DELETE an image:\troot/comment/DELETE/{{commentId}}")
	}
}

// CommentPOST function for posting comment
func CommentPOST(method string, imageID string, comment string, length int, w http.ResponseWriter) {

	// If the length is correct
	if length == 4 {

		// If the comment isn't blank and are under 140 characters
		if comment != "" && len(comment) <= 140 {

			// URL for POSTing comment
			url := "https://api.imgur.com/3/comment"

			// Make payload
			payload := strings.NewReader("{\"image_id\": \"" + imageID + "\", \"comment\": \"" + comment + "\"}")

			// Make headers
			headers := map[string]string{
				"Authorization": "Bearer " + os.Getenv("TOKEN"),
				"Content-Type":  "application/json",
			}

			// Get body and http status code
			body, status := DoStuff(method, url, payload, headers)

			// If status is not OK
			if status != 200 {

				// Give error
				http.Error(w, "Could not post comment", status)
			} else {

				// Print body
				fmt.Fprintln(w, string(body))
			}
		} else {

			// Give error
			http.Error(w, "Comment can not be empty and can not be fore than 140 characters long", http.StatusBadRequest)
		}
	} else {

		// Give error
		http.Error(w, "Comment must have imageHash and comment-text", http.StatusBadRequest)
	}
}

// CommentGET function for getting comment
func CommentGET() {

}

// CommentDELETE function for deleting comment
func CommentDELETE() {

}
