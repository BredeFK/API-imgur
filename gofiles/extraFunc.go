package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// TODO : Comment more here

// DoStuff Does stuff :)
func DoStuff(method string, url string, body io.Reader, headers map[string]string) ([]byte, int) {

	// New request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
	}

	// Add header(s)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	imgBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return imgBody, res.StatusCode
}
