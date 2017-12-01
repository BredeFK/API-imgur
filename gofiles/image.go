package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// HandleImage handles all image requests
func HandleImage(w http.ResponseWriter, r *http.Request) {

	// Split the url by the '/'
	parts := strings.Split(r.URL.Path, "/")

	// Get length of url after root/ - 1
	length := len(parts) - 1

	// Make switch on method
	switch parts[2] {

	// If the method is POST
	case "POST":

		// Handle POST
		ImagePOST(parts[2], w)

	// If the method is GET
	case "GET":

		// Handle GET
		ImageGET(parts[2], parts[3], length, w)

	// If the method is DELETE
	case "DELETE":

		// Handle DELETE
		ImageDELETE(parts[2], parts[3], length, w)

	// If the method is neither of the above
	default:

		// Give info about how to navigate images
		fmt.Fprintln(w, "Path: root/image/\n\nTo POST an image:\troot/image/POST/\nTo GET an image:\troot/image/GET/{{imageHash}}\nTo DELETE an image:\troot/image/DELETE/{{imageDeleteHash}}")
	}

}

// TODO : Write more error handling in ImagePOST

// ImagePOST function for posting image
func ImagePOST(method string, w http.ResponseWriter) {

	// URL for POSTing
	url := "https://api.imgur.com/3/image"

	// Make payload
	payload := strings.NewReader("https://i.ytimg.com/vi/GUyxqjunVe0/maxresdefault.jpg")

	// Make headers
	headers := map[string]string{
		"Authorization": "Bearer " + os.Getenv("TOKEN"),
		"Content-Type":  "application/json",
	}

	// Get body, but not the status code
	body, _ := DoStuff(method, url, payload, headers)

	// Print body
	fmt.Fprintln(w, string(body))
}

// ImageGET function for getting image
func ImageGET(method string, imageHash string, length int, w http.ResponseWriter) {

	// If the length is correct
	if length == 3 {

		// If the imageHash isn't blank
		if imageHash != "" {

			// URL for GETing
			url := "https://api.imgur.com/3/image/" + imageHash

			// Make headers
			headers := map[string]string{
				"Authorization": "Client-ID " + os.Getenv("CLIENT"),
			}

			// Get body and status code
			body, status := DoStuff(method, url, nil, headers)

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
func ImageDELETE(method string, imageDeleteHash string, length int, w http.ResponseWriter) {

	// If the length is correct
	if length == 3 {

		// If the imageDeleteHash isn't blank
		if imageDeleteHash != "" {

			// URL for DELETE image
			url := "https://api.imgur.com/3/image/" + imageDeleteHash

			// Make headers
			headers := map[string]string{
				"Authorization": "Bearer " + os.Getenv("TOKEN"),
			}

			// Get status code, but not body
			_, status := DoStuff(method, url, nil, headers)

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
