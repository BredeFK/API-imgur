// Author: Brede F. Klausen

// TODO : Comment code
// TODO : Make code more pretty and add more error handling
// TODO : Figure out how to add title and stuff when POSTing
// TODO : Maybe set up db and store som stuff there (like imageDeleteHash)
// TODO : Split to more files
// TODO : Make more helpingFunctions
// TODO : Change it so I have to use postman to post and delete
// TODO : Always run fmt, vet and lint before pushing!
// TODO : MAke more todos

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// DoStuff Does stuff :)
func DoStuff(method string, url string, body io.Reader, key string) ([]byte, int) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("authorization", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return payload, res.StatusCode
}

// HandleImage handles image requests
func HandleImage(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	length := len(parts) - 1

	switch parts[2] {

	case "POST":

		//	imageURL := strings.SplitAfter(r.URL.Path, "POST/")

		url := "https://api.imgur.com/3/image"

		//image := imageURL[1]

		payload := strings.NewReader("https://i.ytimg.com/vi/GUyxqjunVe0/maxresdefault.jpg")
		client := "Client-ID " + os.Getenv("CLIENT")
		body, _ := DoStuff(parts[2], url, payload, client)

		fmt.Fprintln(w, string(body))

	case "GET":
		if length == 3 {
			if parts[3] != "" {

				url := "https://api.imgur.com/3/image/" + parts[3]

				client := "Client-ID " + os.Getenv("CLIENT")

				body, status := DoStuff(parts[2], url, nil, client)
				if status != 200 {
					http.Error(w, "Could not get image", status)
				} else {
					fmt.Fprintln(w, string(body))
				}

			} else {
				http.Error(w, "imageHash can't be blank", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "No imageHash", http.StatusBadRequest)
		}

	case "DELETE":
		if length == 3 {
			if parts[3] != "" {

				url := "https://api.imgur.com/3/image/" + parts[3]

				client := "Client-ID " + os.Getenv("CLIENT")
				_, status := DoStuff(parts[2], url, nil, client)

				if status != 200 {
					http.Error(w, "Could not delete image", status)
				} else {
					fmt.Fprintln(w, "Deleted image with imageDeleteHash: ", parts[3])
				}

			} else {
				http.Error(w, "imageDeleteHash can't be blank", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "No imageDeleteHash", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Method has to be: POST, GET or DELETE", http.StatusBadRequest)
	}

}

func main() {

	// Handle Images
	http.HandleFunc("/image/", HandleImage)

	// Listen and serve for address "localhost:8080"
	http.ListenAndServe("localhost:8080", nil)
}
