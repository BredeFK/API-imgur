package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// HandleComment handles all comment requests
func HandleComment(w http.ResponseWriter, r *http.Request) {

	// Get Access token
	_, token, client := CheckEnv()

	// Split the url by the '/'
	parts := strings.Split(r.URL.Path, "/")

	// Make switch on method
	switch r.Method {

	// If the method is POST
	case "POST":

		// Handle POST
		CommentPOST(token, r, w)

	// If the method is GET
	case "GET":

		// Handle GET
		CommentGET(parts, client, r, w)

	// If the method is DELETE
	case "DELETE":

		// Handle DELETE
		//CommentDELETE(parts, client, r, w)

		// Give error while I try to fix CommentDELETE()
		http.Error(w, "Under maintenance", http.StatusServiceUnavailable)

	// If the method is neither of the above
	default:

		// Give info about how to navigate comments
		fmt.Fprintln(w, "Path: root/comment/\n\nTo POST an comment:\troot/comment/\nTo GET an image:\troot/comment/{{commentId}}\nTo DELETE an image:\troot/comment/{{commentId}}")
	}
}

// CommentPOST function for posting comment
func CommentPOST(token string, r *http.Request, w http.ResponseWriter) {

	// Make new commentPost struct
	comment := CommentPost{}

	// Get body from post to new struct
	err := json.NewDecoder(r.Body).Decode(&comment)

	// If there is an error
	if err != nil {

		// Give error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Close this
	defer r.Body.Close()

	// If the ImageID isn't blank
	if comment.ImageID != "" {

		// If the comment isn't blank and are under 140 characters
		if comment.Comment != "" && len(comment.Comment) <= 140 {

			// URL for POSTing comment
			url := "https://api.imgur.com/3/comment"

			// Make payload
			payload := strings.NewReader("{\"image_id\": \"" + comment.ImageID + "\", \"comment\": \"" + comment.Comment + "\", \"parent_id\": \"" + comment.ParentID + "\"}")

			// Make headers
			headers := map[string]string{
				"Authorization": token,
				"Content-Type":  "application/json",
			}

			// Get body and http status code
			body, status := DoStuff(r.Method, url, payload, headers)

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
		http.Error(w, "imageID can not be empty", http.StatusBadRequest)
	}
}

// CommentGET function for getting comment
func CommentGET(list []string, client string, r *http.Request, w http.ResponseWriter) {

	length := len(list) - 1
	commentID := list[2]

	// If the length is correct
	if length == 2 {

		// If the commentId isn't blank
		if commentID != "" {

			// Url for GETing comment
			url := "https://api.imgur.com/3/comment/" + commentID

			// Make headers
			headers := map[string]string{
				"Authorization": client,
				"Content-Type":  "application/json",
			}

			// Get body and status code
			body, status := DoStuff(r.Method, url, nil, headers)

			// If the status is not OK
			if status != 200 {

				// Give error
				http.Error(w, "Could not get comment", status)
			} else {

				// Print body
				fmt.Fprint(w, string(body))
			}
		} else {

			// Give error
			http.Error(w, "CommentId can not be blank", http.StatusBadRequest)
		}
	} else {

		// Give error
		http.Error(w, "No commentId", http.StatusBadRequest)
	}

}

// CommentDELETE function for deleting comment
func CommentDELETE(list []string, token string, r *http.Request, w http.ResponseWriter) {

	length := len(list) - 1
	commentID := list[2]

	// If the length is correct
	if length == 2 {

		// If the imageDeleteHash isn't blank
		if commentID != "" {

			// URL for DELETE image
			url := "https://api.imgur.com/3/comment/" + commentID

			// Make headers
			headers := map[string]string{
				"Authorization": token,
			}

			// Get status code, but not body
			body, status := DoStuff(r.Method, url, nil, headers)

			// If the status code is not OK
			if status != 200 {

				// Give error
				//	http.Error(w, "Could not delete comment", status)
				fmt.Fprint(w, string(body))
			} else {

				// Print confirmation
				fmt.Fprintln(w, "Deleted comment with commentId: ", commentID)
			}

		} else {

			// Give error
			http.Error(w, "commentId can't be blank", http.StatusBadRequest)
		}
	} else {

		// Give error
		http.Error(w, "No commentId", http.StatusBadRequest)
	}
}
