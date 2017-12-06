package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO : Comment more here

// CheckEnv Checks the environment variables
func CheckEnv() (int, string, string) {

	// If environment variables are empty
	if os.Getenv("TOKEN") == "" || os.Getenv("CLIENT") == "" {

		// Say they are empty
		fmt.Println("Environment variables can not be empty!")

		// Return 1 (for 1 error)
		return 1, "", ""

	}

	// Return environment variables if they aren't blank
	return 0, "Bearer " + os.Getenv("TOKEN"), "Client-ID " + os.Getenv("CLIENT")

}

// DoStuff Does stuff :)
func DoStuff(method string, url string, body io.Reader, headers map[string]string) ([]byte, int) {

	// New request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Add header(s)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	imgBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return imgBody, res.StatusCode
}
