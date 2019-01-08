package main

import (
	"errors"
	"os"
)

// Option The option received from the user
type Option struct {
	fontURL     string
	destination string
}

func isset(array []string, index int) bool {
	return len(array) > index && array[index] != ""
}

func getOption() (Option, error) {
	args := os.Args[1:]

	if !isset(args, 0) {
		return Option{}, errors.New("Please specify the font url")
	} else if !isset(args, 1) {
		return Option{}, errors.New("Please specify the destination")
	}

	// options
	googleFontURL := args[0]
	destinationPath := args[1]

	return Option{fontURL: googleFontURL, destination: destinationPath}, nil
}
