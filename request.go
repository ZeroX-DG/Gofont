package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// FetchHTML fetch the html response from a url
func FetchHTML(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

// DownloadFile download a file from url
func DownloadFile(url string, destination string) error {
	// Create the destination file
	destinationOutput, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationOutput.Close()

	// Send get request to URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(destinationOutput, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
