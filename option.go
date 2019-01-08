package main

import (
	"errors"
	"os"
	"strings"
)

// Option The option received from the user
type Option struct {
	fontURL     string
	destination string
	cssFile     string
}

func isset(array []string, index int) bool {
	return len(array) > index && array[index] != ""
}

func getOption() (Option, error) {
	rawArgs := os.Args[1:]

	var args []string
	var flags = make(map[string]string)

	for _, arg := range rawArgs {
		if strings.HasPrefix(arg, "--") {
			rawFlag := strings.Replace(arg, "--", "", 1)
			flagParts := strings.Split(rawFlag, "=")
			flagName := flagParts[0]
			flagValue := flagParts[1]
			flags[flagName] = flagValue
		} else {
			args = append(args, arg)
		}
	}

	if !isset(args, 0) {
		return Option{}, errors.New("Please specify the font url")
	} else if !isset(args, 1) {
		return Option{}, errors.New("Please specify the destination")
	}

	// options
	googleFontURL := args[0]
	destinationPath := args[1]

	finalOption := Option{fontURL: googleFontURL, destination: destinationPath}

	if cssFile, ok := flags["cssFile"]; ok {
		finalOption.cssFile = cssFile
	}

	return finalOption, nil
}
