package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// HandleImage handles all image requests
func HandleImage(w http.ResponseWriter, r *http.Request) {

	// Get Access token and client-ID
	_, token, client := CheckEnv()

	// Split the url by the '/'
	parts := strings.Split(r.URL.Path, "/")

	// Make switch on method
	switch r.Method {

	// If the method is POST
	case "POST":

		// Handle POST
		ImagePOST(token, r, w)

	// If the method is GET
	case "GET":

		// Handle GET
		ImageGET(parts, client, r, w)

	// If the method is DELETE
	case "DELETE":

		// Handle DELETE
		ImageDELETE(parts, token, r, w)

	// If the method is neither of the above
	default:

		// Give info about how to navigate images
		fmt.Fprintln(w, "Path: root/image/\n\nTo POST an image:\troot/image/\nTo GET an image:\troot/image/{{imageHash}}\nTo DELETE an image:\troot/image/{{imageDeleteHash}}")
	}

}

// TODO : Write more error handling in ImagePOST

// ImagePOST function for posting image
func ImagePOST(token string, r *http.Request, w http.ResponseWriter) {

	post := ImagePost{}

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// URL for POSTing
	url := "https://api.imgur.com/3/image"

	// Make payload
	payload := strings.NewReader("{\"title\": \"" + post.Title + "\",\"image\": \"" + post.Image + "\"}")

	// Make headers
	headers := map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	}

	// Get body, but not the status code
	body, _ := DoStuff(r.Method, url, payload, headers)

	// Print body
	fmt.Fprintln(w, string(body))

}

// ImageGET function for getting image
func ImageGET(list []string, client string, r *http.Request, w http.ResponseWriter) {
	length := len(list) - 1
	imageHash := list[2]

	// If the length is correct
	if length == 2 {

		// If the imageHash isn't blank
		if imageHash != "" {

			// URL for GETing image
			url := "https://api.imgur.com/3/image/" + imageHash

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
				http.Error(w, "Could not get image", status)
			} else {

				// Print body
				fmt.Fprintln(w, string(body))
			}

		} else {

			// Give error
			http.Error(w, "imageHash can't be blank", http.StatusBadRequest)
		}
	} else {

		// Give error
		http.Error(w, "No imageHash", http.StatusBadRequest)
	}
}

// ImageDELETE function for deleting image
func ImageDELETE(list []string, token string, r *http.Request, w http.ResponseWriter) {

	length := len(list) - 1
	imageDeleteHash := list[2]

	// If the length is correct
	if length == 2 {

		// If the imageDeleteHash isn't blank
		if imageDeleteHash != "" {

			// URL for DELETE image
			url := "https://api.imgur.com/3/image/" + imageDeleteHash

			// Make headers
			headers := map[string]string{
				"Authorization": token,
			}

			// Get status code, but not body
			_, status := DoStuff(r.Method, url, nil, headers)

			// If the status code is not OK
			if status != 200 {

				// Give error
				http.Error(w, "Could not delete image", status)
			} else {

				// Print confirmation
				fmt.Fprintln(w, "Deleted image with imageDeleteHash: ", imageDeleteHash)
			}

		} else {

			// Give error
			http.Error(w, "imageDeleteHash can't be blank", http.StatusBadRequest)
		}
	} else {

		// Give error
		http.Error(w, "No imageDeleteHash", http.StatusBadRequest)
	}
}
